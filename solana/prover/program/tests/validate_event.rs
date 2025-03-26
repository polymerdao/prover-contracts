mod setup;

use {
    mollusk_svm::{result::Check, Mollusk},
    prover::{id, instruction},
};

#[test]
fn validate_event() {
    let mut mollusk = Mollusk::new(&id(), "prover");
    mollusk.compute_budget.compute_unit_limit = 500_000; // last known 2252

    let proof = setup::read_and_decode_proof_file("tests/op-proof-v2.hex")
        .expect("could not read proof file");

    mollusk.process_and_validate_instruction(
        &instruction::validate_event(&id(), proof).unwrap(),
        &[mollusk.sysvars.keyed_account_for_rent_sysvar()],
        &[Check::success()],
    );
}
