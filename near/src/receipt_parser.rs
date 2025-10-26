// Copyright 2024, Polymer Labs
// Licensed under the Apache License, Version 2.0

use sha2::{Digest, Sha256};

/// Parse an event from raw bytes
///
/// # Arguments
/// * `raw_event` - Raw event data containing emitting contract, topics, and unindexed data
/// * `num_topics` - Number of 32-byte topics in the event
///
/// # Returns
/// Tuple of (emitting_contract as hex string, topics bytes, unindexed_data bytes)
pub fn parse_event(raw_event: &[u8], num_topics: u8) -> (String, Vec<u8>, Vec<u8>) {
    assert!(raw_event.len() >= 20, "Raw event too short");

    // Extract emitting contract (first 20 bytes as Ethereum address)
    let emitting_contract = &raw_event[0..20];
    let emitting_contract_hex = format!("0x{}", hex::encode(emitting_contract));

    // Calculate topics end position (20 bytes for address + 32 bytes per topic)
    let topics_end = 20 + (32 * num_topics as usize);
    assert!(raw_event.len() >= topics_end, "Raw event too short for topics");

    // Extract topics
    let topics = raw_event[20..topics_end].to_vec();

    // Extract unindexed data
    let unindexed_data = if raw_event.len() > topics_end {
        raw_event[topics_end..].to_vec()
    } else {
        Vec::new()
    };

    (emitting_contract_hex, topics, unindexed_data)
}

/// Generate the event root key for IAVL proof verification
///
/// Format: "chain/{chainId}/storedLogs/{clientType}/{height}/{receiptIndex}/{logIndex}"
pub fn event_root_key(
    chain_id: u32,
    client_type: &str,
    height: u64,
    receipt_index: u32,
    log_index: u32,
) -> Vec<u8> {
    format!(
        "chain/{}/storedLogs/{}/{}/{}/{}",
        chain_id, client_type, height, receipt_index, log_index
    )
    .into_bytes()
}

/// Generate the Solana event root key for IAVL proof verification
///
/// Format: "chain/{chainId}/storedLogs/{clientType}/{height}/{txSignatureHigh}{txSignatureLow}/{programID}"
pub fn solana_event_root_key(
    chain_id: u32,
    client_type: &str,
    height: u64,
    tx_signature_high: [u8; 32],
    tx_signature_low: [u8; 32],
    program_id: [u8; 32],
) -> Vec<u8> {
    format!(
        "chain/{}/storedLogs/{}/{}/{}{}/{}",
        chain_id,
        client_type,
        height,
        hex::encode(tx_signature_high),
        hex::encode(tx_signature_low),
        hex::encode(program_id)
    )
    .into_bytes()
}

/// Verify IAVL membership proof
///
/// This verifies that a given key-value pair exists in an IAVL tree with the given root.
///
/// # Proof Format
/// ```text
/// +----------------------------------------------------------------------------------------------------+
/// | header:  | number of paths (1B) | path-0 start (1B) | prefix... | varint(len(key))                |
/// +----------------------------------------------------------------------------------------------------+
/// | path-0:  | path-0 suffix start (1B) | path-0 suffix end (1B) | path-0 prefix... | path-0 suffix... |
/// +----------------------------------------------------------------------------------------------------+
/// | ...      |                                        ...                                             |
/// +----------------------------------------------------------------------------------------------------+
/// | path-n:  | path-n suffix start (1B) | path-n suffix end (1B) | path-n prefix... | path-n suffix... |
/// +----------------------------------------------------------------------------------------------------+
/// ```
///
/// # Arguments
/// * `root` - The IAVL tree root hash to verify against
/// * `key` - The key to verify in the tree
/// * `value` - The value hash to verify (keccak256 of the actual value)
/// * `proof` - The IAVL proof bytes
///
/// # Panics
/// Panics if the proof is invalid or doesn't match the root
pub fn verify_membership(root: [u8; 32], key: &[u8], value: [u8; 32], proof: &[u8]) {
    assert!(proof.len() >= 2, "Proof too short");

    let num_paths = proof[0] as usize;
    let path0_start = proof[1] as usize;

    assert!(path0_start <= proof.len(), "Invalid path0 start");

    // Initial prehash: sha256(prefix || key || 0x20 || sha256(value))
    let mut prehash_data = Vec::new();
    prehash_data.extend_from_slice(&proof[2..path0_start]);
    prehash_data.extend_from_slice(key);
    prehash_data.push(0x20); // Length prefix for value hash (32 bytes)

    // Hash the value
    let value_hash = Sha256::digest(&value);
    prehash_data.extend_from_slice(&value_hash);

    let mut prehash = Sha256::digest(&prehash_data);
    let mut offset = path0_start;

    // Iterate through each path in the proof
    for _ in 0..num_paths {
        assert!(offset + 2 <= proof.len(), "Invalid proof structure");

        let suffix_start = proof[offset] as usize;
        let suffix_end = proof[offset + 1] as usize;

        assert!(offset + suffix_end <= proof.len(), "Invalid suffix bounds");

        // Hash: sha256(prefix || prehash || suffix)
        let mut hash_data = Vec::new();
        hash_data.extend_from_slice(&proof[offset + 2..offset + suffix_start]);
        hash_data.extend_from_slice(&prehash);
        hash_data.extend_from_slice(&proof[offset + suffix_start..offset + suffix_end]);

        prehash = Sha256::digest(&hash_data);
        offset += suffix_end;
    }

    // Verify the final hash matches the root
    let mut prehash_array = [0u8; 32];
    prehash_array.copy_from_slice(&prehash);

    assert_eq!(prehash_array, root, "Invalid proof root");
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_parse_event() {
        // Mock event: 20-byte address + 2 topics (32 bytes each) + 10 bytes unindexed data
        let mut raw_event = vec![0x12; 20]; // address
        raw_event.extend_from_slice(&[0x34; 32]); // topic 1
        raw_event.extend_from_slice(&[0x56; 32]); // topic 2
        raw_event.extend_from_slice(&[0x78; 10]); // unindexed data

        let (address, topics, unindexed) = parse_event(&raw_event, 2);

        assert_eq!(address, "0x1212121212121212121212121212121212121212");
        assert_eq!(topics.len(), 64); // 2 topics * 32 bytes
        assert_eq!(unindexed.len(), 10);
    }

    #[test]
    fn test_event_root_key() {
        let key = event_root_key(1, "test-client", 100, 5, 2);
        let key_str = String::from_utf8(key).unwrap();
        assert_eq!(key_str, "chain/1/storedLogs/test-client/100/5/2");
    }

    #[test]
    fn test_solana_event_root_key() {
        let tx_high = [0x11; 32];
        let tx_low = [0x22; 32];
        let program_id = [0x33; 32];

        let key = solana_event_root_key(2, "solana-client", 200, tx_high, tx_low, program_id);
        let key_str = String::from_utf8(key).unwrap();

        assert!(key_str.starts_with("chain/2/storedLogs/solana-client/200/"));
        assert!(key_str.contains(&hex::encode(tx_high)));
        assert!(key_str.contains(&hex::encode(tx_low)));
        assert!(key_str.contains(&hex::encode(program_id)));
    }
}
