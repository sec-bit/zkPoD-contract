const zkPoDExchange = artifacts.require("zkPoDExchange");
const { getWeb3 } = require("./helpers");
const web3 = getWeb3();
const fs = require('fs');
const testdataPath = "test/testdata";
const truffleAssert = require('truffle-assertions');

contract("zkPoDExchange", async (accounts) => {

    let zkPoDEX;

    let alice = accounts[0];
    let bob = accounts[1];

    before(async () => {
        zkPoDEX = await zkPoDExchange.deployed();
    });

    it("should keep basic settings", async () => {
        let s = await zkPoDEX.s_();
        let t1 = await zkPoDEX.t1();
        // let t2 = await zkPoDEX.t2();
        let t3 = await zkPoDEX.t3();
        let t4 = await zkPoDEX.t4()
        assert.equal(s, 65, "max s count changed");
        assert.equal(t1, 8 * 60 * 60, "t1 changed");
        // assert.equal(t2, 12 * 60 * 60, "t2 changed");
        assert.equal(t3, 24 * 60 * 60, "t3 changed");
        assert.equal(t4, 3 * 24 * 60 * 60, "t3 changed");
    });

    it("should publish plain correctly", async () => {
        let blt = JSON.parse(fs.readFileSync(testdataPath + "/bulletin.plain.json"));

        let _size = blt.size;
        let _s = blt.s;
        let _n = blt.n;
        let _sigma_mkl_root = "0x" + blt.sigma_mkl_root;
        let _blt_type = 0;

        await zkPoDEX.publish(_size, _s, _n, _sigma_mkl_root, 0, _blt_type);

        let _bltKey = web3.utils.soliditySha3({ t: 'uint64', v: _size }, { t: 'uint64', v: _s }, { t: 'uint64', v: _n }, { t: 'uint256', v: _sigma_mkl_root });
        let bulletin = await zkPoDEX.bulletins_(_bltKey);

        assert.equal(bulletin.owner, alice, "wrong owner");
        assert.equal(bulletin.size, blt.size, "wrong size");
        assert.equal(bulletin.s, blt.s, "wrong s");
        assert.equal(bulletin.n, blt.n, "wrong n");
        assert.equal(bulletin.sigma_mkl_root, web3.utils.toBN(_sigma_mkl_root).toString(), "wrong sigma_mkl_root");
        assert.equal(bulletin.stat, 0, "wrong status");
        assert.equal(bulletin.blt_type, _blt_type, "wrong blt_type");

        await truffleAssert.reverts(
            zkPoDEX.publish(_size, _s, _n, _sigma_mkl_root, 0, _blt_type),
            "blt occupied"
        );

    });

    it("should publish table correctly", async () => {
        let blt = JSON.parse(fs.readFileSync(testdataPath + "/bulletin.table.json"));

        let _size = 0;
        let _s = blt.s;
        let _n = blt.n;
        let _sigma_mkl_root = "0x" + blt.sigma_mkl_root;
        let _vrf_meta_digest = "0x" + blt.vrf_meta_digest;
        let _blt_type = 1;

        await zkPoDEX.publish(_size, _s, _n, _sigma_mkl_root, _vrf_meta_digest, _blt_type);

        let _bltKey = web3.utils.soliditySha3({ t: 'uint64', v: _s }, { t: 'uint64', v: _n }, { t: 'uint256', v: _sigma_mkl_root }, { t: 'uint256', v: _vrf_meta_digest });
        let bulletin = await zkPoDEX.bulletins_(_bltKey);

        assert.equal(bulletin.owner, alice, "wrong owner");
        assert.equal(bulletin.s, blt.s, "wrong s");
        assert.equal(bulletin.n, blt.n, "wrong n");
        assert.equal(bulletin.sigma_mkl_root, web3.utils.toBN(_sigma_mkl_root).toString(), "wrong sigma_mkl_root");
        assert.equal(bulletin.vrf_meta_digest, web3.utils.toBN(_vrf_meta_digest).toString(), "wrong vrf_meta_digest");
        assert.equal(bulletin.stat, 0, "wrong status");
        assert.equal(bulletin.blt_type, _blt_type, "wrong blt_type");

        await truffleAssert.reverts(
            zkPoDEX.publish(_size, _s, _n, _sigma_mkl_root, _vrf_meta_digest, _blt_type),
            "blt occupied"
        );
    });

    it("should submitProofComplaint correctly (for not evil)", async () => {
        let path = testdataPath + "/complaint/not_evil";
        let receipt = JSON.parse(fs.readFileSync(path + "/receipt"));
        let secret = JSON.parse(fs.readFileSync(path + "/secret"));

        // let _sessionId = Math.random().toString().slice(2, 17);
        let _sessionId = 0;
        let _from = bob;
        let _seed0 = "0x" + secret.s;
        let _seed2 = "0x" + receipt.s;
        let _k_mkl_root = "0x" + receipt.k;
        let _count = receipt.c;
        let _price = web3.utils.toWei('0.5', 'ether');
        let _expireAt = Math.floor(Date.now() / 1000) + 3600;

        await zkPoDEX.bobDeposit(alice, {from: bob, value: _price});

        let hash = web3.utils.soliditySha3({ t: 'uint256', v: _sessionId }, { t: 'address', v: _from }, { t: 'bytes32', v: _seed2 }, { t: 'bytes32', v: _k_mkl_root }, { t: 'uint64', v: _count }, { t: 'uint256', v: _price }, { t: 'uint256', v: _expireAt });
        let signature = await web3.eth.sign(hash, bob);
        // console.log("signature:", signature);

        // before submit
        let bdpsitBefore = await zkPoDEX.bobDeposits_(bob, alice);

        // session record
        let _result = await zkPoDEX.submitProofComplaint(_seed0, _sessionId, _from, _seed2, _k_mkl_root, _count, _price, _expireAt, signature, {from: alice});
        let _txblock = await web3.eth.getBlock(_result.receipt.blockNumber);
        let _sRecord = await zkPoDEX.getSessionRecord(alice, bob, _sessionId);
        assert.equal(_sRecord.submitAt.toNumber(), _txblock.timestamp, "wrong time");
        assert.equal(_sRecord.mode, 0, "wrong mode");
        assert.equal(_sRecord.stat, 1, "wrong stat");

        // event
        await truffleAssert.eventEmitted(_result, 'OnComplaintKey', { _a: alice, _b: bob });

        // bobDeposits_
        let bdpsitAfter = await zkPoDEX.bobDeposits_(bob, alice);
        assert.equal(bdpsitAfter.pendingCnt, bdpsitBefore.pendingCnt.toNumber()+1, "wrong pending cnt");

        await truffleAssert.reverts(
            zkPoDEX.submitProofComplaint(_seed0, _sessionId, bob, _seed2, _k_mkl_root, _count, _price, _expireAt, signature, {from: alice})
        );
    });


    it("should submitProofComplaint correctly (for evil)", async () => {
        let path = testdataPath + "/complaint/evil";
        let receipt = JSON.parse(fs.readFileSync(path + "/receipt"));
        let secret = JSON.parse(fs.readFileSync(path + "/secret"));

        // let _sessionId = Math.random().toString().slice(2, 17);
        let _sessionId = 1;
        let _from = bob;
        let _seed0 = "0x" + secret.s;
        let _seed2 = "0x" + receipt.s;
        let _k_mkl_root = "0x" + receipt.k;
        let _count = receipt.c;
        let _price = web3.utils.toWei('0.5', 'ether');
        let _expireAt = Math.floor(Date.now() / 1000) + 3600;

        await zkPoDEX.bobDeposit(alice, {from: bob, value: _price});

        let hash = web3.utils.soliditySha3({ t: 'uint256', v: _sessionId }, { t: 'address', v: _from }, { t: 'bytes32', v: _seed2 }, { t: 'bytes32', v: _k_mkl_root }, { t: 'uint64', v: _count }, { t: 'uint256', v: _price }, { t: 'uint256', v: _expireAt });
        let signature = await web3.eth.sign(hash, bob);

        await zkPoDEX.submitProofComplaint(_seed0, _sessionId, _from, _seed2, _k_mkl_root, _count, _price, _expireAt, signature, { from: alice });


        await truffleAssert.reverts(
            zkPoDEX.submitProofComplaint(_seed0, _sessionId, _from, _seed2, _k_mkl_root, _count, _price, _expireAt, signature, { from: alice }),
            "not new"
        )
    });

    it("should claimComplaint correctly (for evil)", async () => {
        let path = testdataPath + "/complaint/evil";
        let bulletin = JSON.parse(fs.readFileSync(testdataPath + "/bulletin.plain.json"));
        let claim = JSON.parse(fs.readFileSync(path + "/claim"));
        let _sessionId = 1;

        let _sCnt = bulletin.s;

        let _i = claim.i;
        let _j = claim.j;
        let _tk = claim.k.split(" ");
        let _tx = _tk[1];
        let _ty = _tk[2];
        let _mkl_path = claim.m.map(function (i) {
            return '0x' + i;
        })

        await zkPoDEX.claimComplaint(alice, _sessionId, _i, _j, _tx, _ty, _mkl_path, _sCnt, { from: bob })
    });

    it("should claimComplaint fail (for not evil)", async () => {
        let path = testdataPath + "/complaint/evil";
        let bulletin = JSON.parse(fs.readFileSync(testdataPath + "/bulletin.plain.json"));
        let claim = JSON.parse(fs.readFileSync(path + "/claim"));
        let _sessionId = 0;

        let _sCnt = bulletin.s;

        let _i = claim.i;
        let _j = claim.j;
        let _tk = claim.k.split(" ");
        let _tx = _tk[1];
        let _ty = _tk[2];
        let _mkl_path = claim.m.map(function (i) {
            return '0x' + i;
        })

        await truffleAssert.reverts(
            zkPoDEX.claimComplaint(alice, _sessionId, _i, _j, _tx, _ty, _mkl_path, _sCnt, { from: bob }),
            "invalid mkl proof"
        )
    });

    it("should submitProofAtomicSwap correctly (not evil)", async () => {
        let path = testdataPath + "/atomic-swap/not_evil";
        let receipt = JSON.parse(fs.readFileSync(path + "/receipt"));
        let secret = JSON.parse(fs.readFileSync(path + "/secret"));
        let bulletin = JSON.parse(fs.readFileSync(testdataPath + "/bulletin.plain.json"));
        let _sCnt = bulletin.s;
        let _sessionId = 2;

        let _seed0 = "0x" + secret.s;
        let _from = bob;
        let _seed2 = "0x" + receipt.s;
        let _sigma_vw = receipt.vw;
        let _count = receipt.c;
        let _price = web3.utils.toWei('0.5', 'ether');
        let _expireAt = Math.floor(Date.now() / 1000) + 3600;

        await zkPoDEX.bobDeposit(alice, {from: bob, value: _price});

        let hash = web3.utils.soliditySha3({ t: 'uint256', v: _sessionId }, { t: 'address', v: _from }, { t: 'bytes32', v: _seed2 }, { t: 'uint256', v: _sigma_vw }, { t: 'uint64', v: _count }, { t: 'uint256', v: _price }, { t: 'uint256', v: _expireAt });
        let signature = await web3.eth.sign(hash, bob);

        await zkPoDEX.submitProofAtomicSwap(_seed0, _sCnt, _sessionId, _from, _seed2, _sigma_vw, _count, _price, _expireAt, signature, { from: alice })

        await truffleAssert.reverts(zkPoDEX.submitProofAtomicSwap(_seed0, _sCnt, _sessionId, _from, _seed2, _sigma_vw, _count, _price, _expireAt, signature, { from: alice }), "not new")
    });

    it("should submitProofAtomicSwap fail (evil)", async () => {
        let path = testdataPath + "/atomic-swap/evil";
        let receipt = JSON.parse(fs.readFileSync(path + "/receipt"));
        let secret = JSON.parse(fs.readFileSync(path + "/secret"));
        let bulletin = JSON.parse(fs.readFileSync(testdataPath + "/bulletin.plain.json"));
        let _sCnt = bulletin.s;
        let _sessionId = 3;

        let _seed0 = "0x" + secret.s;
        let _from = bob;
        let _seed2 = "0x" + receipt.s;
        let _sigma_vw = receipt.vw;
        let _count = receipt.c;
        let _price = web3.utils.toWei('0.5', 'ether');
        let _expireAt = Math.floor(Date.now() / 1000) + 3600;

        await zkPoDEX.bobDeposit(alice, {from: bob, value: _price});

        let hash = web3.utils.soliditySha3({ t: 'uint256', v: _sessionId }, { t: 'address', v: _from }, { t: 'bytes32', v: _seed2 }, { t: 'uint256', v: _sigma_vw }, { t: 'uint64', v: _count }, { t: 'uint256', v: _price }, { t: 'uint256', v: _expireAt });
        let signature = await web3.eth.sign(hash, bob);

        await truffleAssert.reverts(zkPoDEX.submitProofAtomicSwap(_seed0, _sCnt, _sessionId, _from, _seed2, _sigma_vw, _count, _price, _expireAt, signature, { from: alice }), "invalid proof")
    });

    it("should submitProofAtomicSwapVC correctly (not evil)", async () => {
        let path = testdataPath + "/atomic-swap-vc/not_evil";
        let receipt = JSON.parse(fs.readFileSync(path + "/receipt"));
        let secret = JSON.parse(fs.readFileSync(path + "/secret"));
        let bulletin = JSON.parse(fs.readFileSync(testdataPath + "/atomic-swap-vc/bulletin.plain.json"));
        let _sessionId = 4;

        let _seed0 = secret.s;
        let _seed0_rand = secret.r;
        let _from = bob;
        let _seed0_mimc3_digest = receipt.d;
        let _price = web3.utils.toWei('0.5', 'ether');
        let _expireAt = Math.floor(Date.now() / 1000) + 3600;

        await zkPoDEX.bobDeposit(alice, {from: bob, value: _price});

        let hash = web3.utils.soliditySha3({ t: 'uint256', v: _sessionId }, { t: 'address', v: _from }, { t: 'uint256', v: _seed0_mimc3_digest }, { t: 'uint256', v: _price }, { t: 'uint256', v: _expireAt });
        let signature = await web3.eth.sign(hash, bob);

        await zkPoDEX.submitProofAtomicSwapVC(_seed0, _seed0_rand, _sessionId, _from, _seed0_mimc3_digest, _price, _expireAt, signature, { from: alice })

        await truffleAssert.reverts(zkPoDEX.submitProofAtomicSwapVC(_seed0, _seed0_rand, _sessionId, _from, _seed0_mimc3_digest, _price, _expireAt, signature, { from: alice }), "not new")
    });

    it("should submitProofAtomicSwapVC fail (evil)", async () => {
        let path = testdataPath + "/atomic-swap-vc/evil";
        let receipt = JSON.parse(fs.readFileSync(path + "/receipt"));
        let secret = JSON.parse(fs.readFileSync(path + "/secret"));
        let bulletin = JSON.parse(fs.readFileSync(testdataPath + "/atomic-swap-vc/bulletin.plain.json"));
        let _sessionId = 5;

        let _seed0 = secret.s;
        let _seed0_rand = secret.r;
        let _from = bob;
        let _seed0_mimc3_digest = receipt.d;
        let _price = web3.utils.toWei('0.5', 'ether');
        let _expireAt = Math.floor(Date.now() / 1000) + 3600;

        await zkPoDEX.bobDeposit(alice, {from: bob, value: _price});

        let hash = web3.utils.soliditySha3({ t: 'uint256', v: _sessionId }, { t: 'address', v: _from }, { t: 'uint256', v: _seed0_mimc3_digest }, { t: 'uint256', v: _price }, { t: 'uint256', v: _expireAt });
        let signature = await web3.eth.sign(hash, bob);

        await truffleAssert.reverts(zkPoDEX.submitProofAtomicSwapVC(_seed0, _seed0_rand, _sessionId, _from, _seed0_mimc3_digest, _price, _expireAt, signature, { from: alice }), "invalid proof")
    });

    // it("", async () => {

    // });

});