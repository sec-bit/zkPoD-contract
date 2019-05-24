const PODEX = artifacts.require("PODEX");
const { getWeb3 } = require("./helpers");
const web3 = getWeb3();
const fs = require('fs');
const testdataPath = "test/testdata";
const truffleAssert = require('truffle-assertions');

contract("PODEX", async (accounts) => {

    let podex;

    let seller = accounts[0];
    let buyer = accounts[1];

    before(async () => {
        podex = await PODEX.deployed();
    });

    it("should keep basic settings", async () => {
        let s = await podex.s_();
        let t1 = await podex.t1();
        // let t2 = await podex.t2();
        let t3 = await podex.t3();
        let t4 = await podex.t4()
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

        await podex.publish(_size, _s, _n, _sigma_mkl_root, 0, _blt_type);

        let _bltKey = web3.utils.soliditySha3({ t: 'uint64', v: _size }, { t: 'uint64', v: _s }, { t: 'uint64', v: _n }, { t: 'uint256', v: _sigma_mkl_root });
        let bulletin = await podex.bulletins_(_bltKey);

        assert.equal(bulletin.owner, seller, "wrong owner");
        assert.equal(bulletin.size, blt.size, "wrong size");
        assert.equal(bulletin.s, blt.s, "wrong s");
        assert.equal(bulletin.n, blt.n, "wrong n");
        assert.equal(bulletin.sigma_mkl_root, web3.utils.toBN(_sigma_mkl_root).toString(), "wrong sigma_mkl_root");
        assert.equal(bulletin.stat, 0, "wrong status");
        assert.equal(bulletin.blt_type, _blt_type, "wrong blt_type");

        await truffleAssert.reverts(
            podex.publish(_size, _s, _n, _sigma_mkl_root, 0, _blt_type),
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

        await podex.publish(_size, _s, _n, _sigma_mkl_root, _vrf_meta_digest, _blt_type);

        let _bltKey = web3.utils.soliditySha3({ t: 'uint64', v: _s }, { t: 'uint64', v: _n }, { t: 'uint256', v: _sigma_mkl_root }, { t: 'uint256', v: _vrf_meta_digest });
        let bulletin = await podex.bulletins_(_bltKey);

        assert.equal(bulletin.owner, seller, "wrong owner");
        assert.equal(bulletin.s, blt.s, "wrong s");
        assert.equal(bulletin.n, blt.n, "wrong n");
        assert.equal(bulletin.sigma_mkl_root, web3.utils.toBN(_sigma_mkl_root).toString(), "wrong sigma_mkl_root");
        assert.equal(bulletin.vrf_meta_digest, web3.utils.toBN(_vrf_meta_digest).toString(), "wrong vrf_meta_digest");
        assert.equal(bulletin.stat, 0, "wrong status");
        assert.equal(bulletin.blt_type, _blt_type, "wrong blt_type");

        await truffleAssert.reverts(
            podex.publish(_size, _s, _n, _sigma_mkl_root, _vrf_meta_digest, _blt_type),
            "blt occupied"
        );
    });

    it("should submitProofBatch1 correctly (for not evil)", async () => {
        let path = testdataPath + "/batch1/not_evil";
        let receipt = JSON.parse(fs.readFileSync(path + "/receipt"));
        let secret = JSON.parse(fs.readFileSync(path + "/secret"));

        // let _sessionId = Math.random().toString().slice(2, 17);
        let _sessionId = 0;
        let _from = buyer;
        let _seed0 = "0x" + secret.s;
        let _seed2 = "0x" + receipt.s;
        let _k_mkl_root = "0x" + receipt.k;
        let _count = receipt.c;
        let _price = 0;
        let _expireAt = 0;

        let hash = web3.utils.soliditySha3({ t: 'uint256', v: _sessionId }, { t: 'address', v: _from }, { t: 'bytes32', v: _seed2 }, { t: 'bytes32', v: _k_mkl_root }, { t: 'uint64', v: _count }, { t: 'uint256', v: _price }, { t: 'uint256', v: _expireAt });
        let signature = await web3.eth.sign(hash, buyer);
        // console.log("signature:", signature);

        await podex.submitProofBatch1(_seed0, _sessionId, _from, _seed2, _k_mkl_root, _count, _price, _expireAt, signature, {from: seller});

        await truffleAssert.reverts(
            podex.submitProofBatch1(_seed0, _sessionId, buyer, _seed2, _k_mkl_root, _count, _price, _expireAt, signature, {from: seller})
        );
    });

    it("should submitProofBatch1 correctly (for evil)", async () => {
        let path = testdataPath + "/batch1/evil";
        let receipt = JSON.parse(fs.readFileSync(path + "/receipt"));
        let secret = JSON.parse(fs.readFileSync(path + "/secret"));
        let claim = JSON.parse(fs.readFileSync(path + "/claim"));

        // let _sessionId = Math.random().toString().slice(2, 17);
        let _sessionId = 1;
        let _from = buyer;
        let _seed0 = "0x" + secret.s;
        let _seed2 = "0x" + receipt.s;
        let _k_mkl_root = "0x" + receipt.k;
        let _count = receipt.c;
        let _price = 0;
        let _expireAt = 0;

        let hash = web3.utils.soliditySha3({ t: 'uint256', v: _sessionId }, { t: 'address', v: _from }, { t: 'bytes32', v: _seed2 }, { t: 'bytes32', v: _k_mkl_root }, { t: 'uint64', v: _count }, { t: 'uint256', v: _price }, { t: 'uint256', v: _expireAt });
        let signature = await web3.eth.sign(hash, buyer);

        await podex.submitProofBatch1(_seed0, _sessionId, _from, _seed2, _k_mkl_root, _count, _price, _expireAt, signature, { from: seller });
    });

    it("should claimBatch1 correctly (for evil)", async () => {
        let path = testdataPath + "/batch1/evil";
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

        await podex.claimBatch1(seller, _sessionId, _i, _j, _tx, _ty, _mkl_path, _sCnt, { from: buyer })
    });

    it("should claimBatch1 fail (for not evil)", async () => {
        let path = testdataPath + "/batch1/evil";
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
            podex.claimBatch1(seller, _sessionId, _i, _j, _tx, _ty, _mkl_path, _sCnt, { from: buyer }),
            "invalid mkl proof"
        )
    });

    it("should submitProofBatch2 correctly (not evil)", async () => {
        let path = testdataPath + "/batch2/not_evil";
        let receipt = JSON.parse(fs.readFileSync(path + "/receipt"));
        let secret = JSON.parse(fs.readFileSync(path + "/secret"));
        let bulletin = JSON.parse(fs.readFileSync(testdataPath + "/bulletin.plain.json"));
        let _sCnt = bulletin.s;
        let _sessionId = 2;

        let _seed0 = "0x" + secret.s;
        let _from = buyer;
        let _seed2 = "0x" + receipt.s;
        let _sigma_vw = receipt.vw;
        let _count = receipt.c;
        let _price = 0;
        let _expireAt = 0;

        let hash = web3.utils.soliditySha3({ t: 'uint256', v: _sessionId }, { t: 'address', v: _from }, { t: 'bytes32', v: _seed2 }, { t: 'uint256', v: _sigma_vw }, { t: 'uint64', v: _count }, { t: 'uint256', v: _price }, { t: 'uint256', v: _expireAt });
        let signature = await web3.eth.sign(hash, buyer);

        await podex.submitProofBatch2(_seed0, _sCnt, _sessionId, _from, _seed2, _sigma_vw, _count, _price, _expireAt, signature, { from: seller })

        await truffleAssert.reverts(podex.submitProofBatch2(_seed0, _sCnt, _sessionId, _from, _seed2, _sigma_vw, _count, _price, _expireAt, signature, { from: seller }), "not new")
    });

    it("should submitProofBatch2 fail (evil)", async () => {
        let path = testdataPath + "/batch2/evil";
        let receipt = JSON.parse(fs.readFileSync(path + "/receipt"));
        let secret = JSON.parse(fs.readFileSync(path + "/secret"));
        let bulletin = JSON.parse(fs.readFileSync(testdataPath + "/bulletin.plain.json"));
        let _sCnt = bulletin.s;
        let _sessionId = 3;

        let _seed0 = "0x" + secret.s;
        let _from = buyer;
        let _seed2 = "0x" + receipt.s;
        let _sigma_vw = receipt.vw;
        let _count = receipt.c;
        let _price = 0;
        let _expireAt = 0;

        let hash = web3.utils.soliditySha3({ t: 'uint256', v: _sessionId }, { t: 'address', v: _from }, { t: 'bytes32', v: _seed2 }, { t: 'uint256', v: _sigma_vw }, { t: 'uint64', v: _count }, { t: 'uint256', v: _price }, { t: 'uint256', v: _expireAt });
        let signature = await web3.eth.sign(hash, buyer);

        await truffleAssert.reverts(podex.submitProofBatch2(_seed0, _sCnt, _sessionId, _from, _seed2, _sigma_vw, _count, _price, _expireAt, signature, { from: seller }), "invalid proof")
    });

    // it("", async () => {

    // });

});