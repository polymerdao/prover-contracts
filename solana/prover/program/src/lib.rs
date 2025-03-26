#![allow(clippy::arithmetic_side_effects)]
#![deny(missing_docs)]
#![cfg_attr(not(test), forbid(unsafe_code))]

//! An ERC20-like Token program for the Solana blockchain

pub mod error;
pub mod instruction;
pub mod processor;

#[cfg(not(feature = "no-entrypoint"))]
mod entrypoint;

/// Export current sdk types for downstream users building with a different sdk
/// version
pub mod solana_program {
    #![allow(missing_docs)]
    pub mod entrypoint {
        pub use solana_program_error::ProgramResult;
    }
    pub mod instruction {
        pub use solana_instruction::{AccountMeta, Instruction};
    }
    pub mod program_error {
        pub use solana_program_error::{PrintProgramError, ProgramError};
    }
    pub mod program_option {
        pub use solana_program_option::COption;
    }
    pub mod program_pack {
        pub use solana_program_pack::{IsInitialized, Pack, Sealed};
    }
    pub mod pubkey {
        pub use solana_pubkey::{Pubkey, PUBKEY_BYTES};
    }
}

solana_pubkey::declare_id!("1111111QLbz7JHiBTspS962RLKV8GndWFwiEaqKM");
