mod setup;

use {
    mollusk_svm::{result::Check, Mollusk},
    prover::{id, instruction},
};

#[test]
fn initialize_mint() {
    let mut mollusk = Mollusk::new(&id(), "prover");
    mollusk.compute_budget.compute_unit_limit = 5_000; // last known 2252

    mollusk.process_and_validate_instruction(
        &instruction::validate_event(&id(), vec![1, 1]).unwrap(),
        &[mollusk.sysvars.keyed_account_for_rent_sysvar()],
        &[Check::success()],
    );
}
