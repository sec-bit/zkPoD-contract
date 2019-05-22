const PODEX = artifacts.require("PODEX");
const { getWeb3 } = require("./helpers");
const web3 = getWeb3();
const fs = require('fs');
const testdataPath = "test/testdata";

contract("PODEX", async (accounts) => {

    let podex;

    before(async () => {
        podex = await PODEX.deployed();
    });

    it("should keep basic settings", async () => {
        let s = await podex.s_();
        let t1 = await podex.t1();
        let t2 = await podex.t2();
        let t3 = await podex.t3();
        assert.equal(s, 65, "max s count changed");
        assert.equal(t1, 8 * 60 * 60, "t1 changed");
        assert.equal(t2, 12 * 60 * 60, "t2 changed");
        assert.equal(t3, 24 * 60 * 60, "t3 changed");
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

        assert.equal(bulletin.size, blt.size, "wrong size");
        assert.equal(bulletin.s, blt.s, "wrong s");
        assert.equal(bulletin.n, blt.n, "wrong n");
        assert.equal(bulletin.sigma_mkl_root, web3.utils.toBN(_sigma_mkl_root).toString(), "wrong sigma_mkl_root");
        assert.equal(bulletin.status, 0, "wrong status");
        assert.equal(bulletin.blt_type, _blt_type, "wrong blt_type");

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

        assert.equal(bulletin.s, blt.s, "wrong s");
        assert.equal(bulletin.n, blt.n, "wrong n");
        assert.equal(bulletin.sigma_mkl_root, web3.utils.toBN(_sigma_mkl_root).toString(), "wrong sigma_mkl_root");
        assert.equal(bulletin.vrf_meta_digest, web3.utils.toBN(_vrf_meta_digest).toString(), "wrong vrf_meta_digest");
        assert.equal(bulletin.status, 0, "wrong status");
        assert.equal(bulletin.blt_type, _blt_type, "wrong blt_type");

    });

    // it("", async () => {

    // });

});