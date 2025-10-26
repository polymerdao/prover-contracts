# CrossL2ProverV2 - NEAR Smart Contract

A NEAR smart contract implementation of the CrossL2ProverV2, which validates cross-chain event proofs from Polymer's prove API.

## Overview

This contract is a Rust port of the Solidity `CrossL2ProverV2` contract, adapted for the NEAR blockchain. It validates event proofs for both EVM-based chains and Solana chains through Polymer's cross-chain proof infrastructure.

## Features

- **ECDSA Signature Verification**: Validates sequencer signatures over state roots
- **Event Validation**: Validates cross-chain event proofs for EVM chains
- **Solana Log Validation**: Validates Solana program log proofs
- **IAVL Proof Verification**: Verifies membership proofs in IAVL trees
- **Light Client Support**: Implements SequencerLightClient type

## Prerequisites

- Rust 1.70 or later
- NEAR CLI: `npm install -g near-cli-rs@latest`
- cargo-near: Install via `cargo install cargo-near`

## Building

Build the contract using cargo-near:

```bash
cd near/
cargo near build
```

This will compile the contract to WebAssembly and place the output in `target/near/`.

## Testing

**Note:** Tests require Rust nightly due to NEAR SDK dependencies. See [TESTING.md](./TESTING.md) for details.

Run the unit tests (requires nightly Rust):

```bash
rustup override set nightly
cargo test
```

Or skip tests and just build:

```bash
cargo near build
```

## Deployment

### 1. Create a NEAR Account

First, create a NEAR account for your contract (testnet example):

```bash
near account create-account fund-myself <your-contract-name>.testnet '1 NEAR' \
  use-manually-provided-public-key <your-public-key> \
  network-config testnet
```

### 2. Deploy the Contract

Deploy the compiled contract:

```bash
cargo near deploy <your-contract-name>.testnet without-init-call \
  network-config testnet \
  sign-with-keychain send
```

### 3. Initialize the Contract

Initialize with your parameters:

```bash
near contract call-function as-transaction <your-contract-name>.testnet new \
  json-args '{"client_type":"SequencerLightClient","sequencer":"0x<SEQUENCER_ADDRESS>","chain_id":"0x<CHAIN_ID_HEX>"}' \
  prepaid-gas '30 TeraGas' \
  attached-deposit '0 NEAR' \
  sign-as <your-account>.testnet \
  network-config testnet
```

Replace:
- `<SEQUENCER_ADDRESS>`: The Ethereum address of the trusted sequencer (40 hex chars, no 0x prefix)
- `<CHAIN_ID_HEX>`: The chain ID as a 32-byte hex string (64 hex chars, no 0x prefix)

## Usage

### Validate an Event Proof

```bash
near contract call-function as-transaction <your-contract-name>.testnet validate_event \
  json-args '{"proof":[<proof_bytes_array>]}' \
  prepaid-gas '300 TeraGas' \
  attached-deposit '0 NEAR' \
  sign-as <your-account>.testnet \
  network-config testnet
```

### Validate Solana Logs

```bash
near contract call-function as-transaction <your-contract-name>.testnet validate_sol_logs \
  json-args '{"proof":[<proof_bytes_array>]}' \
  prepaid-gas '300 TeraGas' \
  attached-deposit '0 NEAR' \
  sign-as <your-account>.testnet \
  network-config testnet
```

### View Functions

Inspect log identifier from a proof:

```bash
near contract call-function as-read-only <your-contract-name>.testnet inspect_log_identifier \
  json-args '{"proof":[<proof_bytes_array>]}' \
  network-config testnet
```

Get client type:

```bash
near contract call-function as-read-only <your-contract-name>.testnet get_client_type \
  json-args '{}' \
  network-config testnet
```

Ping (for health checks):

```bash
near contract call-function as-transaction <your-contract-name>.testnet ping \
  json-args '{}' \
  prepaid-gas '30 TeraGas' \
  attached-deposit '0 NEAR' \
  sign-as <your-account>.testnet \
  network-config testnet
```

## Contract API

### Initialization

```rust
pub fn new(
    client_type: String,
    sequencer: String,      // Hex string (with or without 0x prefix)
    chain_id: String,       // 32-byte hex string (with or without 0x prefix)
) -> Self
```

### Main Functions

#### validate_event
Validates an event proof from Polymer's prove API for non-Solana chains.

**Parameters:**
- `proof`: Vec<u8> - The complete proof bytes from Polymer's API

**Returns:**
```rust
ValidatedEvent {
    chain_id: u32,
    emitting_contract: String,  // Hex address
    topics: Vec<u8>,
    unindexed_data: Vec<u8>,
}
```

#### validate_sol_logs
Validates an event proof for Solana chains.

**Parameters:**
- `proof`: Vec<u8> - The complete proof bytes from Polymer's API

**Returns:**
```rust
ValidatedSolLogs {
    chain_id: u32,
    program_id: [u8; 32],
    log_messages: Vec<String>,
}
```

### View Functions

- `get_client_type() -> String`
- `get_light_client_type() -> LightClientType`
- `inspect_log_identifier(proof: Vec<u8>) -> (u32, u64, u32, u32)`
- `inspect_polymer_state(proof: Vec<u8>) -> ([u8; 32], u64, Vec<u8>)`

## Differences from Solidity Version

1. **Async Operations**: NEAR contracts don't support synchronous cross-contract calls like EVM
2. **Storage Model**: Uses NEAR's key-value storage instead of EVM's state trie
3. **Gas Model**: NEAR uses a different gas metering system
4. **Return Types**: View functions can return complex types directly
5. **Error Handling**: Uses Rust's `panic!` and `assert!` instead of Solidity's `revert`
6. **No External Calls**: The Solidity version uses `this.verifyMembership()` and `this.parseEvent()` - in Rust these are just direct function calls

## Security Considerations

- The sequencer address and chain ID are immutable after initialization
- All signature verification uses battle-tested cryptographic libraries (k256, sha2, sha3)
- IAVL proof verification follows the same algorithm as the Solidity version
- The contract panics on invalid proofs or signatures (fails closed)

## License

Apache-2.0 - Copyright 2024, Polymer Labs

## Resources

- [NEAR Documentation](https://docs.near.org)
- [Polymer Docs](https://docs.polymer.zone)
- [Original Solidity Contract](../contracts/core/prove_api/CrossL2ProverV2.sol)
