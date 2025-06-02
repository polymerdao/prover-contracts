var CrossL2ProverV2 = artifacts.require("CrossL2ProverV2");

module.exports = function(deployer) {
  deployer.deploy(CrossL2ProverV2,"proof_api" , "0x4632e987c6f31d6351e39324dec6f1404af56209", "0x00000000000000000000000000000000000000000000000000000000000001bc");
};
