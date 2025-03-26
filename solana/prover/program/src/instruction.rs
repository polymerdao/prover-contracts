//! Instruction types

use {
    crate::error::ProverError, solana_instruction::Instruction, solana_program_error::ProgramError,
    solana_pubkey::Pubkey, std::mem::size_of,
};

/// Instructions supported by the prover program.
#[repr(C)]
#[derive(Clone, Debug, PartialEq)]
pub enum ProverInstruction {
    /// Validates an event encoded in the provided proof
    ValidateEvent {
        /// Binary encoded proof from the prove-api
        proof: Vec<u8>,
    },
}

impl ProverInstruction {
    /// Unpacks a byte buffer into a `ProverInstruction`
    pub fn unpack(input: &[u8]) -> Result<Self, ProgramError> {
        use ProverError::InvalidInstruction;

        let (&tag, rest) = input.split_first().ok_or(InvalidInstruction)?;
        Ok(match tag {
            0 => {
                let proof: Vec<u8> = Vec::from(rest);
                Self::ValidateEvent { proof }
            }
            _ => return Err(ProverError::InvalidInstruction.into()),
        })
    }

    /// Packs a `ProverInstruction` into a byte buffer
    pub fn pack(&self) -> Vec<u8> {
        let mut buf = Vec::with_capacity(size_of::<Self>());
        match self {
            &Self::ValidateEvent { ref proof } => {
                buf.push(0);
                buf.extend_from_slice(proof.as_slice());
            }
        };
        buf
    }
}

/// Creates a `ValidateEvent` instruction
pub fn validate_event(program_id: &Pubkey, proof: Vec<u8>) -> Result<Instruction, ProgramError> {
    let data = ProverInstruction::ValidateEvent { proof }.pack();
    Ok(Instruction {
        program_id: *program_id,
        accounts: vec![],
        data,
    })
}

#[cfg(test)]
mod test {
    use {super::*, proptest::prelude::*};

    #[test]
    fn test_instruction_packing() {
        let proof = vec![1, 2, 3, 4, 5];
        let check = ProverInstruction::ValidateEvent {
            proof: Vec::from(proof.as_slice()),
        };
        let packed = check.pack();
        let mut expect = Vec::from([0u8]);
        expect.extend_from_slice(&proof);
        assert_eq!(packed, expect);
        let unpacked = ProverInstruction::unpack(&expect).unwrap();
        assert_eq!(unpacked, check);
    }

    #[test]
    fn test_instruction_unpack_panic() {
        for i in 0..255u8 {
            for j in 1..10 {
                let mut data = vec![0; j];
                data[0] = i;
                let _no_panic = ProverInstruction::unpack(&data);
            }
        }
    }

    proptest! {
        #![proptest_config(ProptestConfig::with_cases(1024))]
        #[test]
        fn test_instruction_unpack_proptest(
            data in prop::collection::vec(any::<u8>(), 0..255)
        ) {
            let _no_panic = ProverInstruction::unpack(&data);
        }
    }
}
