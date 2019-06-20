const zkPoDExchange = artifacts.require("zkPoDExchange");
const { getWeb3 } = require("./helpers");
const web3 = getWeb3();
const truffleAssert = require('truffle-assertions');


contract("Bob", async (accounts) => {
    let zkPoDEX;

    let alice = accounts[0];
    let bob = accounts[1];

    before(async () => {
        zkPoDEX = await zkPoDExchange.deployed();
    });

    it("should bobDeposit correctly", async () => {
        let _value = web3.utils.toWei('3', 'ether');
        
        await zkPoDEX.bobDeposit(alice, {from: bob, value: _value});

        let bdpsitAfter = await zkPoDEX.bobDeposits_(bob, alice);
        assert.equal(bdpsitAfter.value, _value, "wrong deposit value");
    });

    it("should bobUnDeposit correctly", async () => {
        await zkPoDEX.bobUnDeposit(alice, {from: bob});

        let bdpsitAfter = await zkPoDEX.bobDeposits_(bob, alice);
        assert.equal(bdpsitAfter.stat, 1, "wrong deposit stat");

        await truffleAssert.reverts(
            zkPoDEX.bobUnDeposit(alice, {from: bob}),
            "wrong stat"
        )
    });

    it("should not bobUnDeposit repeatedly", async () => {
        await truffleAssert.reverts(
            zkPoDEX.bobUnDeposit(alice, {from: bob}),
            "wrong stat"
        )
    });



});
