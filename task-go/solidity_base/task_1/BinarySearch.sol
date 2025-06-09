// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.10;

contract BinarySearch {
   //[1,3,7,8,15,99,101,102],99
   function binarySearch(uint[] memory arr,uint taget) public pure returns (int){
      uint left =0;
      uint right = arr.length-1;
      while(left <= right){
         uint mid = (left+ right)/2;
         if(taget > arr[mid]){//说明在右半段一段
            left = mid+1;
         }else if(taget == arr[mid]){
            return  int(mid);
         } else{
            right =mid-1;
         }
      }
      return -1;
   }
}
