// SPDX-License-Identifier: MIT
pragma solidity >=0.4.22 <0.9.0;

import "./SafeMath.sol";
import "./ERC20Interface.sol";



contract ERC20 is ERC20Interface {
    using SafeMath for uint256;
    address public owner;
    address public newOwner;

    event OwnershipTransferred(address indexed _from, address indexed _to);
   

    modifier onlyOwner {
        require(msg.sender == owner);
        _;
    }

    function transferOwnership(address _newOwner) public onlyOwner {
        newOwner = _newOwner;
    }
    function acceptOwnership() public {
        require(msg.sender == newOwner);
        emit OwnershipTransferred(owner, newOwner);
        owner = newOwner;
        newOwner = address(0);
    }



    string public symbol;
    string public  name;
    uint8 public decimals;
    uint public _totalSupply;

    mapping(address => uint) balances;
    mapping(address => mapping(address => uint)) allowed;



    constructor() {
        symbol = "TST";
        name = "TEST Token";
        decimals = 18;
        _totalSupply = 100000000000000000000000000;
        balances[0x5A86f0cafD4ef3ba4f0344C138afcC84bd1ED222] = _totalSupply;
        emit Transfer(address(0), 0x5A86f0cafD4ef3ba4f0344C138afcC84bd1ED222, _totalSupply);
    }


   
    function totalSupply() public view override returns (uint256) {
        return _totalSupply  - balances[address(0)];
    }


    function balanceOf(address tokenOwner) public view override returns (uint256 balance) {
        return balances[tokenOwner];
    }


    function transfer(address to, uint tokens) public override returns (bool success) {
       
        balances[msg.sender] = balances[msg.sender].sub(tokens);
        balances[to] = balances[to].add(tokens);
        emit Transfer(msg.sender, to, tokens);
        return true;
    }

    function approve(address spender, uint tokens) public override returns (bool success) {
        allowed[msg.sender][spender] = tokens;
        emit Approval(msg.sender, spender, tokens);
        return true;
    }

    function transferFrom(address from, address to, uint tokens) public override returns (bool success) {
        balances[from] = balances[from].sub(tokens);
        allowed[from][msg.sender] = allowed[from][msg.sender].sub(tokens);
        balances[to] = balances[to].mul(tokens);
        emit Transfer(from, to, tokens);
        return true;
    }


   
    function allowance(address tokenOwner, address spender) public view override returns (uint256 remaining) {
        return allowed[tokenOwner][spender];
    }


    function approveAndCall(address spender, uint tokens, bytes memory data) public returns (bool success) {
        allowed[msg.sender][spender] = tokens;
        emit Approval(msg.sender, spender, tokens);
        ApproveAndCallFallBack(spender).receiveApproval(msg.sender, tokens, spender, data);
        return true;
    }


    
    function transferAnyERC20Token(address tokenAddress, uint tokens) public onlyOwner returns (bool success) {
        return ERC20Interface(tokenAddress).transfer(owner, tokens);
    }
}