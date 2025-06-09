// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

contract Morge{
    //合并有序数组
    function mergeSortedArray(uint[] memory arr1, uint[] memory arr2) public pure returns (uint[] memory){

        uint  sumLehgth = arr1.length+arr2.length;
        uint[] memory result = new uint[](sumLehgth);
        sumLehgth++;
        uint i;
        uint j;
        uint r;
        //只要一个遍历完了就停止
        while (i< arr1.length && j<arr2.length){
            if(arr1[i]>= arr2[j]){
                result[r]=arr2[j];
                r++;
                j++;
            }else{
                result[r]=arr1[i];
                r++;
                i++;
            }
        }
        //说明没有走完，继续添加后边的数据
        while(i< arr1.length){
            result[r]=arr1[i];
            i++;
            r++;
        }
        //说明没有走完，继续添加后边的数据
        while(j<arr2.length){
            result[r]=arr2[j];
            j++;
            r++;
        }
        return result;
    }
}