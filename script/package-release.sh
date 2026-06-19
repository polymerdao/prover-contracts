#!/usr/bin/env bash
#
# Freezes the deployable prove_api contracts into release/ so a release publishes immutable,
# deploy-ready artifacts:
#
#   <Contract>.creationCode.hex   frozen creation bytecode (deployed as-is, never recompiled)
#   <Contract>.abi.json           ABI
#   <Contract>.metadata.json      solc metadata (standard-JSON / explorer verification)
#   manifest.json                 tag, commit, compiler recipe, per-contract creationCode keccak
#   deploy.sh                     self-contained CREATE2 deploy script (bytes-in, no recompile)
#   SHA256SUMS                    download-integrity checksums over everything above
#
# Run via `make release` (which builds first). Inspired by ../solana-prover-contracts.
# See docs/deployments.md.
set -euo pipefail

ROOT="$(realpath "$(dirname "$(realpath "$0")")"/..)"
cd "$ROOT"

RELEASE_DIR="$ROOT/release"
FACTORY="0x4e59b44847b379578588920cA78FbF26c0B4956C" # canonical Arachnid CREATE2 deployer
TAG="${RELEASE_TAG:-$(git describe --tags --always --dirty)}"
COMMIT="$(git rev-parse HEAD)"

# Deployable prove_api contracts (default / metadata-free profile). name:path
CONTRACTS=(
  "CrossL2ProverV2:contracts/core/prove_api/CrossL2ProverV2.sol"
  "CrossL2Executor:contracts/core/prove_api/CrossL2Executor.sol"
)

rm -rf "$RELEASE_DIR"
mkdir -p "$RELEASE_DIR"

contracts_json="{}"
for entry in "${CONTRACTS[@]}"; do
  name="${entry%%:*}"
  id="${entry#*:}:${name}"

  code="$(forge inspect "$id" bytecode)"

  # The artifact MUST be metadata-free (no embedded ipfs source hash), or comments/formatting
  # would move the CREATE2 address. Guards against an accidental recipe regression.
  if [[ "$code" == *a264697066* ]]; then
    echo "ERROR: $name creationCode embeds an ipfs metadata hash -- not metadata-free." >&2
    echo "       The default profile must set bytecode_hash=\"none\" (see foundry.toml)." >&2
    exit 1
  fi

  printf '%s' "$code" >"$RELEASE_DIR/${name}.creationCode.hex"
  forge inspect "$id" abi --json >"$RELEASE_DIR/${name}.abi.json"
  forge inspect "$id" metadata --json >"$RELEASE_DIR/${name}.metadata.json"

  hash="$(cast keccak "$code")"
  ctor="$(
    jq -r '[.[] | select(.type=="constructor") | .inputs[].type] | "constructor(" + join(",") + ")"' \
      "$RELEASE_DIR/${name}.abi.json"
  )"
  contracts_json="$(
    jq --arg n "$name" --arg h "$hash" --arg c "$ctor" --arg f "${name}.creationCode.hex" \
      '.[$n] = {creationCodeHash: $h, constructorSignature: $c, creationCode: $f}' <<<"$contracts_json"
  )"
done

# compiler block mirrors foundry.toml [profile.default] (the deploy recipe) -- keep in sync
jq -n \
  --arg tag "$TAG" --arg commit "$COMMIT" --arg factory "$FACTORY" \
  --argjson contracts "$contracts_json" \
  '{
     tag: $tag, commit: $commit, factory: $factory,
     compiler: {
       solc: "0.8.15", profile: "default", via_ir: false,
       optimizer: true, optimizer_runs: 200, evm_version: "london",
       bytecode_hash: "none", cbor_metadata: false
     },
     contracts: $contracts
   }' >"$RELEASE_DIR/manifest.json"

# Ship the deploy script with the release so each version deploys exactly as it did when cut --
# infra downloads the release and runs release/deploy.sh, no repo checkout, no script drift.
cp "$ROOT/script/deploy.sh" "$RELEASE_DIR/deploy.sh"
chmod +x "$RELEASE_DIR/deploy.sh"

(cd "$RELEASE_DIR" && sha256sum -- *.hex *.json *.sh >SHA256SUMS)

echo "Release $TAG packaged in release/:"
ls -1 "$RELEASE_DIR"
