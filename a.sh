# Make sure to set required env vars to run:
# - CONTRACTS_DIR (path to the prover contracts)
# - CLI_PATH  (path to the prover cli)
# - RPC URLS: 
#     - $BASE_SEPOLIA_RPC_URL   
#     - $OP_SEPOLIA_RPC_URL   
#     - $ETH_SEPOLIA_RPC_URL  
# - 

export CONTRACTS_DIR=~/Code/prover-contracts
export CLI_PATH=~/Code/fallback-prover
export FOUNDRY_DISABLE_NIGHTLY_WARNING=TRUE
export OP_RPC_URL=$OP_SEPOLIA_RPC_URL
export BASE_RPC_URL=$BASE_SEPOLIA_RPC_URL

export DEPLOYER=$(cast wallet address $PKEY)

# Kill any existing anvil processes
# kill $(lsof -t -i:8546) & kill $(lsof -t -i:8547) & kill $(lsof -t -i:8545)
# sleep 2
# anvil -f $BASE_SEPOLIA_RPC_URL -p 8547 &> anvil-base.log \
#     & anvil -f $OP_SEPOLIA_RPC_URL -p 8546 &> anvil-op.log &
# sleep 2

# anvil -f $BASE_SEPOLIA_RPC_URL -p 8547 &> anvil-base.log &

# These are the local fork urls. Not the same as the ones in the anvil command above.



# echo DEPLOYING Prover addr contracts
# forge script script/DeployOpProver.s.sol --private-key $PKEY --rpc-url $OP_RPC_URL --broadcast -vv &> deploy-op-prover-op.txt
# export OP_PROVER=$(cat deploy-op-prover-op.txt | grep "OpstackCannonProver:  "  | awk '{print $2}')
# echo deployed prover address to op @$OP_PROVER
# forge script script/DeployOpProver.s.sol --private-key $PKEY --rpc-url $BASE_RPC_URL --broadcast -vv &> deploy-op-prover-base.txt
# export BASE_PROVER=$(cat deploy-op-prover-base.txt | grep "OpstackCannonProver:  " ./output.txt | awk '{print $2}')
# echo deployed prover address to base @$BASE_PROVER

# echo DEPLOYING L1 Contracts: 
# Deploy L1 contracts 
# forge script script/DeployRegistry.s.sol --private-key $PKEY --broadcast --rpc-url $ETH_SEPOLIA_RPC_URL -vv &> deploy-eth.txt 
# export SETTLEMENT_REGISTRY=$(grep "^0:\saddress\s0x[0-9a-fA-F]\{40\}$" deploy-eth.txt | awk '{print $3}')
# echo L1Contracts Deployed: $SETTLEMENT_REGISTRY


# Now that it's been deployed, we can now deploy each l2 contracts. 
# export CHAIN_ID=$(cast chain-id -r $OP_RPC_URL)
# forge script script/DeployNativeProver.s.sol:DeployNativeProverScript --private-key $PKEY --broadcast --rpc-url $OP_RPC_URL -vv &> deploy-op.txt 
# export OP_NATIVE_PROVER=$(cat deploy-op.txt | grep "nativeProver: " | awk '{print $2}')
# echo Op native prover Deployed: $OP_NATIVE_PROVER

# export CHAIN_ID=$(cast chain-id -r $BASE_RPC_URL)
# forge script script/DeployNativeProver.s.sol:DeployNativeProverScript --private-key $PKEY --broadcast --rpc-url $BASE_RPC_URL -vv &> deploy-base.txt
# export BASE_NATIVE_PROVER=$(cat deploy-base.txt | grep "nativeProver: " | awk '{print $2}')
# echo Base native prover Deployed: $BASE_NATIVE_PROVER

# --------------------------------------------------------------------------------------------------------------------

export OP_PROVER=0x4Dba9439caE5B8fC4EC881d73738c13545E6ecc8
export BASE_PROVER=0x4Dba9439caE5B8fC4EC881d73738c13545E6ecc8
export SETTLEMENT_REGISTRY=0xa102b8A5F1a46C0C3D516a9b2237F0D886B21479
export OP_NATIVE_PROVER=0xAf35b23F38fe8eE793AA6E03d7D54c93c8a6F4B3
export BASE_NATIVE_PROVER=0x63C0d99458E2d2153cF95ad8b800e191633F0d69


cd $CLI_PATH && make build 
OUTFILE=~/Code/prover-contracts/test/native_fallback/payload/native-proof.hex
$CLI_PATH/bin/native-proof proveNative \
    --src-l2-chain-id $(cast chain-id -r $OP_RPC_URL) \
    --dst-l2-chain-id $(cast chain-id -r $BASE_RPC_URL) \
    --src-l2-http-path $OP_RPC_URL \
    --dst-l2-http-path $BASE_RPC_URL \
    --src-contract-address 0x1BDD24840e119DC2602dCC587Dd182812427A5Cc \
    --src-storage-slot 0x0000000000000000000000000000000000000000000000000000000000000002 \
    --l1-http-path $ETH_SEPOLIA_RPC_URL \
    --l1-registry-address $SETTLEMENT_REGISTRY &> $OUTFILE

echo "Proof output written to $OUTFILE"

export PROOF=$(tail -n 1 $OUTFILE)
cd $CONTRACTS_DIR 

export CHAIN_ID=$(cast chain-id -r $BASE_RPC_URL)
TEST_OUTPUT=~/Code/prover-contracts/integration_test.txt
forge test --match-test test_integration_proof -vvvv &> $TEST_OUTPUT 
echo "Test output written to $TEST_OUTPUT"

