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
CONTRACT_ABI_PATTERNS = $(addsuffix .sol/*.abi.json,$(addprefix ./out/,$(CONTRACT_NAMES)))
CONTRACT_JSON_PATTERNS := $(addsuffix .sol/*.json,$(addprefix ./out/,$(CONTRACT_NAMES)))

# Use wildcard to expand each pattern
CONTRACT_ABI_FILES = $(foreach pattern,$(CONTRACT_ABI_PATTERNS),$(wildcard $(pattern)))
CONTRACT_BOTH_FILES = $(foreach pattern,$(CONTRACT_JSON_PATTERNS),$(wildcard $(pattern)))
CONTRACT_JSON_FILES = $(filter-out $(CONTRACT_ABI_FILES),$(CONTRACT_BOTH_FILES))

.PHONY: test
test:
	forge test

.PHONY: build-contracts
build-contracts:
	echo "Building contracts"; \
	rm -frd ./out; \
	forge install; \
	forge build --skip test script -C contracts \
		--lib-paths lib \
		--extra-output-files abi --force

# Libraries do not generate the correct ABI code (ChannelState enum causes errors)
# So the Ibc.sol abi generated code from abigen throws errors but is not needed.
# This is because the types exported are included in the ABIs of contracts and
# correctly interpreted (enums -> uint8 etc), the library methods are used inside
# of other contract methods and thus bindings for them do not need to be generated
# as they are not publicly exposed, but rather used within the contract itself.
#
# 	ABIGen issue ref: https://github.com/ethereum/solidity/issues/9278
.PHONY: bindings-gen-go
bindings-gen-go: build-contracts
	echo "Generating Go Prover Contracts bindings..."; \
	rm -rfd ./bindings/go/* ; \
	for abi_file in $(CONTRACT_ABI_FILES); do \
		abi_base=$$(basename $$(dirname $$abi_file)); \
		if [ "$$abi_base" = "Ibc.sol" ]; then \
			continue; \
		fi; \
		type=$$(basename $$abi_file .abi.json); \
		pkg=$$(basename $$abi_base .sol | tr "[:upper:]" "[:lower:]"); \
		mkdir -p ./bindings/go/$$pkg; \
		abigen --abi $$abi_file --pkg $$pkg --type $$type --out ./bindings/go/$$pkg/$$type.go; \
	done; \
	echo "Done."

.PHONY: bindings-gen-ts
bindings-gen-ts: build-contracts
	echo "Generating TypeScript bindings..."; \
	rm -rfd ./src/evm/contracts/*; \
	npx typechain --target ethers-v6 --out-dir ./src/evm/contracts $(CONTRACT_JSON_FILES); \
	echo "Done."
