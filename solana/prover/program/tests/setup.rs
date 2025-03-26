use hex;
use std::fs::File;
use std::io::Read;

pub fn read_and_decode_proof_file(file_path: &str) -> Result<Vec<u8>, Box<dyn std::error::Error>> {
    // Open the file
    let mut file = File::open(file_path)?;

    // Read the file contents into a string
    let mut contents = String::new();
    file.read_to_string(&mut contents)?;

    // Hex decode the contents
    let decoded = hex::decode(&contents.trim().as_bytes()[2..])?;

    Ok(decoded)
}
