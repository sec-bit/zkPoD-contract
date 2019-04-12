pragma experimental ABIEncoderV2;

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

    struct PlainRangeProof1ForClaim {
        uint256 seed0;
        PlainRangeReceipt1ForClaim receipt;
        Signature receiptSig;
    }

    struct SessionRecord {
        uint256 seed0;
        PlainRangeReceipt1ForClaim receipt;
        uint256 submitAt;
    }

    struct PlainRangeReceipt1ForClaim {
        uint256 sessionId;
        address from;
        bytes32 seed2;
        bytes32 k_mkl_root;
        uint64 count;
        uint256 price;
        uint256 expireAt;
    }

    enum DepositStatus {
        OK,
        CANCELING,
        CANCELED
    }

    enum BltType {
        PLAIN,
        TABLE
    }

    mapping (bytes32 => Bulletin) public bulletins_;

    // B => A => Deposit
    mapping (address => mapping(address => Deposit)) public buyerDeposits_;

    // A => B => SessionId => Receipt
    mapping (address => mapping(address => mapping(uint256 => SessionRecord))) sessionRecords_;
    
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
            revert();
        }
        require(bulletins_[_bltKey].owner != address(0));
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

    // plain range
    // function submitProof1() public;

    function checkSig1(address addr, PlainRangeReceipt1ForClaim memory r1, Signature memory sig)
        public
        pure
        returns (bool)
    {
        bytes32 hash = keccak256(abi.encodePacked(r1.sessionId, r1.from, r1.seed2, r1.k_mkl_root, r1.count, r1.price, r1.expireAt));
        return addr == ecrecover(hash, sig.v, sig.r, sig.s);
    }

    function submitProof1WaitClaim(PlainRangeProof1ForClaim memory proof1)
        public
    {
        address _from = proof1.receipt.from;
        uint256 _sessionId = proof1.receipt.sessionId;
        require(checkSig1(_from, proof1.receipt, proof1.receiptSig));
        require(now < proof1.receipt.expireAt);
        require(sessionRecords_[msg.sender][proof1.receipt.from][_sessionId].receipt.from == address(0));
        sessionRecords_[msg.sender][_from][_sessionId] = SessionRecord({
            seed0: proof1.seed0,
            receipt: proof1.receipt,
            submitAt: now
        });
    }

    // function verifyPath() public;

    struct G1Point {
        uint256 X;
        uint256 Y;
    }

    function hashOfTwoSha3(bytes32 _x, bytes32 _y)
        public
        view
        returns (bytes32)
    {
        return keccak256(abi.encodePacked(_x, _y));
    }

    function hashOfTwoSha256(bytes32 _x, bytes32 _y)
        public
        view
        returns (bytes32)
    {
        return sha256(abi.encodePacked(_x, _y));
    }

    function log2ub(uint256 _n)
        public
        view
        returns (uint256)   
    {
        if (_n == 1) return 0;
        if (_n % 2 != 0) _n++;
        return 1 + log2ub(_n / 2);
    }

    // // for debug
    // event LogVerifyPath(uint256 _i, uint64 _pos, bytes32 _value);
    // event LogBeforeBytes32(bytes32 _mkl_path_i, bytes32 _value);

    function verifyPathOfK(bytes32 _x, bytes32 _y, uint64 _ij, uint64 _ns, bytes32 _root, bytes32[] memory _mkl_path)
        public
        // view
        returns (bool)
    {
        (bytes32 _value_b) = hashOfTwoSha3(_x, _y);
        uint256 _depth = log2ub(_ns);
        if (_mkl_path.length != _depth) {
            return false;
        }
        bytes32 _value = _value_b;

        uint64 _pos = _ij;
        for (uint256 _i = 0; _i < _depth; _i++) {
            // emit LogBeforeBytes32(_mkl_path[_i], _value);
            if (_pos % 2 != 0) {
                _value = hashOfTwoSha256(_mkl_path[_i], _value);
            } else {
                _value = hashOfTwoSha256(_value, _mkl_path[_i]);
            }
            _pos /= 2;
            // emit LogVerifyPath(_i, _pos, _value);
        }
        return (_value == _root);
    }

    uint64 public s_ = 8;

    function claim(address _a, uint256 _sessionId, uint64 _i, uint64 _j, bytes32[2] memory _tij, bytes32[] memory _mkl_path)
        public
    {
        // loadReceipt();

        SessionRecord memory _sessionRecord = sessionRecords_[_a][msg.sender][_sessionId];
        
        require(verifyPathOfK(_tij[0], _tij[1], _i*s_+_j, _sessionRecord.receipt.count*s_, _sessionRecord.receipt.k_mkl_root, _mkl_path));

        // verifyClaim();
    }

    function submitProof2() public {
    
    }
}
