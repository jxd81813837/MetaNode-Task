// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

contract Voting{
    //候选人得票数存储
    mapping(string=>uint) public candidates;
    //投票人记录周期 避免重复投票
    mapping(address=>uint) public votesRound;
    string[] names;
    uint currentRound;


    // 存储合约创建者
    address public owner;

    //初始化投票人信息
    constructor(){
        owner =msg.sender;
        names =["voteA","voteB","voteC"];
        currentRound=1;//默认第一轮
        resetVotes();
    }

    //给某个候选人投票默认一票
    function vote(string memory name) external {
        require(votesRound[msg.sender]!= currentRound," You have already cast your vote");
        //默认当前轮次
        votesRound[msg.sender] =currentRound;
        candidates[name]+=1;
    }
    // 返回某个候选人得票数
    function getVotes(string memory name) external view returns (uint){
        return candidates[name];
    }
    //重制所有候选人得票数
    function resetVotes() public {
        require(owner ==msg.sender," You have not owner");
        for(uint256 i=0;i <names.length;i++){
            candidates[names[i]]=0;
        }
        // 只需增加轮次号即可"重置"投票状态
        currentRound++;
    }
}