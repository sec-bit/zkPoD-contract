pragma solidity ^0.5.0;
pragma experimental ABIEncoderV2;

import "./interface/PublicVarInterface.sol";
import "./lib/ECDSA.sol";

contract PODEX {

    struct Bulletin {
        address owner;
        uint64 size;
        uint64 s;
        uint64 n;
        uint256 sigma_mkl_root;
        uint256 vrf_meta_digest;
        uint256 pledge_value;
        uint256 unDepositAt;
        BltType blt_type;
        DepositStat stat;
    }

    struct Deposit {
        uint256 value;
        uint256 unDepositAt;
        uint256 pendingCnt;
        DepositStat stat;
    }

    /* Receipt Struct */

    struct Batch1Receipt {
        uint256 sid;
        address from;
        bytes32 seed2;
        bytes32 k_mkl_root;
        uint64 count;
        uint256 price;
        uint256 expireAt;
    }

    struct Batch2Receipt {
        uint256 sid;
        address from;
        bytes32 seed2;
        uint256 sigma_vw;
        uint64 count;
        uint256 price;
        uint256 expireAt;
    }

    struct Batch3Receipt {
        uint256 sid;
        address from;
        G1Point u0_x0_lgs;
        G1Point u0d;
        uint256 price;
        uint256 expireAt;
    }

    struct VRFReceipt {
        uint256 sid;
        address from;
        G1Point g_exp_r;
        uint256 price;
        uint256 expireAt;
    }

    /* Session Record Struct */
    struct SessionRecord {
        uint256 submitAt;
        TradeMode mode;
        TradeStat stat;
    }

    struct Batch1Record {
        bytes32 seed0;
        Batch1Receipt receipt;
    }

    struct Batch2Record {
        bytes32 seed0;
    }

    struct Batch3Record {
        uint256 d;
        uint256 x0_lgs;
    }

    struct VRFRecord {
        uint256 r;
    }

    /* Others */

    struct G1Point {
        uint256 X;
        uint256 Y;
    }

    /* Enum */

    enum BltType {
        PLAIN,
        TABLE
    }

    enum DepositStat {
        OK,
        CANCELING,
        CANCELED
    }

    enum TradeStat {
        DEAL,
        WAIT,
        CLAIMED
    }

    enum TradeMode {
        BATCH1,
        BATCH2,
        BATCH3,
        VRF
    }

    /* Mappings */

    mapping (bytes32 => Bulletin) public bulletins_;
    // B => A => Deposit
    mapping (address => mapping(address => Deposit)) public buyerDeposits_;
    // A => B => sid => Record
    mapping (address => mapping(address => mapping(uint256 => SessionRecord))) internal sessionRecords_;
    mapping (address => mapping(address => mapping(uint256 => Batch1Record))) internal batch1Records_;
    mapping (address => mapping(address => mapping(uint256 => Batch2Record))) internal batch2Records_;
    mapping (address => mapping(address => mapping(uint256 => Batch3Record))) internal batch3Records_;
    mapping (address => mapping(address => mapping(uint256 => VRFRecord))) internal vrfRecords_;

    /* Variables */
    uint256 public t1 = 8 hours; // B must wait for t1 to withdraw after unDeposit
    // uint256 public t2 = 12 hours; // For receipt, not set in contract
    uint256 public t3 = 24 hours; // B must claim in t3 after A submitProofBatch1 or A'll get paid after t3
    uint256 public t4 = 3 days; // A must wait for t4 after unPublish

    uint64 public s_ = 65;
    uint256 internal constant GEN_ORDER = 0x30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f0000001;

    PublicVarInterface public publicVar_;

    /* Events */
    event OnDeal(address indexed _a, address indexed _b, uint256 indexed _sid, TradeMode _mode, uint256 _price);
    event OnBatch1Key(address indexed _a, address indexed _b, uint256 indexed _sid, bytes32 _seed0);
    event OnBatch1Claim(address indexed _a, address indexed _b, uint256 indexed _sid);
    event OnBatch1Deal(address indexed _a, address indexed _b, uint256 indexed _sid, uint256 _price);
    event OnBatch2Deal(address indexed _a, address indexed _b, uint256 indexed _sid, bytes32 _seed0);
    event OnBatch3Deal(address indexed _a, address indexed _b, uint256 indexed _sid, uint256 _d, uint256 _x0_lgs);
    event OnVRFDeal(address indexed _a, address indexed _b, uint256 indexed _sid, uint256 _r);

    /* Debug Events */
    // event LogG1Point(uint256 _x, uint256 _y);

    constructor(address _publicVar)
        public
    {
        publicVar_ = PublicVarInterface(_publicVar);
    }

    /* Core Functions - Fund Related */

    // TODO: minimal pledge_value
    function publish(uint64 _size, uint64 _s, uint64 _n, uint256 _sigma_mkl_root, uint256 _vrf_meta_digest, uint256 _blt_type)
        public
        payable
    {
        bytes32 _bltKey;
        BltType _bltType;
        if (_blt_type == 0) {
            _bltType = BltType.PLAIN;
            _bltKey = keccak256(abi.encodePacked(_size, _s, _n, _sigma_mkl_root));
        } else if (_blt_type == 1) {
            _bltType = BltType.TABLE;
            _bltKey = keccak256(abi.encodePacked(_s, _n, _sigma_mkl_root, _vrf_meta_digest));
        } else {
            revert("wrong type");
        }
        require(bulletins_[_bltKey].owner == address(0) || bulletins_[_bltKey].stat == DepositStat.CANCELED, "blt occupied");
        Bulletin memory file = Bulletin({
             owner: msg.sender,
             size: _size,
             s: _s,
             n: _n,
             sigma_mkl_root: _sigma_mkl_root,
             vrf_meta_digest: _vrf_meta_digest,
             pledge_value: msg.value,
             unDepositAt: 0,
             blt_type: _bltType,
             stat: DepositStat.OK
        });
        bulletins_[_bltKey] = file;
    }

    function unPublish(bytes32 _bltKey)
        public
    {
        require(bulletins_[_bltKey].stat == DepositStat.OK, "wrong stat");
        bulletins_[_bltKey].stat = DepositStat.CANCELING;
        bulletins_[_bltKey].unDepositAt = now;
    }

    function buyerDeposit(address _to)
        public
        payable
    {
        buyerDeposits_[msg.sender][_to].value = buyerDeposits_[msg.sender][_to].value + msg.value;
        if (buyerDeposits_[msg.sender][_to].stat != DepositStat.OK) {
            buyerDeposits_[msg.sender][_to].stat = DepositStat.OK;
        }
    }

    function buyerUnDeposit(address _to)
        public
    {
        require(buyerDeposits_[msg.sender][_to].pendingCnt == 0, "pending not allowed");
        require(buyerDeposits_[msg.sender][_to].stat == DepositStat.OK, "wrong stat");
        buyerDeposits_[msg.sender][_to].stat = DepositStat.CANCELING;
        buyerDeposits_[msg.sender][_to].unDepositAt = now;
    }

    function withdrawA(bytes32 _bltKey)
        public
    {
        require(bulletins_[_bltKey].owner == msg.sender, "not owner");
        require(bulletins_[_bltKey].stat == DepositStat.CANCELING, "wrong stat");
        require(now - bulletins_[_bltKey].unDepositAt > t4, "be patient");
        // transfer
        uint256 _value = bulletins_[_bltKey].pledge_value;
        bulletins_[_bltKey].pledge_value = 0;
        bulletins_[_bltKey].stat == DepositStat.CANCELED;
        msg.sender.transfer(_value);
    }

    function withdrawB(address _to)
        public
    {
        require(buyerDeposits_[msg.sender][_to].stat == DepositStat.CANCELING, "wrong stat");
        require(now - buyerDeposits_[msg.sender][_to].unDepositAt > t1, "be patient");
        // transfer
        uint256 _value = buyerDeposits_[msg.sender][_to].value;
        buyerDeposits_[msg.sender][_to].value = 0;
        buyerDeposits_[msg.sender][_to].stat == DepositStat.CANCELED;
        msg.sender.transfer(_value);
    }

    /* Core Functions - Submit Proof */

    // mode: plain (range_pod/ot_range_pod), table (ot_batch_pod/batch_pod)
    function submitProofBatch1
    (
        bytes32 _seed0,
        // receipt
        uint256 _sid,
        address _b,
        bytes32 _seed2,
        bytes32 _k_mkl_root,
        uint64 _count,
        uint256 _price,
        uint256 _expireAt,
        // sig
        bytes memory _sig
    )
        public
    {
        vrfyCommon(_b, _sid, _expireAt);

        Batch1Receipt memory _receipt = Batch1Receipt({
            sid: _sid,
            from: _b,
            seed2: _seed2,
            k_mkl_root: _k_mkl_root,
            count: _count,
            price: _price,
            expireAt: _expireAt
        });

        require(checkSigBatch1(_b, _receipt, _sig), "wrong sig");

        setBatch1Key(msg.sender, _b, _sid, _receipt, _seed0);
    }

    // mode: plain (range_pod/ot_range_pod), table (ot_batch_pod/batch_pod)
    function claimBatch1(address _a, uint256 _sid, uint64 _i, uint64 _j, uint256 _tx, uint256 _ty, bytes32[] memory _mkl_path, uint64 _sCnt)
        public
    {
        require(sessionRecords_[_a][msg.sender][_sid].stat == TradeStat.WAIT, "wrong stat");
        require(now - sessionRecords_[_a][msg.sender][_sid].submitAt < t3, "timeout");
        // loadReceipt
        Batch1Record memory _record = batch1Records_[_a][msg.sender][_sid];
        require(vrfyProofBatch1(_i, _j, _sCnt, _tx, _ty, _record, _mkl_path), "invalid proof");

        // return back money
        sessionRecords_[_a][msg.sender][_sid].stat = TradeStat.CLAIMED;
        uint256 _value = buyerDeposits_[msg.sender][_a].value;
        buyerDeposits_[msg.sender][_a].value = 0;
        msg.sender.transfer(_value);

        // TODO: panish evil A
        buyerDeposits_[msg.sender][_a].pendingCnt -= 1;

        emit OnBatch1Claim(_a, msg.sender, _sid);
    }

    // mode: plain (range_pod/ot_range_pod), table (ot_batch_pod/batch_pod)
    // TODO: discount if b confirm order before timeout
    function settleBatch1Deal(address payable _a, address _b, uint256 _sid)
        public
    {
        require(sessionRecords_[_a][_b][_sid].stat == TradeStat.WAIT, "wrong stat");
        require(now - sessionRecords_[_a][_b][_sid].submitAt >= t3, "timeout");

        uint256 _value = batch1Records_[_a][_b][_sid].receipt.price;

        // transfer
        settleBalance(_b, _a, _value);

        sessionRecords_[_a][_b][_sid].stat == TradeStat.DEAL;
        buyerDeposits_[_b][_a].pendingCnt -= 1;

        emit OnBatch1Deal(_a, _b, _sid, _value);
        emit OnDeal(_a, _b, _sid, TradeMode.BATCH1, _value);
    }

    // mode: plain (TODO), table (batch2_pod)
    function submitProofBatch2
    (
        bytes32 _seed0,
        uint64 _sCnt,
        // receipt
        uint256 _sid,
        address _b,
        bytes32 _seed2,
        uint256 _sigma_vw,
        uint64 _count,
        uint256 _price,
        uint256 _expireAt,
        // sig
        bytes memory _sig
    )
        public
    {
        vrfyCommon(_b, _sid, _expireAt);

        Batch2Receipt memory _receipt = Batch2Receipt({
            sid: _sid,
            from: _b,
            seed2: _seed2,
            sigma_vw: _sigma_vw,
            count: _count,
            price: _price,
            expireAt: _expireAt
        });

        require(checkSigBatch2(_b, _receipt, _sig), "wrong sig");
        require(vrfyProofBatch2(_count, _sCnt, _seed0, _seed2, _sigma_vw), "invalid proof");

        setBatch2Deal(msg.sender, _b, _sid, _price, _seed0);
    }

    // mode: plain (batch3_pod), table (batch3_pod)
    function submitProofBatch3
    (
        uint256 _s_d, uint256 _s_x0_lgs,
        uint256 _sid,
        address _b,
        uint256[2] memory _r_u0_x0_lgs,
        uint256[2] memory _r_u0d,
        uint256 _price,
        uint256 _expireAt,
        bytes memory _sig
    )
        public
    {
        vrfyCommon(_b, _sid, _expireAt);

        Batch3Receipt memory _receipt = Batch3Receipt({
            sid: _sid,
            from: _b,
            u0_x0_lgs: G1Point(_r_u0_x0_lgs[0], _r_u0_x0_lgs[1]),
            u0d: G1Point(_r_u0d[0], _r_u0d[1]),
            price: _price,
            expireAt: _expireAt
        });

        require(checkSigBatch3(_b, _receipt, _sig), "wrong sig");
        require(vrfyProofBatch3(_r_u0d, _r_u0_x0_lgs, _s_d, _s_x0_lgs), "invalid proof");

        setBatch3Deal(msg.sender, _b, _sid, _price, _s_d, _s_x0_lgs);
    }

    // mode: table (vrf_query/ot_vrf_query)
    function submitProofVRF
    (
        uint256 _s_r, // secret r
        // receipt
        uint256 _sid,
        address _b,
        uint256[2] memory _g_exp_r,
        uint256 _price,
        uint256 _expireAt,
        // sig
        bytes memory _sig
    )
        public
    {
        vrfyCommon(_b, _sid, _expireAt);

        VRFReceipt memory _receipt = VRFReceipt({
            sid: _sid,
            from: _b,
            g_exp_r: G1Point(_g_exp_r[0], _g_exp_r[1]),
            price: _price,
            expireAt: _expireAt
        });

        require(checkSigVRF(_b, _receipt, _sig), "wrong sig");
        require(vrfyProofVRF(_g_exp_r, _s_r), "invalid proof");

        setVRFDeal(msg.sender, _b, _sid, _price, _s_r);
    }

    /* Helper Functions - Settle */

    function settleBalance(address _b, address payable _a, uint256 _value)
        internal
    {
        uint256 _balance = buyerDeposits_[_b][_a].value;
        require(_balance >= _value, "short balance");
        buyerDeposits_[_b][_a].value = _balance - _value;
        _a.transfer(_value);
    }

    function setBatch1Key(address _a, address _b, uint256 _sid, Batch1Receipt memory _receipt, bytes32 _seed0)
        internal
    {
        batch1Records_[_a][_b][_sid] = Batch1Record({
            seed0: _seed0,
            receipt: _receipt
        });
        sessionRecords_[_a][_b][_sid] = SessionRecord({
            submitAt: now,
            mode: TradeMode.BATCH1,
            stat: TradeStat.WAIT
        });
        buyerDeposits_[_b][_a].pendingCnt += 1;

        emit OnBatch1Key(_a, _b, _sid, _seed0);
    }

    function setBatch2Deal(address payable _a, address _b, uint256 _sid, uint256 _price, bytes32 _seed0)
        internal
    {
        batch2Records_[_a][_b][_sid] = Batch2Record({
            seed0: _seed0
        });
        sessionRecords_[_a][_b][_sid] = SessionRecord({
            submitAt: now,
            mode: TradeMode.BATCH2,
            stat: TradeStat.DEAL
        });

        // transfer
        settleBalance(_b, _a, _price);

        emit OnBatch2Deal(_a, _b, _sid, _seed0);
        emit OnDeal(_a, _b, _sid, TradeMode.BATCH2, _price);
    }

    function setBatch3Deal(address payable _a, address _b, uint256 _sid, uint256 _price, uint256 _s_d, uint256 _s_x0_lgs)
        internal
    {
        batch3Records_[_a][_b][_sid] = Batch3Record({
            d: _s_d,
            x0_lgs: _s_x0_lgs
        });
        sessionRecords_[_a][_b][_sid] = SessionRecord({
            submitAt: now,
            mode: TradeMode.BATCH3,
            stat: TradeStat.DEAL
        });

        // transfer
        settleBalance(_b, _a, _price);

        emit OnBatch3Deal(_a, _b, _sid, _s_d, _s_x0_lgs);
        emit OnDeal(_a, _b, _sid, TradeMode.BATCH3, _price);
    }

    function setVRFDeal(address payable _a, address _b, uint256 _sid, uint256 _price, uint256 _s_r)
        internal
    {
        vrfRecords_[_a][_b][_sid] = VRFRecord({
            r: _s_r
        });
        sessionRecords_[_a][_b][_sid] = SessionRecord({
            submitAt: now,
            mode: TradeMode.VRF,
            stat: TradeStat.DEAL
        });

        // transfer
        settleBalance(_b, _a, _price);

        emit OnVRFDeal(_a, _b, _sid, _s_r);
        emit OnDeal(_a, _b, _sid, TradeMode.VRF, _price);
    }

    /* Helper Functions - Verify Proof */

    function vrfyCommon(address _b, uint256 _sid, uint256 _expireAt)
        internal
    {
        require(sessionRecords_[msg.sender][_b][_sid].submitAt == 0, "not new");
        // require(now < _expireAt, "expired");
        DepositStat _stat = buyerDeposits_[_b][msg.sender].stat;
        require(_stat != DepositStat.CANCELED, "deposit canceled");
        if (_stat == DepositStat.CANCELING) {
            buyerDeposits_[_b][msg.sender].stat = DepositStat.OK;
        }
    }

    function vrfyProofBatch1(uint64 _i, uint64 _j, uint64 _sCnt, uint256 _tx, uint256 _ty, Batch1Record memory _record, bytes32[] memory _mkl_path)
        internal
        view
        returns (bool)
    {
        // vrfy mkl path
        // TODO: think about overflow here
        uint64 _index = _i*_sCnt + _j;
        bytes32 _x = convertToBE(bytes32(_tx));
        bytes32 _y = convertToBE(bytes32(_ty));
        require(vrfyPath(_x, _y, _index, _record.receipt.count*_sCnt, _record.receipt.k_mkl_root, _mkl_path), "invalid mkl proof");
        // derive k
        uint256 _v = chain(_record.seed0, _index);
        // calc u^v
        (uint256 _u1_x, uint256 _u1_y) = publicVar_.getEccPubU1(_j);
        G1Point memory _check = scalarMul(_u1_x, _u1_y, _v);
        // emit LogG1Point(_check.X, _check.Y);
        require(_check.X != _tx || _check.Y != _ty, "invalid claim");
        return true;
    }

    function vrfyProofBatch2(
        uint64 count, uint64 s,
        bytes32 seed0, bytes32 seed2,
        uint sigma_vw
    )
        public
        pure
        returns(bool)
    {
        uint256 v;
        uint256 w;
        uint256 check_sigma_vw;
        uint256 sigma_v;
        uint64 offset = count * s;

        for (uint64 j = 0; j < s; j++) {
            v = chain(seed0, offset + j);
            check_sigma_vw = addmod(check_sigma_vw, v, GEN_ORDER);
        }

        for(uint64 i = 0; i < count; i++){
            sigma_v = 0;
            w = chain(seed2, i);
            for(uint64 j = 0; j < s; j++){
                v = chain(seed0, i*s+j);
                sigma_v = addmod(sigma_v, v, GEN_ORDER);
            }
            check_sigma_vw = addmod(check_sigma_vw, mulmod(sigma_v, w, GEN_ORDER), GEN_ORDER);
        }

        return check_sigma_vw == sigma_vw;
    }

    function vrfyProofBatch3(
        uint256[2] memory _r_u0d,
        uint256[2] memory  _r_u0_x0_lgs,
        uint256 _s_d, uint256 _s_x0_lgs
    )
        public
        view
        returns (bool)
    {
        (uint256 _u1_x, uint256 _u1_y) = publicVar_.getEccPubU1(0);
        G1Point memory _check = scalarMul(_u1_x, _u1_y, _s_d);
        require(_r_u0d[0] == _check.X && _r_u0d[1] == _check.Y, "wrong d");
        _check = scalarMul(_u1_x, _u1_y, _s_x0_lgs);
        require(_r_u0_x0_lgs[0] == _check.X && _r_u0_x0_lgs[1] == _check.Y, "wrong x0");
        return true;
    }

    function vrfyProofVRF(
        uint256[2] memory _g_exp_r,
        uint256 _s_r
    )
        public
        view
        returns (bool)
    {
        // g1: 1 1 2
        G1Point memory _check = scalarMul(1, 2, _s_r);
        require(_g_exp_r[0] == _check.X && _g_exp_r[1] == _check.Y, "wrong r");
        return true;
    }

    function vrfyPath(bytes32 _x, bytes32 _y, uint64 _ij, uint64 _ns, bytes32 _root, bytes32[] memory _mkl_path)
        public
        pure
        returns (bool)
    {
        bytes32 _value = hashInSha3(_x, _y);
        uint256 _depth = log2ub(_ns);
        if (_mkl_path.length != _depth) {
            return false;
        }

        uint64 _pos = _ij;
        for (uint256 _i = 0; _i < _depth; _i++) {
            if (_pos % 2 != 0) {
                _value = hashInSha3(_mkl_path[_i], _value);
            } else {
                _value = hashInSha3(_value, _mkl_path[_i]);
            }
            _pos /= 2;
        }
        return (_value == _root);
    }

    /* Helper Functions - Calculation */

    function hashInSha3(bytes32 _x, bytes32 _y)
        public
        pure
        returns (bytes32)
    {
        return keccak256(abi.encodePacked(_x, _y));
    }

    function hashInSha3(bytes32 _x, uint64 _y)
        public
        pure
        returns (bytes32)
    {
        return keccak256(abi.encodePacked(_x, _y));
    }

    function chain(bytes32 seed, uint64 index)
        public
        pure
        returns (uint256)
    {
        bytes32 _ret = hashInSha3(seed, index);
        return uint256(_ret) % GEN_ORDER;
    }

    function scalarMul(uint256 _x, uint256 _y, uint256 _s)
        internal
        view
        returns (G1Point memory r)
    {
        uint[3] memory input;
        input[0] = _x;
        input[1] = _y;
        input[2] = _s;
        bool success;
        assembly {
            success := staticcall(sub(gas, 2000), 7, input, 0x80, r, 0x60)
            // Use "invalid" to make gas estimation work
            switch success case 0 { invalid() }
        }
        require (success, "failed");
    }

    function log2ub(uint256 _n)
        public
        pure
        returns (uint256)
    {
        if (_n == 1) return 0;
        if (_n % 2 != 0) _n++;
        return 1 + log2ub(_n / 2);
    }

    function convertToBE(bytes32 _in)
        internal
        pure
        returns (bytes32)
    {
        bytes memory _bytes = new bytes(32);
        for (uint256 j = 0; j < 32; j++) {
            _bytes[j] = _in[31-j];
        }
        bytes32 _out;
        assembly {
            _out := mload(add(_bytes, 32))
        }
        return _out;
    }

    /* Helper Functions - Receipt Signature */

    function checkSigBatch1(address addr, Batch1Receipt memory r1, bytes memory sig)
        internal
        pure
        returns (bool)
    {
        bytes32 hash = keccak256(abi.encodePacked(r1.sid, r1.from, r1.seed2, r1.k_mkl_root, r1.count, r1.price, r1.expireAt));
        return checkSig(addr, hash, sig);
    }

    function checkSigBatch2(address addr, Batch2Receipt memory r1, bytes memory sig)
        internal
        pure
        returns (bool)
    {
        bytes32 hash = keccak256(abi.encodePacked(r1.sid, r1.from, r1.seed2, r1.sigma_vw, r1.count, r1.price, r1.expireAt));
        return checkSig(addr, hash, sig);
    }

    function checkSigBatch3(address addr, Batch3Receipt memory r1, bytes memory sig)
        internal
        pure
        returns (bool)
    {
        bytes32 hash = keccak256(abi.encodePacked(r1.sid, r1.from, r1.u0_x0_lgs.X, r1.u0_x0_lgs.Y, r1.u0d.X, r1.u0d.Y, r1.price, r1.expireAt));
        return checkSig(addr, hash, sig);
    }

    function checkSigVRF(address addr, VRFReceipt memory r1, bytes memory sig)
        internal
        pure
        returns (bool)
    {
        bytes32 hash = keccak256(abi.encodePacked(r1.sid, r1.from, r1.g_exp_r.X, r1.g_exp_r.Y, r1.price, r1.expireAt));
        return checkSig(addr, hash, sig);
    }

    function checkSig(address addr, bytes32 hash, bytes memory sig)
        internal
        pure
        returns (bool)
    {
        return addr == ECDSA.recover(ECDSA.toEthSignedMessageHash(hash), sig);
    }

    /* Public Getters */

    function getSessionRecord(address _a, address _b, uint256 _sid)
        public
        view
        returns (uint256 submitAt, TradeMode mode, TradeStat stat)
    {
        SessionRecord memory _sRecord = sessionRecords_[_a][_b][_sid];
        submitAt = _sRecord.submitAt;
        mode = _sRecord.mode;
        stat = _sRecord.stat;
    }

    function getRecordBatch1(address _a, address _b, uint256 _sid)
        public
        view
        returns (
            bytes32 seed0,
            bytes32 seed2,
            bytes32 k_mkl_root,
            uint64 count,
            uint256 price,
            uint256 expireAt,
            uint256 submitAt
        )
    {
        Batch1Record memory _record = batch1Records_[_a][_b][_sid];
        seed0 = _record.seed0;
        seed2 = _record.receipt.seed2;
        k_mkl_root = _record.receipt.k_mkl_root;
        count = _record.receipt.count;
        price = _record.receipt.price;
        expireAt = _record.receipt.expireAt;
        submitAt = sessionRecords_[_a][_b][_sid].submitAt;
    }

    function getRecordBatch2(address _a, address _b, uint256 _sid)
        public
        view
        returns (
            bytes32 seed0,
            uint256 submitAt
        )
    {
        Batch2Record memory _record = batch2Records_[_a][_b][_sid];
        seed0 = _record.seed0;
        submitAt = sessionRecords_[_a][_b][_sid].submitAt;
    }

    function getRecordBatch3(address _a, address _b, uint256 _sid)
        public
        view
        returns (
            uint256 d,
            uint256 x0_lgs,
            uint256 submitAt
        )
    {
        Batch3Record memory _record = batch3Records_[_a][_b][_sid];
        d = _record.d;
        x0_lgs = _record.x0_lgs;
        submitAt = sessionRecords_[_a][_b][_sid].submitAt;
    }

    function getRecordVRF(address _a, address _b, uint256 _sid)
        public
        view
        returns (
            uint256 r,
            uint256 submitAt
        )
    {
        VRFRecord memory _record = vrfRecords_[_a][_b][_sid];
        r = _record.r;
        submitAt = sessionRecords_[_a][_b][_sid].submitAt;
    }

}
