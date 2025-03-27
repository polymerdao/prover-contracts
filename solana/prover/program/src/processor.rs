//! Program state processor

use {
    crate::instruction::ProverInstruction,
    hex,
    sha3::{Digest, Keccak256},
    solana_account_info::AccountInfo,
    solana_msg::msg,
    solana_program::{keccak, secp256k1_recover::secp256k1_recover},
    solana_program_error::{ProgramError, ProgramResult},
    solana_pubkey::Pubkey,
};

/// Program state handler.
pub struct Processor {}
impl Processor {
    fn verify_sequencer_signature(
        app_hash: &[u8; 32],
        peptide_height: &[u8; 8],
        signature: &[u8; 64],
        recovery_id: u8,
    ) -> ProgramResult {
        let chain_id: [u8; 32] = [
            0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
            0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
            0x00, 0x00, 0x03, 0x85,
        ];

        msg!("app_hash {:?}", hex::encode(app_hash).as_str());
        msg!("peptide_height {:?}", u64::from_be_bytes(*peptide_height));

        let message_hash = {
            let mut hasher = keccak::Hasher::default();
            hasher.hash(app_hash);
            hasher.hash(peptide_height);
            hasher.result()
        };
        msg!("message_hash {:?}", hex::encode(message_hash.0).as_str());

        let hash = {
            let mut hasher = keccak::Hasher::default();
            hasher.hash(&[0; 32]);
            hasher.hash(&chain_id);
            hasher.hash(&message_hash.0);
            hasher.result()
        };
        // keccak256(bytes.concat(bytes32(0), CHAIN_ID, keccak256(abi.encodePacked(appHash, peptideHeight)))),
        msg!("hash {:?}", hex::encode(hash.0).as_str());

        msg!("recovery_id {}", recovery_id);
        msg!("signature {:?}", hex::encode(signature).as_str());
        let recovered_pubkey =
            secp256k1_recover(&hash.0, recovery_id - 27, signature).map_err(|e| {
                msg!("error {}", u64::from(e));
                ProgramError::InvalidArgument
            })?;

        let eth_address_hash = Keccak256::digest(recovered_pubkey.to_bytes());
        let eth_address = &eth_address_hash[12..32]; // Take last 20 bytes

        msg!(format!(
            "Recovered Ethereum Address: 0x{:?}",
            hex::encode(eth_address)
        )
        .as_str());

        Ok(())
    }

    /// Processes an [`InitializeMint`](enum.TokenInstruction.html) instruction.
    pub fn process_validate_event(proof: Vec<u8>) -> ProgramResult {
        Self::verify_sequencer_signature(
            &<[u8; 32]>::try_from(&proof[0..32]).unwrap(),
            &<[u8; 8]>::try_from(&proof[101..109]).unwrap(),
            &<[u8; 64]>::try_from(&proof[32..96]).unwrap(),
            proof[96],
        )
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
