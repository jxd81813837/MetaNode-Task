// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

contract Reverse{

    //反转字符串
    function reverseString(string memory str) external  pure returns (string memory){
        bytes memory strByte = bytes(str);
        bytes memory resultByte = new bytes(strByte.length);

        for (uint256 i=0;i < strByte.length;i++){
            resultByte[strByte.length-i-1] =strByte[i];
        }
        return  string(resultByte);
    }
}