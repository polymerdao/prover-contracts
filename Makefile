.SILENT:

# Hardcoded for simplicity
CONTRACT_NAMES = CrossL2ProverV2 \
				 MockCrossL2ProverV2 \
				 ReceiptParser

# Create the pattern for each contract
CONTRACT_JSON_PATTERNS := $(addsuffix .sol/*.json,$(addprefix ./out/,$(CONTRACT_NAMES)))

# Use wildcard to expand each pattern
CONTRACT_JSON_FILES = $(foreach pattern,$(CONTRACT_JSON_PATTERNS),$(wildcard $(pattern)))

.PHONY: test
test:
	forge test

# Deep fuzz + invariant campaign. Runs the full suite with cranked run counts; used on every PR
# (Foundry CI workflow) and locally. No separate nightly job needed.
.PHONY: fuzz
fuzz:
	FOUNDRY_FUZZ_RUNS=50000 FOUNDRY_INVARIANT_RUNS=5000 FOUNDRY_INVARIANT_DEPTH=200 forge test -vvv

.PHONY: build-prover
build-prover:
	forge build --force contracts/core/prove_api --sizes --deny warnings

.PHONY: release
release: build-prover
	./script/package-release.sh

.PHONY: build-contracts
build-contracts:
	echo "Building contracts"; \
	rm -frd ./out; \
	forge install; \
	forge build --skip test script -C contracts \
		--lib-paths lib --force

.PHONY: bindings-gen-ts
bindings-gen-ts: build-contracts
	echo "Generating TypeScript bindings..."; \
	rm -rfd ./src/evm/contracts/*; \
	npx typechain --target ethers-v6 --out-dir ./src/evm/contracts $(CONTRACT_JSON_FILES); \
	echo "Done."
