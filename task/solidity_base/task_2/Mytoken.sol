// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;
//✅ 作业 1：ERC20 代币
contract Mytoken {
    //记录转账操作事件
    event Transfer(address indexed from,  address indexed to, uint256  amount);
    //记录授权操作事件
    event Approval(address indexed approveAddr, address indexed spender, uint256  limitAmount);

    //设置一个存储数据
    mapping (address=>uint256) private balances;
    //设置一个存储数据
    mapping (address=>mapping (address=>uint256)) private approves;
    //账户管理者地址
    address  private owner;
    //代币名称
    string private name;
    //代币简称
    string private symbol;

    //代币总量
    uint256 public total;

    constructor(string memory name_,string memory symbol_){
        name=name_;
        symbol=symbol_;
        owner=msg.sender;
    }
    //增发代币
    function _mint(uint256 count) external  returns(uint256){
        require(owner==msg.sender,"Is must owner mint token!!!");
        total +=count;
        balances[owner]+=count;
        //来源于0地址 增发监听事件
        emit  Transfer(address(0),owner,count);
        return total;
    }
    //查询目前总发行量
    function getTotal () external view returns (uint256){
        return total;
    }
    //查询余额
    function getBalanceOf(address addr)  public view returns (uint256){
        return  balances[addr] ;
    }

    //转账
    function transfer( address to, uint256 amount) public returns (bool){
        require(to!= address(0),"cannot transfer to zero address");
        require(balances[msg.sender]>=amount,"Insufficient Balance!!!");
        balances[msg.sender] -=amount;
        balances[to] +=amount;
        emit Transfer(msg.sender,to,amount);
        return true;
    }

    //授权
    function approve(address spender, uint256 limitAmount) public returns (bool){
        require(spender!= address(0),"cannot approve to zero address");
        approves[msg.sender][spender]=limitAmount;
        emit Approval(msg.sender, spender, limitAmount);
        return true;
    }

    //授权转账
    function transferFrom(address from,address to, uint256 amount)public returns (bool){
        require(to!= address(0),"cannot transfer to zero address");
        require(approves[from][msg.sender]>= amount,"Insufficient credit limit!!!");
        require(balances[from]>=amount,"Insufficient Balance!!!");
        approves[from][msg.sender]-=amount;//授权额度减去相映金额
        balances[from]-=amount;
        balances[to] +=amount;
        emit Transfer(from,to,amount);
        return true;
    }

    function getName() external view returns(string memory){
        return name;
    }

    function getSymbol() external view returns(string memory){
        return symbol;
    }
}