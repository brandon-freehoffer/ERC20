// SPDX-License-Identifier: MIT
pragma solidity >=0.4.22 <0.9.0;

abstract contract ERC20Interface { 
    
    function totalSupply() public virtual view returns (uint256);
    function balanceOf(address tokenOwner) public virtual view returns (uint256 balance);
    function allowance(address tokenOwner, address spender) public virtual view returns (uint256 remaining);
    function transfer(address to, uint tokens) public virtual returns (bool success);
    function approve(address spender, uint tokens) public virtual returns (bool success);
    function transferFrom(address from, address to, uint tokens) public virtual returns (bool success);
 

    event Transfer(address indexed from, address indexed to, uint tokens);
    event Approval(address indexed tokenOwner, address indexed spender, uint tokens);
}
abstract contract ApproveAndCallFallBack {
    function receiveApproval(address from, uint256 tokens, address token, bytes memory data) public virtual;
}
