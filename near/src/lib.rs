// Copyright 2024, Polymer Labs
// Licensed under the Apache License, Version 2.0

use near_sdk::borsh::{self, BorshDeserialize, BorshSerialize};
use near_sdk::collections::UnorderedMap;
use near_sdk::serde::{Deserialize, Serialize};
use near_sdk::{env, near_bindgen, AccountId, PanicOnDefault};
use sha2::{Digest, Sha256};
use k256::ecdsa::{RecoveryId, Signature, VerifyingKey};

mod receipt_parser;
use receipt_parser::*;

/// Light client type enum
#[derive(BorshDeserialize, BorshSerialize, Serialize, Deserialize, Clone, Debug, PartialEq)]
#[serde(crate = "near_sdk::serde")]
pub enum LightClientType {
    SequencerLightClient,
}

/// Event log structure for validation results
#[derive(Serialize, Deserialize, Clone, Debug)]
#[serde(crate = "near_sdk::serde")]
pub struct ValidatedEvent {
    pub chain_id: u32,
    pub emitting_contract: String, // EVM address as hex string
    pub topics: Vec<u8>,
    pub unindexed_data: Vec<u8>,
}

/// Validated Solana logs structure
#[derive(Serialize, Deserialize, Clone, Debug)]
#[serde(crate = "near_sdk::serde")]
pub struct ValidatedSolLogs {
    pub chain_id: u32,
    pub program_id: [u8; 32],
    pub log_messages: Vec<String>,
}

/// Main CrossL2Prover contract
#[near_bindgen]
#[derive(BorshDeserialize, BorshSerialize, PanicOnDefault)]
pub struct CrossL2ProverV2 {
    /// The trusted sequencer public key (derived from Ethereum address)
    pub sequencer: [u8; 20], // Ethereum address (20 bytes)

    /// Chain ID of the L2 chain for which the sequencer signs over
    pub chain_id: [u8; 32],

    /// Client type identifier
    pub client_type: String,
}

#[near_bindgen]
impl CrossL2ProverV2 {
    /// Initialize the contract
    #[init]
    pub fn new(client_type: String, sequencer: String, chain_id: String) -> Self {
        // Parse sequencer address from hex string
        let sequencer_bytes = hex::decode(sequencer.trim_start_matches("0x"))
            .expect("Invalid sequencer address");
        assert_eq!(sequencer_bytes.len(), 20, "Sequencer address must be 20 bytes");

        let mut sequencer_array = [0u8; 20];
        sequencer_array.copy_from_slice(&sequencer_bytes);

        // Parse chain_id from hex string
        let chain_id_bytes = hex::decode(chain_id.trim_start_matches("0x"))
            .expect("Invalid chain ID");
        assert_eq!(chain_id_bytes.len(), 32, "Chain ID must be 32 bytes");

        let mut chain_id_array = [0u8; 32];
        chain_id_array.copy_from_slice(&chain_id_bytes);

        // Emit ping event
        env::log_str("Ping");

        Self {
            sequencer: sequencer_array,
            chain_id: chain_id_array,
            client_type,
        }
    }

    /// Get the light client type (constant)
    pub fn get_light_client_type(&self) -> LightClientType {
        LightClientType::SequencerLightClient
    }

    /// Get client type string
    pub fn get_client_type(&self) -> String {
        self.client_type.clone()
    }

    /// Emit a ping event for health checks
    pub fn ping(&self) {
        env::log_str("Ping");
    }

    /// Validates an event proof from Polymer's prove API for a non-Solana chain
    ///
    /// # Proof Format
    /// - state root (32 bytes) [0:32]
    /// - signature (65 bytes) [32:97]
    /// - source chain ID (big endian, 4 bytes) [97:101]
    /// - peptide height (big endian, 8 bytes) [101:109]
    /// - source chain block height (big endian, 8 bytes) [109:117]
    /// - receipt index (big endian, 4 bytes) [117:121]
    /// - event index (big endian, 4 bytes) [121:125]
    /// - number of topics (1 byte) [125]
    /// - event data end (big endian, 2 bytes) [126:128]
    /// - event emitter (contract address) (20 bytes) [128:148]
    /// - topics (32 bytes * number of topics)
    /// - event data (variable)
    /// - iavl proof (variable)
    pub fn validate_event(&self, proof: Vec<u8>) -> ValidatedEvent {
        assert!(proof.len() >= 148, "Proof too short");

        // Extract chain ID
        let chain_id = u32::from_be_bytes([
            proof[97], proof[98], proof[99], proof[100]
        ]);

        // Extract state root
        let mut state_root = [0u8; 32];
        state_root.copy_from_slice(&proof[0..32]);

        // Extract peptide height
        let peptide_height = u64::from_be_bytes([
            proof[101], proof[102], proof[103], proof[104],
            proof[105], proof[106], proof[107], proof[108]
        ]);

        // Extract signature components (v, r, s)
        let v = proof[96];
        let mut r = [0u8; 32];
        r.copy_from_slice(&proof[32..64]);
        let mut s = [0u8; 32];
        s.copy_from_slice(&proof[64..96]);

        // Verify sequencer signature
        self.verify_sequencer_signature(state_root, peptide_height, v, r, s);

        // Extract event data end position
        let event_end = u16::from_be_bytes([proof[126], proof[127]]) as usize;
        assert!(event_end <= proof.len(), "Invalid event end position");

        // Extract raw event
        let raw_event = &proof[128..event_end];

        // Extract block height, receipt index, and log index for IAVL proof
        let block_height = u64::from_be_bytes([
            proof[109], proof[110], proof[111], proof[112],
            proof[113], proof[114], proof[115], proof[116]
        ]);
        let receipt_index = u32::from_be_bytes([
            proof[117], proof[118], proof[119], proof[120]
        ]);
        let log_index = u32::from_be_bytes([
            proof[121], proof[122], proof[123], proof[124]
        ]);

        // Generate event root key
        let key = event_root_key(
            chain_id,
            &self.client_type,
            block_height,
            receipt_index,
            log_index
        );

        // Compute keccak256 hash of raw event
        let value = keccak256(raw_event);

        // Verify IAVL proof
        let iavl_proof = &proof[event_end..];
        verify_membership(state_root, &key, value, iavl_proof);

        // Parse event
        let num_topics = proof[125];
        let (emitting_contract, topics, unindexed_data) = parse_event(raw_event, num_topics);

        ValidatedEvent {
            chain_id,
            emitting_contract,
            topics,
            unindexed_data,
        }
    }

