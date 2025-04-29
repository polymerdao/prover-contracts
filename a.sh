export CONTRACTS_DIR=~/Code/prover-contracts
# Anvil private key. Public knowledge.
export FUNDER_PKEY=0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80
export PKEY=$(cast wallet new | grep "Private key:" | awk '{print $3}')
export DEPLOYER=$(cast wallet address $PKEY)
export ETH_RPC_URL=127.0.0.1:8545
export OP_RPC_URL=127.0.0.1:8546
export BASE_RPC_URL=127.0.0.1:8547
cast send $DEPLOYER -r $ETH_RPC_URL --private-key $FUNDER_PKEY --value 1ether --chain-id 1
cast send $DEPLOYER -r $OP_RPC_URL --private-key $FUNDER_PKEY --value 1ether --chain-id 10
cast send $DEPLOYER -r $BASE_RPC_URL --private-key $FUNDER_PKEY --value 1ether --chain-id 8453 

export CLI_PATH=~/Code/fallback-prover/bin/native-proof

echo DEPLOYING Prover addr contracts
export OP_PROVER=$(forge create --private-key $PKEY ./contracts/core/native_fallback/L2/OPStackCannonProver.sol:OPStackCannonProver --chain-id 10 --broadcast -r $OP_RPC_URL | grep "Deployed to:" | awk '{print $3}')
echo deployed prover address to op @$OP_PROVER
export BASE_PROVER=$(forge create --private-key $PKEY ./contracts/core/native_fallback/L2/OPStackCannonProver.sol:OPStackCannonProver --chain-id 8453 --broadcast -r $BASE_RPC_URL | grep "Deployed to:" | awk '{print $3}')
echo deployed prover address to base @$BASE_PROVER

echo DEPLOYING L1 Contracts: 
# Deploy L1 contracts 
export SETTLEMENT_REGISTRY=$(forge script script/DeployRegistry.s.sol --private-key $PKEY --broadcast --rpc-url $ETH_RPC_URL -vv | grep "settlementRegistry: " | awk '{print $2}') 

echo L1Contracts Deployed: $SETTLEMENT_REGISTRY

# Now that it's bee ndeployed, we can now deploy each l2 contracts. 
export CHAIN_ID=10
export OP_NATIVE_PROVER=$(forge script script/DeployNativeProver.s.sol:DeployNativeProverScript --private-key $PKEY --broadcast --rpc-url $OP_RPC_URL -vv | grep "nativeProver: " | awk '{print $2}') 
echo Op native prover Deployed: $OP_NATIVE_PROVER

export CHAIN_ID=8453
export BASE_NATIVE_PROVER=$(forge script script/DeployNativeProver.s.sol:DeployNativeProverScript --private-key $PKEY --broadcast --rpc-url $BASE_RPC_URL -vv | grep "nativeProver: " | awk '{print $2}') 
echo Base native prover Deployed: $BASE_NATIVE_PROVER

PROOF=$($CLI_PATH prove \
    --src-l2-chain-id 10 \
    --dst-l2-chain-id 8453 \
    --src-l2-http-path http://127.0.0.1:8546 \
    --dst-l2-http-path http://127.0.0.1:8547 \
    --src-l2-contract-address 0x2E480B1096E5F705CC3e049685846ED3022aA226 \
    --src-l2-storage-slot 0xc74f5e5058880b8779a44a8a762f001e42cef4d58647fb1d7d1548fb703c99f5 \
    --l1-http-path http://127.0.0.1:8545 \
    --l1-registry-address $SETTLEMENT_REGISTRY)


OUTFILE=~/Code/prover-contracts/test/native_fallback/payload/native-proof.hex
echo $PROOF > $OUTFILE
echo "Proof written to $OUTFILE"

cast code $OP_NATIVE_PROVER --rpc-url $OP_RPC_URL > ~/Code/prover-contracts/test/native_fallback/payload/op-prover-bytecode.hex 
cast code $BASE_NATIVE_PROVER --rpc-url $OP_RPC_URL > ~/Code/prover-contracts/test/native_fallback/payload/base-prover-bytecode.hex 
