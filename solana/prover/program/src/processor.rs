//! Program state processor

use {
    crate::instruction::ProverInstruction, solana_account_info::AccountInfo, solana_msg::msg,
    solana_program_error::ProgramResult, solana_pubkey::Pubkey,
};

/// Program state handler.
pub struct Processor {}
impl Processor {
    /// Processes an [`InitializeMint`](enum.TokenInstruction.html) instruction.
    pub fn process_validate_event(_proof: Vec<u8>) -> ProgramResult {
        // TODO call prover here
        Ok(())
    }

    /// Processes an [`Instruction`](enum.Instruction.html).
    pub fn process(_program_id: &Pubkey, _accounts: &[AccountInfo], input: &[u8]) -> ProgramResult {
        let instruction = ProverInstruction::unpack(input)?;

        match instruction {
            ProverInstruction::ValidateEvent { proof } => {
                msg!("Instruction: ValidateEvent");
                Self::process_validate_event(proof)
            }
        }
    }
}