    /// Validates an event proof from Polymer's prove API for Solana chains
    ///
    /// # Proof Format
    /// - state root (32 bytes) [0:32]
    /// - signature (65 bytes) [32:97]
    /// - source chain ID (big endian, 4 bytes) [97:101]
    /// - peptide height (big endian, 8 bytes) [101:109]
    /// - source chain block height (big endian, 8 bytes) [109:117]
    /// - number of log messages (1 byte) [117]
    /// - txSignature (high) (32 bytes) [118:150]
    /// - txSignature (low) (32 bytes) [150:182]
    /// - programID (32 bytes) [182:214]
    /// - (currLogMsgDataEnd, logMsg) pairs (2 bytes + X bytes each)
    /// - iavl proof (variable)
    pub fn validate_sol_logs(&self, proof: Vec<u8>) -> ValidatedSolLogs {
        assert!(proof.len() >= 214, "Proof too short for Solana");

        // Extract chain ID
        let chain_id = u32::from_be_bytes([
            proof[97], proof[98], proof[99], proof[100]
        ]);

        // Extract state root
        let mut state_root = [0u8; 32];
        state_root.copy_from_slice(&proof[0..32]);

        // Extract peptide height
        let peptide_height = u64::from_be_bytes([
            proof[101], proof[102], proof[103], proof[104],
            proof[105], proof[106], proof[107], proof[108]
        ]);

        // Extract signature components
        let v = proof[96];
        let mut r = [0u8; 32];
        r.copy_from_slice(&proof[32..64]);
        let mut s = [0u8; 32];
        s.copy_from_slice(&proof[64..96]);

        // Verify sequencer signature
        self.verify_sequencer_signature(state_root, peptide_height, v, r, s);

        // Extract program ID
        let mut program_id = [0u8; 32];
        program_id.copy_from_slice(&proof[182..214]);

        // Extract log messages
        let num_log_messages = proof[117] as usize;
        let mut log_messages = Vec::new();
        let mut curr_log_message_start = 214;
        let mut current_log_message_end = 214;

        for _ in 0..num_log_messages {
            current_log_message_end = u16::from_be_bytes([
                proof[curr_log_message_start],
                proof[curr_log_message_start + 1]
            ]) as usize;

            let log_msg = String::from_utf8(
                proof[curr_log_message_start + 2..current_log_message_end].to_vec()
            ).expect("Invalid UTF-8 in log message");

            log_messages.push(log_msg);
            curr_log_message_start = current_log_message_end;
        }

        // Construct raw event for hashing
        let mut raw_event = program_id.to_vec();
        for msg in &log_messages {
            raw_event.extend_from_slice(msg.as_bytes());
        }

        // Extract tx signature components
        let mut tx_sig_high = [0u8; 32];
        tx_sig_high.copy_from_slice(&proof[118..150]);
        let mut tx_sig_low = [0u8; 32];
        tx_sig_low.copy_from_slice(&proof[150..182]);

        // Extract block height
        let block_height = u64::from_be_bytes([
            proof[109], proof[110], proof[111], proof[112],
            proof[113], proof[114], proof[115], proof[116]
        ]);

        // Generate Solana event root key
        let key = solana_event_root_key(
            chain_id,
            &self.client_type,
            block_height,
            tx_sig_high,
            tx_sig_low,
            program_id
        );

        // Compute keccak256 hash of raw event
        let value = keccak256(&raw_event);

        // Verify IAVL proof
        let iavl_proof = &proof[current_log_message_end..];
        verify_membership(state_root, &key, value, iavl_proof);

        ValidatedSolLogs {
            chain_id,
            program_id,
            log_messages,
        }
    }

