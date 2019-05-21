const Migrations = artifacts.require("Migrations");
const PODEX = artifacts.require("PODEX");

module.exports = function(deployer) {
  deployer.deploy(Migrations);
  deployer.deploy(PODEX);
};
