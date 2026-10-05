#!/usr/bin/env bash
#
# Deploys a frozen release artifact via CREATE2. This script SHIPS INSIDE THE RELEASE, next to the
# creationCode it deploys, so every version deploys exactly as it did when cut -- no repo checkout,
# no script drift. It NEVER recompiles: it reads the prebuilt creationCode, appends the env-specific
# constructor args, and deploys through the canonical Arachnid factory (deterministic across chains).
#
# Run from a downloaded release: `gh release download <version> -D release && bash release/deploy.sh`
# (CI/infra) or `make release && bash release/deploy.sh` (local). See docs/deployments.md.
#
# Required env: RPC_URL, DEPLOYER_PRIVATE_KEY, CREATE2_SALT, SEQUENCER_PUB_KEY, PEPTIDE_CHAIN_ID
# Optional env: EXPECTED_ADDRESS (assert the canonical address), CLIENT_TYPE (default "proof_api")
set -euo pipefail

# The frozen creationCode ships alongside this script in the release; resolve it relative to the
# script so a downloaded release is fully self-contained.
SCRIPT_DIR="$(cd "$(dirname "$0")" && pwd)"

FACTORY="0x4e59b44847b379578588920cA78FbF26c0B4956C" # canonical Arachnid CREATE2 deployer
CLIENT_TYPE="${CLIENT_TYPE:-proof_api}"
ARTIFACT="$SCRIPT_DIR/CrossL2ProverV2.creationCode.hex"

for v in RPC_URL DEPLOYER_PRIVATE_KEY CREATE2_SALT SEQUENCER_PUB_KEY PEPTIDE_CHAIN_ID; do
  if [ -z "${!v:-}" ]; then
    echo "ERROR: missing required env var: $v" >&2
    exit 1
  fi
done

if [ ! -f "$ARTIFACT" ]; then
  echo "ERROR: $ARTIFACT not found -- run 'gh release download <version> -D release' or 'make release'" >&2
  exit 1
fi

# initCode = frozen creationCode ++ abi.encode(constructor args). The artifact is deployed as-is.
code="$(cat "$ARTIFACT")"
args="$(cast abi-encode 'constructor(string,address,bytes32)' "$CLIENT_TYPE" "$SEQUENCER_PUB_KEY" "$PEPTIDE_CHAIN_ID")"
init="${code}${args#0x}"

# predicted CREATE2 address = keccak256(0xff ++ factory ++ salt ++ keccak256(initCode))[12:]
full="$(cast keccak "$(cast concat-hex 0xff "$FACTORY" "$CREATE2_SALT" "$(cast keccak "$init")")")"
predicted="0x${full: -40}"
echo "predicted CREATE2 address: $predicted"

if [ -n "${EXPECTED_ADDRESS:-}" ]; then
  if [ "$(tr '[:upper:]' '[:lower:]' <<<"$predicted")" != "$(tr '[:upper:]' '[:lower:]' <<<"$EXPECTED_ADDRESS")" ]; then
    echo "ERROR: predicted $predicted != EXPECTED_ADDRESS $EXPECTED_ADDRESS" >&2
    echo "       inputs (salt/args/recipe) or the release version don't match the canonical deploy." >&2
    exit 1
  fi
  echo "matches EXPECTED_ADDRESS"
fi

# The factory MUST be present -- never fall back to a nonce-based deploy (that breaks the address).
if [ "$(cast code "$FACTORY" --rpc-url "$RPC_URL")" = "0x" ]; then
  echo "ERROR: CREATE2 factory $FACTORY is not deployed on this chain. Deploy it first." >&2
  exit 1
fi

if [ "$(cast code "$predicted" --rpc-url "$RPC_URL")" != "0x" ]; then
  echo "already deployed at $predicted -- nothing to do"
  exit 0
fi

cast send \
  "$FACTORY" \
  "${CREATE2_SALT}${init#0x}" \
  --private-key "$DEPLOYER_PRIVATE_KEY" \
  --rpc-url "$RPC_URL" >/dev/null

if [ "$(cast code "$predicted" --rpc-url "$RPC_URL")" = "0x" ]; then
  echo "ERROR: deployment sent but no code at $predicted" >&2
  exit 1
fi

echo "deployed CrossL2ProverV2 at $predicted"
