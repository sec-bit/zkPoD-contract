const ECDSA = artifacts.require("ECDSA");
const PublicVar = artifacts.require("PublicVar");
const PODEX = artifacts.require("PODEX");

module.exports = function(deployer) {
  deployer.deploy(ECDSA);
  deployer.link(ECDSA, PODEX);
  deployer.deploy(PublicVar).then(function() {
    return deployer.deploy(PODEX, PublicVar.address);
  });
};
