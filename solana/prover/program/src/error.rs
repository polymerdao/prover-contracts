//! Error types

use {
    num_derive::FromPrimitive,
    solana_decode_error::DecodeError,
    solana_msg::msg,
    solana_program_error::{PrintProgramError, ProgramError},
    thiserror::Error,
};

/// Errors that may be returned by the Token program.
#[derive(Clone, Debug, Eq, Error, FromPrimitive, PartialEq)]
pub enum ProverError {
    // 0
    /// Lamport balance below rent-exempt threshold.
    #[error("Lamport balance below rent-exempt threshold")]
    NotRentExempt,
    /// Insufficient funds for the operation requested.
    #[error("Insufficient funds")]
    InsufficientFunds,
    /// Invalid Mint.
    #[error("Invalid Mint")]
    InvalidMint,
    /// Account not associated with this Mint.
    #[error("Account not associated with this Mint")]
    MintMismatch,
    /// Owner does not match.
    #[error("Owner does not match")]
    OwnerMismatch,

    // 5
    /// This token's supply is fixed and new tokens cannot be minted.
    #[error("Fixed supply")]
    FixedSupply,
    /// The account cannot be initialized because it is already being used.
    #[error("Already in use")]
    AlreadyInUse,
    /// Invalid number of provided signers.
    #[error("Invalid number of provided signers")]
    InvalidNumberOfProvidedSigners,
    /// Invalid number of required signers.
    #[error("Invalid number of required signers")]
    InvalidNumberOfRequiredSigners,
    /// State is uninitialized.
    #[error("State is uninitialized")]
    UninitializedState,

    // 10
    /// Instruction does not support native tokens
    #[error("Instruction does not support native tokens")]
    NativeNotSupported,
    /// Non-native account can only be closed if its balance is zero
    #[error("Non-native account can only be closed if its balance is zero")]
    NonNativeHasBalance,
    /// Invalid instruction
    #[error("Invalid instruction")]
    InvalidInstruction,
    /// State is invalid for requested operation.
    #[error("State is invalid for requested operation")]
    InvalidState,
    /// Operation overflowed
    #[error("Operation overflowed")]
    Overflow,

    // 15
    /// Account does not support specified authority type.
    #[error("Account does not support specified authority type")]
    AuthorityTypeNotSupported,
    /// This token mint cannot freeze accounts.
    #[error("This token mint cannot freeze accounts")]
    MintCannotFreeze,
    /// Account is frozen; all account operations will fail
    #[error("Account is frozen")]
    AccountFrozen,
    /// Mint decimals mismatch between the client and mint
    #[error("The provided decimals value different from the Mint decimals")]
    MintDecimalsMismatch,
    /// Instruction does not support non-native tokens
    #[error("Instruction does not support non-native tokens")]
    NonNativeNotSupported,
}
impl From<ProverError> for ProgramError {
    fn from(e: ProverError) -> Self {
        ProgramError::Custom(e as u32)
    }
}
impl<T> DecodeError<T> for ProverError {
    fn type_of() -> &'static str {
        "TokenError"
    }
}

impl PrintProgramError for ProverError {
    fn print<E>(&self)
    where
        E: 'static
            + std::error::Error
            + DecodeError<E>
            + PrintProgramError
            + num_traits::FromPrimitive,
    {
        match self {
            ProverError::NotRentExempt => {
                msg!("Error: Lamport balance below rent-exempt threshold")
            }
            ProverError::InsufficientFunds => msg!("Error: insufficient funds"),
            ProverError::InvalidMint => msg!("Error: Invalid Mint"),
            ProverError::MintMismatch => msg!("Error: Account not associated with this Mint"),
            ProverError::OwnerMismatch => msg!("Error: owner does not match"),
            ProverError::FixedSupply => msg!("Error: the total supply of this token is fixed"),
            ProverError::AlreadyInUse => msg!("Error: account or token already in use"),
            ProverError::InvalidNumberOfProvidedSigners => {
                msg!("Error: Invalid number of provided signers")
            }
            ProverError::InvalidNumberOfRequiredSigners => {
                msg!("Error: Invalid number of required signers")
            }
            ProverError::UninitializedState => msg!("Error: State is uninitialized"),
            ProverError::NativeNotSupported => {
                msg!("Error: Instruction does not support native tokens")
            }
            ProverError::NonNativeHasBalance => {
                msg!("Error: Non-native account can only be closed if its balance is zero")
            }
            ProverError::InvalidInstruction => msg!("Error: Invalid instruction"),
            ProverError::InvalidState => msg!("Error: Invalid account state for operation"),
            ProverError::Overflow => msg!("Error: Operation overflowed"),
            ProverError::AuthorityTypeNotSupported => {
                msg!("Error: Account does not support specified authority type")
            }
            ProverError::MintCannotFreeze => msg!("Error: This token mint cannot freeze accounts"),
            ProverError::AccountFrozen => msg!("Error: Account is frozen"),
            ProverError::MintDecimalsMismatch => {
                msg!("Error: decimals different from the Mint decimals")
            }
            ProverError::NonNativeNotSupported => {
                msg!("Error: Instruction does not support non-native tokens")
            }
        }
    }
}
