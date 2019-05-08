pragma experimental ABIEncoderV2;

import "./PublicVar.sol";

contract PODEX is PublicVar {

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
        DepositStatus status;
    }

    struct Deposit {
        uint256 value;
        uint256 unDepositAt;
        DepositStatus status;
    }

    struct Signature {
        uint8 v;
        bytes32 r;
        bytes32 s;
    }

    /* Receipt Struct */

    struct Batch1Receipt {
        uint256 sessionId;
        address from;
        bytes32 seed2;
        bytes32 k_mkl_root;
        uint64 count;
        uint256 price;
        uint256 expireAt;
    }

    struct Batch2Receipt {
        uint256 sessionId;
        address from;
        bytes32 seed2;
        uint256 sigma_vw;
        uint64 count;
        uint256 price;
        uint256 expireAt;
    }

    struct Batch3Receipt {
        uint256 sessionId;
        address from;
        G1Point u0_x0_lgs;
        G1Point u0d;
        uint256 price;
        uint256 expireAt;
    }

    struct VRFReceipt {
        uint256 sessionId;
        address from;
        G1Point g_exp_r;
        uint256 price;
        uint256 expireAt;
    }

    /* Session Record Struct */

    struct Batch1Record {
        bytes32 seed0;
        Batch1Receipt receipt;
        uint256 submitAt;
    }

    struct Batch2Record {
        bytes32 seed0;
        uint256 submitAt;
    }

    struct Batch3Record {
        uint256 d;
        uint256 x0_lgs;
        uint256 submitAt;
    }

    struct VRFRecord {
        uint256 r;
        uint256 submitAt;
    }

    /* Enum */

    enum DepositStatus {
        OK,
        CANCELING,
        CANCELED
    }

    enum BltType {
        PLAIN,
        TABLE
    }

    /* Mappings */

    mapping (bytes32 => Bulletin) public bulletins_;
    // B => A => Deposit
    mapping (address => mapping(address => Deposit)) public buyerDeposits_;
    // A => B => SessionId => Record
    mapping (address => mapping(address => mapping(uint256 => Batch1Record))) internal batch1Records;
    mapping (address => mapping(address => mapping(uint256 => Batch2Record))) internal batch2Records_;
    mapping (address => mapping(address => mapping(uint256 => Batch3Record))) internal batch3Records_;
    mapping (address => mapping(address => mapping(uint256 => VRFRecord))) internal vrfRecords_;

    /* Variables */

    uint64 public s_ = 65;
    uint256 internal constant GEN_ORDER = 0x30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f0000001;

    /* Events */

    /* Debug Events */

    event LogG1Point(uint256 _x, uint256 _y);

    /* Core Functions - Fund Related */

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
        require(bulletins_[_bltKey].owner == address(0), "blt occupied");
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
             status: DepositStatus.OK
        });
        bulletins_[_bltKey] = file;
    }

    function unPublish(bytes32 _bltKey)
        public
    {
        require(bulletins_[_bltKey].status == DepositStatus.OK);
        bulletins_[_bltKey].status = DepositStatus.CANCELING;
        bulletins_[_bltKey].unDepositAt = now;  
    }

    function buyerDeposit(address _to)
        public
        payable
    {
        buyerDeposits_[msg.sender][_to].value = buyerDeposits_[msg.sender][_to].value + msg.value;
    }

    function buyerUnDeposit(address _to)
        public
    {
        require(buyerDeposits_[msg.sender][_to].status == DepositStatus.OK);
        buyerDeposits_[msg.sender][_to].status = DepositStatus.CANCELING;
        buyerDeposits_[msg.sender][_to].unDepositAt = now;
    }

    function withdraw(bytes32 _bltKey)
        public
    {
        require(bulletins_[_bltKey].owner == msg.sender);
        require(bulletins_[_bltKey].status == DepositStatus.CANCELING);
        require(now - bulletins_[_bltKey].unDepositAt > 7 days);
        // transfer
    }

    function withdraw(address _to)
        public
    {
        require(buyerDeposits_[msg.sender][_to].status == DepositStatus.CANCELING);
        require(now - buyerDeposits_[msg.sender][_to].unDepositAt > 7 days);
        // transfer
    }

    /* Core Functions - Submit Proof */

    // mode: plain (range_pod/ot_range_pod), table (ot_batch_pod/batch_pod)
    function submitProofBatch1
    (
        bytes32 _seed0,
        // receipt
        uint256 _sessionId,
        address _b,
        bytes32 _seed2,
        bytes32 _k_mkl_root,
        uint64 _count,
        uint256 _price,
        uint256 _expireAt,
        // sig
        uint8 _v,
        bytes32 _r,
        bytes32 _s
    )
        public
    {
        Batch1Receipt memory _receipt = Batch1Receipt({
            sessionId: _sessionId,
            from: _b,
            seed2: _seed2,
            k_mkl_root: _k_mkl_root,
            count: _count,
            price: _price,
            expireAt: _expireAt
        });

        Signature memory _sig = Signature({
            v: _v,
            r: _r,
            s: _s
        });

        require(checkSigBatch1(_b, _receipt, _sig));
        // require(now < _expireAt);
        require(batch1Records[msg.sender][_b][_sessionId].receipt.from == address(0));
        batch1Records[msg.sender][_b][_sessionId] = Batch1Record({
            seed0: _seed0,
            receipt: _receipt,
            submitAt: now
        });
    }

    // mode: plain (range_pod/ot_range_pod), table (ot_batch_pod/batch_pod)
    function claimBatch1(address _a, uint256 _sessionId, uint64 _i, uint64 _j, uint256 _tx, uint256 _ty, bytes32[] memory _mkl_path, uint64 _sCnt)
        public
    {
        // loadReceipt
        Batch1Record memory _record = batch1Records[_a][msg.sender][_sessionId];
        require(verifyProofBatch1(_i, _j, _sCnt, _tx, _ty, _record, _mkl_path));
        // transfer
    }

    // mode: plain (TODO), table (batch2_pod)
    function submitProofBatch2
    (
        bytes32 _seed0,
        uint64 _sCnt,
        // receipt
        uint256 _sessionId,
        address _b,
        bytes32 _seed2,
        uint256 _sigma_vw,
        uint64 _count,
        uint256 _price,
        uint256 _expireAt,
        // sig
        uint8 _v,
        // bytes32 _r,
        // bytes32 _s
        bytes32[2] memory _rs
    )
        public
    {
        Batch2Receipt memory _receipt = Batch2Receipt({
            sessionId: _sessionId,
            from: _b,
            seed2: _seed2,
            sigma_vw: _sigma_vw,
            count: _count,
            price: _price,
            expireAt: _expireAt
        });

        Signature memory _sig = Signature({
            v: _v,
            r: _rs[0],
            s: _rs[1]
        });

        require(checkSigBatch2(_b, _receipt, _sig));
        // require(now < _expireAt);
        require(batch2Records_[msg.sender][_b][_sessionId].submitAt == 0, "not new");
        require(verifyProofBatch2(_count, _sCnt, _seed0, _seed2, _sigma_vw), "invalid proof");
        // transfer
        batch2Records_[msg.sender][_b][_sessionId] = Batch2Record({
            seed0: _seed0,
            submitAt: now
        });
    }

    // mode: plain (batch3_pod), table (batch3_pod)
    function submitProofBatch3
    (
        uint256 _s_d, uint256 _s_x0_lgs,
        uint256 _sessionId,
        address _b,
        uint256[2] memory _r_u0_x0_lgs,
        uint256[2] memory _r_u0d,
        uint256 _price,
        uint256 _expireAt,
        uint8 _v,
        bytes32 _r,
        bytes32 _s
    )
        public
    {
        Batch3Receipt memory _receipt = Batch3Receipt({
            sessionId: _sessionId,
            from: _b,
            u0_x0_lgs: G1Point(_r_u0_x0_lgs[0], _r_u0_x0_lgs[1]),
            u0d: G1Point(_r_u0d[0], _r_u0d[1]),
            price: _price,
            expireAt: _expireAt
        });

        Signature memory _sig = Signature({
            v: _v,
            r: _r,
            s: _s
        });

        require(checkSigBatch3(_b, _receipt, _sig));
        // require(now < _expireAt);
        require(batch3Records_[msg.sender][_b][_sessionId].submitAt == 0, "not new");
        require(verifyProofBatch3(_r_u0d, _r_u0_x0_lgs, _s_d, _s_x0_lgs), "invalid proof");
        // transfer
        batch3Records_[msg.sender][_b][_sessionId] = Batch3Record({
            d: _s_d,
            x0_lgs: _s_x0_lgs,
            submitAt: now
        });
    }

    // mode: table (vrf_query/ot_vrf_query)
    function submitProofVRF
    (
        uint256 _s_r, // secret r
        // receipt
        uint256 _sessionId,
        address _b,
        uint256[2] memory _g_exp_r,
        uint256 _price,
        uint256 _expireAt,
        // sig
        uint8 _v,
        // bytes32 _r,
        // bytes32 _s
        bytes32[2] memory _rs
    )
        public
    {
        VRFReceipt memory _receipt = VRFReceipt({
            sessionId: _sessionId,
            from: _b,
            g_exp_r: G1Point(_g_exp_r[0], _g_exp_r[1]),
            price: _price,
            expireAt: _expireAt
        });

        Signature memory _sig = Signature({
            v: _v,
            r: _rs[0],
            s: _rs[1]
        });

        require(checkSigVRF(_b, _receipt, _sig));
        // require(now < _expireAt);
        require(vrfRecords_[msg.sender][_b][_sessionId].submitAt == 0, "not new");
        require(verifyProofVRF(_g_exp_r, _s_r), "invalid proof");
        // transfer
        vrfRecords_[msg.sender][_b][_sessionId] = VRFRecord({
            r: _s_r,
            submitAt: now
        });
    }

    /* Helper Functions - Verify Proof */

    function verifyProofBatch1(uint64 _i, uint64 _j, uint64 _sCnt, uint256 _tx, uint256 _ty, Batch1Record memory _record, bytes32[] memory _mkl_path)
        internal
        returns (bool)
    {
        // verify mkl path
        // TODO: think about overflow here
        uint64 _index = _i*_sCnt + _j;
        bytes32 _x = convertToBE(bytes32(_tx));
        bytes32 _y = convertToBE(bytes32(_ty));
        require(verifyPath(_x, _y, _index, _record.receipt.count*_sCnt, _record.receipt.k_mkl_root, _mkl_path), "invalid mkl proof");
        // derive k
        uint256 _v = chain(_record.seed0, _index);
        // calc u^v
        G1Point memory _check = scalarMul(ecc_pub_u1_[_j].X, ecc_pub_u1_[_j].Y, _v);
        emit LogG1Point(_check.X, _check.Y);
        require(_check.X != _tx || _check.Y != _ty, "invalid claim");   
    }

    function verifyProofBatch2(
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

        for(uint64 i=0; i<count; i++){
            sigma_v = 0;
            w = chain(seed2, i);
            for(uint64 j=0; j<s; j++){
                v = chain(seed0, i*s+j);
                sigma_v = addmod(sigma_v, v, GEN_ORDER);
            }
            check_sigma_vw = addmod(check_sigma_vw, mulmod(sigma_v, w, GEN_ORDER), GEN_ORDER);
        }

        return check_sigma_vw == sigma_vw;
    }

    function verifyProofBatch3(
        uint256[2] memory _r_u0d,
        uint256[2] memory  _r_u0_x0_lgs,
        uint256 _s_d, uint256 _s_x0_lgs
    )
        public
        view
        returns (bool)
    {
        uint256 _u1_x = ecc_pub_u1_[0].X;
        uint256 _u1_y = ecc_pub_u1_[0].Y;
        G1Point memory _check = scalarMul(_u1_x, _u1_y, _s_d);
        require(_r_u0d[0] == _check.X && _r_u0d[1] == _check.Y, "wrong d");
        _check = scalarMul(_u1_x, _u1_y, _s_x0_lgs);
        require(_r_u0_x0_lgs[0] == _check.X && _r_u0_x0_lgs[1] == _check.Y, "wrong x0");
        return true;
    }

    function verifyProofVRF(
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

    function verifyPath(bytes32 _x, bytes32 _y, uint64 _ij, uint64 _ns, bytes32 _root, bytes32[] memory _mkl_path)
        public
        pure
        returns (bool)
    {
        bytes32 _value_b = hashInSha3(_x, _y);
        uint256 _depth = log2ub(_ns);
        if (_mkl_path.length != _depth) {
            return false;
        }
        bytes32 _value = _value_b;

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

        bytes memory shasum = new bytes(32);
        uint j;
        for (j=0; j<32; j++) {
            shasum[j] = _ret[j];
        }

        shasum[0] = byte(uint8(shasum[0]) & 63);

        bytes32 revhash;
        assembly {
            revhash := mload(add(shasum, 32))
        }

        bytes32 _retMod = bytes32(uint256(revhash) % GEN_ORDER);
        return uint256(_retMod);
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
        require (success);
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

    function checkSigBatch1(address addr, Batch1Receipt memory r1, Signature memory sig)
        internal
        pure
        returns (bool)
    {
        bytes32 hash = keccak256(abi.encodePacked(r1.sessionId, r1.from, r1.seed2, r1.k_mkl_root, r1.count, r1.price, r1.expireAt));
        return addr == ecrecover(hash, sig.v, sig.r, sig.s);
    }

    function checkSigBatch2(address addr, Batch2Receipt memory r1, Signature memory sig)
        internal
        pure
        returns (bool)
    {
        bytes32 hash = keccak256(abi.encodePacked(r1.sessionId, r1.from, r1.seed2, r1.sigma_vw, r1.count, r1.price, r1.expireAt));
        return addr == ecrecover(hash, sig.v, sig.r, sig.s);
    }

    function checkSigBatch3(address addr, Batch3Receipt memory r1, Signature memory sig)
        internal
        pure
        returns (bool)
    {
        bytes32 hash = keccak256(abi.encodePacked(r1.sessionId, r1.from, r1.u0_x0_lgs.X, r1.u0_x0_lgs.Y, r1.u0d.X, r1.u0d.Y, r1.price, r1.expireAt));
        return addr == ecrecover(hash, sig.v, sig.r, sig.s);
    }

    function checkSigVRF(address addr, VRFReceipt memory r1, Signature memory sig)
        internal
        pure
        returns (bool)
    {
        bytes32 hash = keccak256(abi.encodePacked(r1.sessionId, r1.from, r1.g_exp_r.X, r1.g_exp_r.Y, r1.price, r1.expireAt));
        return addr == ecrecover(hash, sig.v, sig.r, sig.s);
    }

    /* Public Getters */

    function getRecordBatch1(address _a, address _b, uint256 _sessionId)
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
        Batch1Record memory _record = batch1Records[_a][_b][_sessionId];
        seed0 = _record.seed0;
        seed2 = _record.receipt.seed2;
        k_mkl_root = _record.receipt.k_mkl_root;
        count = _record.receipt.count;
        price = _record.receipt.price;
        expireAt = _record.receipt.expireAt;
        submitAt = _record.submitAt;
    }

    function getRecordBatch2(address _a, address _b, uint256 _sessionId)
        public
        view
        returns (
            bytes32 seed0,
            uint256 submitAt
        )
    {
        Batch2Record memory _record = batch2Records_[_a][_b][_sessionId];
        seed0 = _record.seed0;
        submitAt = _record.submitAt;
    }

    function getRecordBatch3(address _a, address _b, uint256 _sessionId)
        public
        view
        returns (
            uint256 d,
            uint256 x0_lgs,
            uint256 submitAt
        )
    {
        Batch3Record memory _record = batch3Records_[_a][_b][_sessionId];
        d = _record.d;
        x0_lgs = _record.x0_lgs;
        submitAt = _record.submitAt;
    }

    function getRecordVRF(address _a, address _b, uint256 _sessionId)
        public
        view
        returns (
            uint256 r,
            uint256 submitAt
        )
    {
        VRFRecord memory _record = vrfRecords_[_a][_b][_sessionId];
        r = _record.r;
        submitAt = _record.submitAt;
    }

}
