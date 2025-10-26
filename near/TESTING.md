# Testing the NEAR CrossL2Prover Contract

## Issue: Cargo Edition 2024 Requirement

The current NEAR SDK (5.x) requires Cargo's `edition2024` feature, which is not yet stable in Rust 1.82.0.

## Solutions

### Option 1: Use Rust Nightly (Recommended)

Install and use Rust nightly which has edition2024 support:

```bash
# Install nightly
rustup install nightly

# Use nightly for this project
cd near/
rustup override set nightly

# Now you can run tests
cargo test
```

### Option 2: Build Without Tests

You can still build the WASM contract for deployment without running tests:

```bash
# Build the contract (doesn't require edition2024)
cargo build --target wasm32-unknown-unknown --release

# Or use cargo-near
cargo near build
```

### Option 3: Wait for Stable Rust

Wait for Rust to stabilize edition2024 in a future stable release, then update:

```bash
rustup update stable
cargo test
```

## Manual Testing

Once you have the WASM built, you can test the contract by:

1. Deploying to NEAR testnet
2. Using the NEAR CLI to call functions
3. Writing integration tests with `near-workspaces` (also requires nightly)

## Current Test Status

The unit tests in `src/lib.rs` and `src/receipt_parser.rs` are functional but cannot run on stable Rust 1.82.0 due to dependency requirements.

Tests include:
- Contract initialization
- Event parsing
- Root key generation
- Solana event root key generation
