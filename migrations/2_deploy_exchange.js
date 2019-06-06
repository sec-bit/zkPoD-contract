const ECDSA = artifacts.require("ECDSA");
const PublicVar = artifacts.require("PublicVar");
const zkPoDExchange = artifacts.require("zkPoDExchange");

module.exports = function(deployer) {
  deployer.deploy(ECDSA);
  deployer.link(ECDSA, zkPoDExchange);
  deployer.deploy(PublicVar).then(function() {
    return deployer.deploy(zkPoDExchange, PublicVar.address);
  });
};
