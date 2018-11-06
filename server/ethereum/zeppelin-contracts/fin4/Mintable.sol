pragma solidity ^0.4.24;

import "../token/ERC20/ERC20Mintable.sol";
/**
 * @title SimpleToken
 * @dev Very simple ERC20 Token example, where all tokens are pre-assigned to the creator.
 * Note they can later distribute these tokens as they wish using `transfer` and other
 * `ERC20` functions.
 */
contract Mintable is ERC20Mintable {

  string public  name;
  string public  symbol;
  uint8 public  decimals;

  uint256 public  INITIAL_SUPPLY = 0;

  /**
   * @dev Constructor that gives msg.sender all of existing tokens.
   */
  constructor(string name_, string symbol_, uint8 decimals_, address minter) public {
    name = name_;
    symbol = symbol_;
    decimals = decimals_;
    _addMinter(minter);
    _mint(msg.sender, INITIAL_SUPPLY);
  }
}
