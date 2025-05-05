# Make sure to set required env vars to run:
# - CONTRACTS_DIR (path to the prover contracts)
# - CLI_PATH  (path to the prover cli)
# - RPC URLS To fork: 
#     - $BASE_MAINNET_RPC_URL   
#     - $OP_MAINNET_RPC_URL   
#     - $ETH_MAINNET_RPC_URL  

export CONTRACTS_DIR=~/Code/prover-contracts
export CLI_PATH=~/Code/fallback-prover
export FOUNDRY_DISABLE_NIGHTLY_WARNING=TRUE

# Anvil private key. Public knowledge.
export FUNDER_PKEY=0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80
export PKEY=$(cast wallet new | grep "Private key:" | awk '{print $3}')
export DEPLOYER=$(cast wallet address $PKEY)


# Kill any existing anvil processes
kill $(lsof -t -i:8546) & kill $(lsof -t -i:8547) & kill $(lsof -t -i:8545)
sleep 2
anvil -f $BASE_MAINNET_RPC_URL -p 8547 &> anvil-base.log \
    & anvil -f $OP_MAINNET_RPC_URL -p 8546 &> anvil-op.log \
    & anvil -f $ETH_MAINNET_RPC_URL -p 8545 &> anvil-eth.log &
sleep 2

# These are the local fork urls. Not the same as the ones in the anvil command above.
export ETH_RPC_URL=http://127.0.0.1:8545
export OP_RPC_URL=http://127.0.0.1:8546
export BASE_RPC_URL=http://127.0.0.1:8547
cast send $DEPLOYER -r $ETH_RPC_URL --private-key $FUNDER_PKEY --value 1ether --chain-id 1 &> fund-accounts.log
cast send $DEPLOYER -r $OP_RPC_URL --private-key $FUNDER_PKEY --value 1ether --chain-id 10 &> fund-accounts.log
cast send $DEPLOYER -r $BASE_RPC_URL --private-key $FUNDER_PKEY --value 1ether --chain-id 8453  &> fund-accounts.log

echo DEPLOYING Prover addr contracts
export OP_PROVER=$(forge create --private-key $PKEY ./contracts/core/native_fallback/L2/OPStackCannonProver.sol:OPStackCannonProver --chain-id 10 --broadcast -r $OP_RPC_URL | grep "Deployed to:" | awk '{print $3}')
echo deployed prover address to op @$OP_PROVER
export BASE_PROVER=$(forge create --private-key $PKEY ./contracts/core/native_fallback/L2/OPStackCannonProver.sol:OPStackCannonProver --chain-id 8453 --broadcast -r $BASE_RPC_URL | grep "Deployed to:" | awk '{print $3}')
echo deployed prover address to base @$BASE_PROVER

echo DEPLOYING L1 Contracts: 
# Deploy L1 contracts 
forge script script/DeployRegistry.s.sol --private-key $PKEY --broadcast --rpc-url $ETH_RPC_URL -vv &> deploy-eth.txt 
export SETTLEMENT_REGISTRY=$(grep "^0:\saddress\s0x[0-9a-fA-F]\{40\}$" deploy-eth.txt | awk '{print $3}')
echo L1Contracts Deployed: $SETTLEMENT_REGISTRY

# Now that it's bee ndeployed, we can now deploy each l2 contracts. 
export CHAIN_ID=10
forge script script/DeployNativeProver.s.sol:DeployNativeProverScript --private-key $PKEY --broadcast --rpc-url $OP_RPC_URL -vv &> deploy-op.txt 
export OP_NATIVE_PROVER=$(cat deploy-op.txt | grep "nativeProver: " | awk '{print $2}')
echo Op native prover Deployed: $OP_NATIVE_PROVER

export CHAIN_ID=8453
forge script script/DeployNativeProver.s.sol:DeployNativeProverScript --private-key $PKEY --broadcast --rpc-url $BASE_RPC_URL -vv &> deploy-base.txt
export BASE_NATIVE_PROVER=$(cat deploy-base.txt | grep "nativeProver: " | awk '{print $2}')

echo Base native prover Deployed: $BASE_NATIVE_PROVER

cd $CLI_PATH && make build 
OUTFILE=~/Code/prover-contracts/test/native_fallback/payload/native-proof.hex
$CLI_PATH/bin/native-proof update-and-prove \
    --src-l2-chain-id 10 \
    --dst-l2-chain-id 8453 \
    --src-l2-http-path $OP_RPC_URL \
    --dst-l2-http-path $BASE_RPC_URL \
    --src-contract-address 0x2E480B1096E5F705CC3e049685846ED3022aA226 \
    --src-storage-slot 0xc74f5e5058880b8779a44a8a762f001e42cef4d58647fb1d7d1548fb703c99f5 \
    --l1-http-path $ETH_RPC_URL \
    --l1-registry-address $SETTLEMENT_REGISTRY &> $OUTFILE
echo "Proof output written to $OUTFILE"

export PROOF=$(tail -n 1 $OUTFILE)
cd $CONTRACTS_DIR 

TEST_OUTPUT=~/Code/prover-contracts/integration_test.txt
forge test --match-test test_integration_proof -vvvv &> $TEST_OUTPUT 
echo "Test output written to $TEST_OUTPUT"

# cast code $OP_NATIVE_PROVER -r $OP_RPC_URL > ~/Code/prover-contracts/test/native_fallback/payload/op-prover-bytecode.hex 
# cast code $BASE_NATIVE_PROVER -r $OP_RPC_URL > ~/Code/prover-contracts/test/native_fallback/payload/base-prover-bytecode.hex 