    /// Inspect log identifier from proof
    pub fn inspect_log_identifier(&self, proof: Vec<u8>) -> (u32, u64, u32, u32) {
        assert!(proof.len() >= 125, "Proof too short");

        let src_chain = u32::from_be_bytes([
            proof[97], proof[98], proof[99], proof[100]
        ]);
        let block_number = u64::from_be_bytes([
            proof[109], proof[110], proof[111], proof[112],
            proof[113], proof[114], proof[115], proof[116]
        ]);
        let receipt_index = u32::from_be_bytes([
            proof[117], proof[118], proof[119], proof[120]
        ]);
        let log_index = u32::from_be_bytes([
            proof[121], proof[122], proof[123], proof[124]
        ]);

        (src_chain, block_number, receipt_index, log_index)
    }

    /// Inspect Polymer state from proof
    pub fn inspect_polymer_state(&self, proof: Vec<u8>) -> ([u8; 32], u64, Vec<u8>) {
        assert!(proof.len() >= 109, "Proof too short");

        let mut state_root = [0u8; 32];
        state_root.copy_from_slice(&proof[0..32]);

        let height = u64::from_be_bytes([
            proof[101], proof[102], proof[103], proof[104],
            proof[105], proof[106], proof[107], proof[108]
        ]);

        let signature = proof[32..97].to_vec();

        (state_root, height, signature)
    }

    /// Internal: Verify sequencer signature over apphash
    fn verify_sequencer_signature(
        &self,
        app_hash: [u8; 32],
        peptide_height: u64,
        v: u8,
        r: [u8; 32],
        s: [u8; 32],
    ) {
        // Construct the message that was signed
        // keccak256(bytes.concat(bytes32(0), CHAIN_ID, keccak256(abi.encodePacked(appHash, peptideHeight))))

        // First: keccak256(abi.encodePacked(appHash, peptideHeight))
        let mut inner_data = Vec::new();
        inner_data.extend_from_slice(&app_hash);
        inner_data.extend_from_slice(&peptide_height.to_be_bytes());
        let inner_hash = keccak256(&inner_data);

        // Second: keccak256(bytes.concat(bytes32(0), CHAIN_ID, inner_hash))
        let mut outer_data = Vec::new();
        outer_data.extend_from_slice(&[0u8; 32]); // bytes32(0)
        outer_data.extend_from_slice(&self.chain_id);
        outer_data.extend_from_slice(&inner_hash);
        let message_hash = keccak256(&outer_data);

        // Recover the signer from the signature
        let recovered_address = recover_address(&message_hash, v, &r, &s);

        // Verify it matches the sequencer address
        assert_eq!(
            recovered_address, self.sequencer,
            "Invalid sequencer signature"
        );
    }
}

/// Helper function to compute keccak256 hash
fn keccak256(data: &[u8]) -> [u8; 32] {
    use sha3::{Keccak256, Digest};
    let mut hasher = Keccak256::new();
    hasher.update(data);
    let result = hasher.finalize();
    let mut hash = [0u8; 32];
    hash.copy_from_slice(&result);
    hash
}

/// Recover Ethereum address from ECDSA signature
fn recover_address(message_hash: &[u8; 32], v: u8, r: &[u8; 32], s: &[u8; 32]) -> [u8; 20] {
    use k256::ecdsa::{RecoveryId, Signature, VerifyingKey};

    // Construct signature from r and s
    let mut sig_bytes = [0u8; 64];
    sig_bytes[..32].copy_from_slice(r);
    sig_bytes[32..].copy_from_slice(s);

    let signature = Signature::from_bytes(&sig_bytes.into())
        .expect("Invalid signature");

    // Recovery ID is v - 27 for Ethereum
    let recovery_id = RecoveryId::from_byte(v - 27)
        .expect("Invalid recovery ID");

    // Recover the public key
    let recovered_key = VerifyingKey::recover_from_prehash(
        message_hash,
        &signature,
        recovery_id
    ).expect("Failed to recover public key");

    // Convert public key to Ethereum address
    // Ethereum address is the last 20 bytes of keccak256(public_key)
    let pub_key_bytes = recovered_key.to_encoded_point(false);
    let pub_key_uncompressed = pub_key_bytes.as_bytes();
    // Skip the first byte (0x04 prefix for uncompressed key)
    let pub_key_hash = keccak256(&pub_key_uncompressed[1..]);

    let mut address = [0u8; 20];
    address.copy_from_slice(&pub_key_hash[12..]);
    address
}

#[cfg(test)]
mod tests {
    use super::*;
    use near_sdk::test_utils::{VMContextBuilder, accounts};
    use near_sdk::testing_env;

    #[test]
    fn test_initialization() {
        let context = VMContextBuilder::new()
            .predecessor_account_id(accounts(0))
            .build();
        testing_env!(context);

        let contract = CrossL2ProverV2::new(
            "test-client".to_string(),
            "0x1234567890123456789012345678901234567890".to_string(),
            "0x0000000000000000000000000000000000000000000000000000000000000001".to_string(),
        );

        assert_eq!(contract.get_client_type(), "test-client");
        assert_eq!(
            contract.get_light_client_type(),
            LightClientType::SequencerLightClient
        );
    }
}
