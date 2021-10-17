const SimpleStRiddlerorage = artifacts.require("Riddler");

module.exports = function(deployer) {
  deployer.deploy(Riddler);
};
