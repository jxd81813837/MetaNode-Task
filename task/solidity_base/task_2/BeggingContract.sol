// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "@openzeppelin/contracts/access/Ownable.sol";

// 合约部署地址 https://sepolia.etherscan.io/tx/0xf51636d615529b50df0ae34644beb47cba0ea2d0674085fc9a175dd1f9e83626
//合约地址 0x26bb001397d3170b914ef71439e052e6EE121dA9
//发布地址 0x8aaCA58C897F28d2b0DedeBe63cB6eef03EA3E9B

//✅ 作业3：编写一个讨饭合约
contract BeggingContract is Ownable {
    //记录捐赠金额
    mapping(address=>uint256)  donates;

    // 新增：时间设置事件
    event DonationPeriodSet(uint256 start, uint256 end);

    //新增：捐赠事件，记录捐赠地址和金额
    event Donation(address indexed addr, uint256 amount);

    //前三名维护地址数组
    address[] topThree;

    uint256 public donateStartTme;
    uint256 public donateEndTime;
    //校验区块时间
    modifier donatePeriod(){
        require(block.timestamp>=donateStartTme&&block.timestamp<donateEndTime,"not donatePeriod ,time is stop");
        _;
    }

    constructor(address _owner) Ownable(_owner) {
        topThree = new address[](3);
    }

    //捐赠信息记录
    function donate() public payable donatePeriod(){
        //对金额进行累加
        uint256 sum = donates[msg.sender]+msg.value;
        donates[msg.sender]=sum;

        //发布捐赠事件
        emit Donation(msg.sender, msg.value);

        //考虑到大部分捐赠都是小于第三名 所有从后向前遍历
        uint256 i=topThree.length-1;
        //获取最后一名金额
        uint256 thirdAmount = donates[topThree[i]];
        int alreadyIndex =-1;

        //找到是否存在
        for(uint256 j=0;j<topThree.length-1;j++){
            if(topThree[j]==msg.sender){
                alreadyIndex=int(i);
                break;
            }
        }

        //如果不存在且小于最后一名不进入循环
        if(alreadyIndex < 0 &&thirdAmount>=sum){
            return;
        }
        //说明这个捐赠者存在 前三名中
        if(alreadyIndex>=0){
            // 说明 他捐赠的金额还达不到排位调整的程度，不需要移动
            // 如果alreadyIndex =0 说明他本来就是第一 也没必要移动
            if(alreadyIndex == 0||donates[topThree[i]]>=sum){
                return;
            }else {
                //说明需要调整 ，把他从排位中清除掉，然后重新计算排名，最后一位设定成0地址
                //这里注意一定是j<topThree.length-1  因为最后一位不用管 会被赋值为0地址
                for(uint256 j=uint256(alreadyIndex);j<topThree.length-1;j++){
                    topThree[j]=topThree[j+1];
                }
                topThree[topThree.length-1]=address(0);
            }
        }
        while(true){
            //如果等于0 说明，当前捐赠者是第一名
            if(i==0){
                //重新排序 从最后开始依次互换
                for(uint j=topThree.length-1 ;j>i ;j--){
                    topThree[j]=topThree[j-1];
                }
                topThree[i]=msg.sender;
                break;
            }
            if(donates[topThree[i]]<=sum && donates[topThree[i-1]]>sum){
                //重新排序 从最后开始依次互换
                for(uint j=topThree.length-1 ;j>i ;j--){
                    topThree[j]=topThree[j-1];
                }
                topThree[i]=msg.sender;
                break;
            }
            i--;
        }

    }
    //合约所有者提取所有金额
    function withdraw() external payable onlyOwner{
        uint256 balance = address(this).balance;
        require(balance!=0,"balance is zero");
        payable (owner()).transfer(balance);
    }

    //查询金额依据地址
    function getDonation (address addr) view external returns (uint256){
        return donates[addr];
    }

    //显示捐赠前三名
    function donateTopThree() view external returns (address[]memory){

        return topThree;
    }

    //设置周期
    function setDonationPeriod (uint256 start, uint256 end) external {
        require(start<end,"Invalid time range ");
        donateStartTme = start;
        donateEndTime = end;
        emit DonationPeriodSet(start,end);
    }

    //block.timestamp; 这个时间和北京时间相差8小时 只得是主网运行时间
    function gettimestamp () external view returns (uint256){
        return block.timestamp;
    }
}