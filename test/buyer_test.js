const PODEX = artifacts.require("PODEX");
const { getWeb3 } = require("./helpers");
const web3 = getWeb3();
const truffleAssert = require('truffle-assertions');


contract("Buyer", async (accounts) => {
    let podex;

    let seller = accounts[0];
    let buyer = accounts[1];

    before(async () => {
        podex = await PODEX.deployed();
    });

    it("should buyerDeposit correctly", async () => {
        let _value = web3.utils.toWei('3', 'ether');
        
        await podex.buyerDeposit(seller, {from: buyer, value: _value});

        let bdpsitAfter = await podex.buyerDeposits_(buyer, seller);
        assert.equal(bdpsitAfter.value, _value, "wrong deposit value");
    });

    it("should buyerUnDeposit correctly", async () => {
        await podex.buyerUnDeposit(seller, {from: buyer});

        let bdpsitAfter = await podex.buyerDeposits_(buyer, seller);
        assert.equal(bdpsitAfter.stat, 1, "wrong deposit stat");

        await truffleAssert.reverts(
            podex.buyerUnDeposit(seller, {from: buyer}),
            "wrong stat"
        )
    });

    it("should not buyerUnDeposit repeatedly", async () => {
        await truffleAssert.reverts(
            podex.buyerUnDeposit(seller, {from: buyer}),
            "wrong stat"
        )
    });



});
