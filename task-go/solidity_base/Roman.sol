// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.10;


contract Roman{
    mapping (bytes1=> int) romanMap;
    constructor(){
        romanMap["I"]=1;
        romanMap["V"]=5;
        romanMap["X"]=10;
        romanMap["L"]=50;
        romanMap["C"]=100;
        romanMap["D"]=500;
        romanMap["M"]=1000;

    }
    //罗马数字转换为整数
    function romanToInt(string memory str) public view returns (int){
        bytes memory strBytes =  bytes(str);
        int result;
        int pre;//之前的数据
        for (uint256 i = 0; i < strBytes.length; i++) {
            //这个方法能获取到字符串中的每个资方
            bytes1 charStr =strBytes[strBytes.length-i-1];
            int current = romanMap[charStr];
            if(current>=pre){
                result += current;
            }else{// 特殊情况
                result -=current;
            }
            pre =current;
        }
        return result;
    }
    //整数数字转换为罗马数字
    function intToRoman(uint16  num) public pure returns (string memory) {
        require(num > 0 && num < 4000, "Number must be 1-3999");
        uint16[13] memory values =[1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1];
        string[13] memory romans =["M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"];
        bytes memory result =new bytes(0);
        for(uint256 i=0; i<values.length ;i++){
            while(num >= values[i]){
                //获取到数字对应的罗马数组
                bytes  memory romanBytes = bytes(romans[i]);
                bytes  memory resultNew =appendStr(result,romanBytes);
                //追加完成 重新赋值
                result= resultNew;
                num = num -values[i];
            }
        }
        return string(result);
    }
    //旧版
    function oldIntToRoman(uint16  num) public pure returns (string memory) {
        require(num > 0 && num < 4000, "Number must be 1-3999");
        uint16[13] memory values =[1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1];
        string[13] memory romans =["M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"];
        bytes memory result =new bytes(0);
        for(uint256 i=0; i<values.length ;i++){
            if(num >= values[i]){
                //获取到数字对应的罗马数组
                bytes  memory romanBytes = bytes(romans[i]);
                bytes  memory resultNew =appendStr(result,romanBytes);
                //追加完成 重新赋值
                result= resultNew;
                num = num -values[i];
                //因为有时候需要多次检查比如800 -500 等于300 要多次减去100
                //单这样做 又时候会造成多一次循环
                if(num >= values[i]){
                    i--;
                }
            }
        }
        return string(result);
    }
    // str1+str2 输出
    function appendStr(bytes memory str1Bytes, bytes memory str2Bytes) internal  pure returns (bytes memory){
        bytes  memory result =new bytes(str1Bytes.length+str2Bytes.length);
        //先把str1Bytes复制 到result
        for(uint256  i=0;i<str1Bytes.length ;i++){
            result[i]=str1Bytes[i];
        }
        for(uint256  i=0;i<str2Bytes.length ;i++){
            result[str1Bytes.length+i]=str2Bytes[i];
        }
        return result;
    }
}