var CrossL2Prover = artifacts.require("CrossL2ProverV2");

module.exports = function(deployer) {
  deployer.deploy(CrossL2Prover,"proof_api" , "0x8d3921b96a3815f403fb3a4c7ff525969d16f9e0", "0x0000000000000000000000000000000000000000000000000000000000000385");
};
