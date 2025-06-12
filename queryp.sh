#!/bin/bash

# Check if job ID and output file are provided
if [ $# -ne 1 ]; then
    echo "Usage: $0 <job_id> "
    exit 1
fi

JOB_ID=$1
OUT_FILE_NAME=devnet-log-$JOB_ID.txt

# Query a proof and save it to file
B64_PROOF_BYTES=$(curl -k -X POST "https://proof.testnet.polymer.zone/" \
 -H "Content-Type: application/json" \
 -H "Authorization: ${POLYMER_KEY_TESTNET}" \
 -d '{
        "jsonrpc": "2.0",
        "method": "polymer_queryProof",
        "params": ['$JOB_ID'],
        "id": 1
 }' | jq -r '.result.proof')

export PROOF_BYTES=0x$(echo $B64_PROOF_BYTES | base64 -d | xxd -p | tr -d ' \n' )
echo proof bytes: $PROOF_BYTES

echo running test and outputing to $OUT_FILE_NAME
forge test --match-test test_testnet_repro -vvvv > $OUT_FILE_NAME 