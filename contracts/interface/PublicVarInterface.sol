pragma solidity ^0.5.0;
interface PublicVarInterface {

    function getEccPubU1(uint256 i) external view returns (uint256 _x, uint256 _y);
}
