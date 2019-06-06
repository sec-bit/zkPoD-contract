const zkPoDExchange = artifacts.require("zkPoDExchange");
const { getWeb3 } = require("./helpers");
const web3 = getWeb3();
const truffleAssert = require('truffle-assertions');


contract("Buyer", async (accounts) => {
    let zkPoDEX;

    let seller = accounts[0];
    let buyer = accounts[1];

    before(async () => {
        zkPoDEX = await zkPoDExchange.deployed();
    });

    it("should buyerDeposit correctly", async () => {
        let _value = web3.utils.toWei('3', 'ether');
        
        await zkPoDEX.buyerDeposit(seller, {from: buyer, value: _value});

        let bdpsitAfter = await zkPoDEX.buyerDeposits_(buyer, seller);
        assert.equal(bdpsitAfter.value, _value, "wrong deposit value");
    });

    it("should buyerUnDeposit correctly", async () => {
        await zkPoDEX.buyerUnDeposit(seller, {from: buyer});

        let bdpsitAfter = await zkPoDEX.buyerDeposits_(buyer, seller);
        assert.equal(bdpsitAfter.stat, 1, "wrong deposit stat");

        await truffleAssert.reverts(
            zkPoDEX.buyerUnDeposit(seller, {from: buyer}),
            "wrong stat"
        )
    });

    it("should not buyerUnDeposit repeatedly", async () => {
        await truffleAssert.reverts(
            zkPoDEX.buyerUnDeposit(seller, {from: buyer}),
            "wrong stat"
        )
    });



});
