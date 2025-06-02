var CrossL2ProverV2 = artifacts.require("CrossL2ProverV2");

module.exports = function(deployer) {
  deployer.deploy(CrossL2ProverV2,"proof_api" , "0xBbba797f031f630Ba321F042a9c89F077BCb9703", "0x0000000000000000000000000000000000000000000000000000000001bc1bc1");
};
