.SILENT:

# Hardcoded for simplicity
CONTRACT_NAMES = CrossL2ProverV2 \
				 MockCrossL2ProverV2 \
				 ReceiptParser \
				 SequencerSignatureVerifier \
				 NativeProver \
				 Registry \
				 OPStackBedrockProver \
				 OPStackCannonProver \
				 RegistryTypes

# Create the pattern for each contract
CONTRACT_JSON_PATTERNS := $(addsuffix .sol/*.json,$(addprefix ./out/,$(CONTRACT_NAMES)))

# Use wildcard to expand each pattern
CONTRACT_JSON_FILES = $(foreach pattern,$(CONTRACT_JSON_PATTERNS),$(wildcard $(pattern)))

# Two profiles: default (deploy recipe, via_ir=false) covers prove_api; native (via_ir=true) covers
# native_fallback. Run BOTH so nothing is missed -- neither profile compiles the whole repo alone.
.PHONY: test
test:
	forge test
	FOUNDRY_PROFILE=native forge test

# Deep fuzz + invariant campaign. Runs the full suite with cranked run counts under both profiles;
# used on every PR (Foundry CI workflow) and locally. No separate nightly job needed.
.PHONY: fuzz
fuzz:
	FOUNDRY_FUZZ_RUNS=50000 FOUNDRY_INVARIANT_RUNS=5000 FOUNDRY_INVARIANT_DEPTH=200 forge test -vvv
	FOUNDRY_PROFILE=native FOUNDRY_FUZZ_RUNS=50000 FOUNDRY_INVARIANT_RUNS=5000 FOUNDRY_INVARIANT_DEPTH=200 forge test -vvv

.PHONY: build-prover
build-prover:
	forge build --force contracts/core/prove_api --sizes --deny warnings

.PHONY: build-native
build-native:
	FOUNDRY_PROFILE=native forge build contracts/core/native_fallback --sizes --deny warnings

.PHONY: release
release: build-prover
	./script/package-release.sh

# Bindings need every ABI, so build under the native profile (via_ir=true compiles all contracts;
# the default profile can't compile native_fallback).
.PHONY: build-contracts
build-contracts:
	echo "Building contracts"; \
	rm -frd ./out; \
	forge install; \
	FOUNDRY_PROFILE=native forge build --skip test script -C contracts \
		--lib-paths lib --force

.PHONY: bindings-gen-ts
bindings-gen-ts: build-contracts
	echo "Generating TypeScript bindings..."; \
	rm -rfd ./src/evm/contracts/*; \
	npx typechain --target ethers-v6 --out-dir ./src/evm/contracts $(CONTRACT_JSON_FILES); \
	echo "Done."
