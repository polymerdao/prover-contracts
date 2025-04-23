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
  ContractItem,
} from "./evm/schemas/contract";
import { Logger } from "./utils/cli";
import { CREATE_2_FACTORY, DEFAULT_DEPLOYER } from "./utils/constants";
import { Chain } from "./evm/chain";
import * as proverContractFactories from "./evm/contracts/index";
import { isMultisig } from "./evm/schemas/multisig";
import { SendingAccountRegistry } from "./evm/schemas/sendingAccount";
import { Wallet } from "ethers";
import { HDNodeWallet } from "ethers";

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
      `[${chain.chainName}-${chain.deploymentEnvironment}]: deploying ${contract.name
      } with args: [${constructorData.args}] with libraries: ${JSON.stringify(
        constructorData.libraries
      )} ${create2Salt ? `and create2 salt: ${create2Salt}` : ""}`
    );

    const deployer = accountRegistry.mustGet(
      contract.deployer ? contract.deployer : DEFAULT_DEPLOYER
    );

    if (isMultisig(deployer)) {
      throw new Error("Contract Deployments not supported for multisig wallets!");
    }

    const updatedNonces = await updateNoncesForSender(
      nonces,
      deployer.address,
      () => deployer.getNonce()
    );
    const nonce = updatedNonces[deployer.address];

    const deployedAddr = await deploy(
      deployer,
      constructorData.factory,
      constructorData.args,
      contract,
      nonce,
      create2Salt,
      chain.chainName,
      dryRun,
      logger
    )

    env[contract.name] = deployedAddr;
    // update contract in registry as output result
    contract.address = deployedAddr;
    contract.deployer = deployer.address;
    contract.abi = constructorData.contractFactoryConstructor.abi;

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
    logger.error(`[${chain.chainName}-${chain.deploymentEnvironment}] deploy ${contract.name} failed: ${err}`);
    throw err;
  }
};

const deploy = async (
  deployer: Wallet | HDNodeWallet,
  factory: any,
  args: any[],
  contract: ContractItem,
  nonce: number,
  create2Salt: string | undefined,
  chainName: string,
  dryRun: boolean,
  logger: Logger,
): Promise<string> => {
  if (dryRun) {
    logger.info('dry run is set. will not deploy')
    return `new.${contract.name}.address`
  }

  // no create2 is used so do a normal contract deployment
  if (!create2Salt) {
    logger.info('no create2 salt provided. doing normal deployment')
    const deployed = await factory
      .connect(deployer)
      .deploy(...args, {
        nonce,
      });
    await deployed.deploymentTransaction()?.wait(1);
    return await deployed.getAddress();
  }

  // this handles an idempotent create2 deployment. The workflow goes as follows:
  // 1. calculate the create2 address where the contract would be deployed at given its bytecode
  // 2. check if there's code already deployed to that address
  // 3. if there is, don't do anything. Otherwise, deploy it
  // 4. return the calculated address

  const initData = await factory.connect(deployer).getDeployTransaction(...args);
  const deployedAddr = getCreate2Address(CREATE_2_FACTORY, create2Salt, keccak256(initData.data))
  logger.info(`calculated create2 contract address: ${deployedAddr}`)

  const code = await deployer.provider?.getCode(deployedAddr)

  // getCode() will return `0x` for non-existing contracts
  if (code !== undefined && code.length > 2) {
    logger.info(
      `found contract at ${deployedAddr} on chain ${chainName}, with size ${code.length / 2 - 1}. ` +
      'will not re-deploy'
    )
    return deployedAddr
  }

  logger.info(`deploying contract to ${chainName}, on address ${deployedAddr}`)
  const txResponse = await deployer.sendTransaction({
    to: CREATE_2_FACTORY,
    data: `${create2Salt}${initData.data.slice(2)}`,
    nonce: nonce,
  });
  await txResponse.wait(1);

  return deployedAddr
}
