import { getCreate2Address, keccak256 } from "ethers";
import {
  DeployedContractObject,
  StringToStringMap,
  renderArgs,
  renderString,
  writeDeployedContractToFile,
} from "./utils/io";
import assert from "assert";
import {
  SingleSigAccountRegistry,
} from "./evm/schemas/account";
import {
  ContractRegistry,
  ContractRegistryLoader,
  ContractItem,
} from "./evm/schemas/contract";
import { Logger } from "./utils/cli";
import { CREATE_2_FACTORY, DEFAULT_DEPLOYER } from "./utils/constants";
import { Chain } from "./evm/chain";
import * as proverContractFactories from "./evm/contracts/index";
import { isMultisig } from "./evm/schemas/multisig";
import { SendingAccountRegistry } from "./evm/schemas/sendingAccount";

export async function updateNoncesForSender(
  nonces: Record<string, number>,
  address: string,
  nonceUpdate: () => Promise<number>
) {
  // To avoid nonce too low bug, we manually increment nonces for each account
  if (!(address in nonces)) {
    nonces[address] = await nonceUpdate();
  } else {
    nonces[address] += 1;
  }
  return nonces;
}

// Converts a factory to a name. If a solc version is specified (which needs to happen for multiple solc versions in a foundry/hardhat project, then it will return the file with the constructed name)
export const getFactoryFileName = (
  factoryName: string,
  solcVersion: string | undefined
) => {
  if (!solcVersion) return `${factoryName}__factory`;

  // Filter version string to remove periods e.g. 0.8.15 -> 0815
  const versionStr = solcVersion
    .split("")
    .filter((c) => c !== "." && c !== "v")
    .join("");
  return `${factoryName}${versionStr}__factory`;
};

/**
 * Return deployment libraries, factory, factory constructor,
 * and rendered arguments for a contract deployment
 */
const getDeployData = (
  factoryName: string,
  deployArgs: any[] | undefined,
  env: StringToStringMap,
  libraries: any[] = [],
  init: { args: any[]; signature: string } | undefined,
  contractFactories: Record<string, any>,
  solcVersion: string | undefined
) => {
  const contractFactoryFileName = getFactoryFileName(factoryName, solcVersion);
  // @ts-ignore
  const contractFactoryConstructor = contractFactories[contractFactoryFileName];
  assert(
    contractFactoryConstructor,
    `cannot find contract factory constructor for contract: ${factoryName}`
  );

  let libs: any = {};
  libraries.forEach((arg: any) => {
    libs[arg.name] = renderString(arg.address, env);
  });
  libs = [libs];

  const factory = new contractFactoryConstructor(...libs);
  if (!factory) {
    throw new Error(
      `cannot load contract factory for contract: ${factoryName} from factory file: ${contractFactoryFileName}`
    );
  }

  return {
    args: renderArgs(deployArgs, init, env),
    libraries: libs,
    factory,
    contractFactoryConstructor,
  };
};

export const deployContract = async (
  chain: Chain,
  accountRegistry: SingleSigAccountRegistry | SendingAccountRegistry,
  contract: ContractItem,
  logger: Logger,
  dryRun: boolean = false,
  writeContracts: boolean = true, // True if you want to save persisted artifact files.
  extraContractFactories: Record<string, object> = {},
  nonces: Record<string, number> = {},
  env: StringToStringMap = {},
  create2Salt: string | undefined = undefined
) => {
  // merge extra contracts into the registry
  const contractFactories = {
    ...proverContractFactories,
    ...extraContractFactories,
  };

  try {
    const factoryName = contract.factoryName
      ? contract.factoryName
      : contract.name;

    const constructorData = getDeployData(
      factoryName,
      contract.deployArgs,
      env,
      contract.libraries,
      contract.init,
      contractFactories,
      contract.solcVersion
    );

    logger.info(
      `[${chain.chainName}-${chain.deploymentEnvironment}]: deploying ${
        contract.name
      } with args: [${constructorData.args}] with libraries: ${JSON.stringify(
        constructorData.libraries
      )} ${create2Salt ? `and create2 salt: ${create2Salt}` : ""}`
    );
    let deployedAddr = `new.${contract.name}.address`;
    const deployer = accountRegistry.mustGet(
      contract.deployer ? contract.deployer : DEFAULT_DEPLOYER
    );

    if (isMultisig(deployer)) {
      throw new Error(
        "Contract Deployments not supported for multisig wallets!"
      );
    }

    const updatedNonces = await updateNoncesForSender(
      nonces,
      deployer.address,
      () => deployer.getNonce()
    );
    const nonce = updatedNonces[deployer.address];
    if (!dryRun) {
      if (create2Salt) {
        // Make a call to create2factory instead of a normal eth deploy tx if create2Salt is provided
        const initData = await constructorData.factory
          .connect(deployer)
          .getDeployTransaction(...constructorData.args);

        const txResponse = await deployer.sendTransaction({
          to: CREATE_2_FACTORY,
          data: `${create2Salt}${initData.data.slice(2)}`,
          nonce: nonce,
        });
        await txResponse.wait(1);
        deployedAddr = getCreate2Address(
          CREATE_2_FACTORY,
          create2Salt,
          keccak256(initData.data)
        );
      } else {
        // Otherwise proceed with normal contract deployment
        const deployed = await constructorData.factory
          .connect(deployer)
          .deploy(...constructorData.args, {
            nonce,
          });
        await deployed.deploymentTransaction()?.wait(1);
        deployedAddr = await deployed.getAddress();
      }
    }
    // save deployed contract address for its dependencies
    logger.info(
      `deployed contract ${chain.chainName} ${contract.name} at ${deployedAddr}`
    );
    env[contract.name] = deployedAddr;
    // update contract in registry as output result
    contract.address = deployedAddr;
    contract.deployer = deployer.address;
    contract.abi = constructorData.contractFactoryConstructor.abi;
    logger.info(
      `[${chain.chainName}-${chain.deploymentEnvironment}]: deployed ${contract.name} to address: ${deployedAddr}`
    );
    if (writeContracts) {
      const contractObject: DeployedContractObject = {
        factory: factoryName,
        address: deployedAddr,
        abi: constructorData.contractFactoryConstructor.abi,
        bytecode: constructorData.factory.bytecode,
        name: contract.name,
        args: constructorData.args,
        libraries: constructorData.libraries,
        solcVersion: contract.solcVersion,
      };
      writeDeployedContractToFile(chain, contractObject);
    }
    return contract;
  } catch (err) {
    logger.error(
      `[${chain.chainName}-${chain.deploymentEnvironment}] deploy ${contract.name} failed: ${err}`
    );
    throw err;
  }
};
