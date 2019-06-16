pragma solidity ^0.4.24;

import "../token/ERC20/ERC20Mintable.sol";
import "../token/ERC20/ERC20Burnable.sol";
import "../token/ERC20/ERC20Capped.sol";
import "../token/ERC20/ERC20Pausable.sol";


/**
 * @title AllPurpose
 * @dev Very simple ERC20 Token example, where all tokens are pre-assigned to the creator.
 * Note they can later distribute these tokens as they wish using `transfer` and other
 * `ERC20` functions.
 */
contract AllPurpose is ERC20Mintable, ERC20Burnable, ERC20Pausable {

  string public  name;
  string public  symbol;
  uint8  public  decimals;
  bool   public  isBurnable;
  bool   public  isTransferable;
  bool   public  isMintable;

  bool   private constructing = true;

  uint256 public  INITIAL_SUPPLY = 5;

    /**
   * @dev Burns a specific amount of tokens.
   * @param value The amount of token to be burned.
   */
  function burn(uint256 value) public {
      if (isBurnable == true){
        super.burn(value);
      }
  }

// TODO: untested!
  /**
   * @dev Burns a specific amount of tokens.
   * @param value The amount of token to be burned.
   */
  function burnFrom(address from, uint256 value) public {
      if (isBurnable == true){
        super.burnFrom(from, value);
      }
  }
// TODO: Make sure super modifier works
  function mint(
    address to,
    uint256 value
  )
    public
    returns (bool)
  {
    if (isMintable == true){
      emit Transfer(address(0), to, value);
      return super.mint(to, value);
    }
  }

  function pause() public{
    if (constructing){
      super.pause();
    }
  }

  // see whether capped takes care of itself. Maybe since
  // it's contructor won't accept 0, it will just not start
  // those functions


  /**
   * @dev Constructor that gives msg.sender all of existing tokens.
   */
  constructor(
    string name_,
    string symbol_,
    uint8 decimals_,
    address minter,
    bool isBurnable_,
    bool isTransferable_,
    bool isMintable_)
      public
  {
    name = name_;
    symbol = symbol_;
    decimals = decimals_;
    isBurnable = isBurnable_;
    isTransferable = isTransferable_;
    isMintable = isMintable_;
    if(isTransferable){
      pause();
    }
    _mint(msg.sender, INITIAL_SUPPLY);
    constructing = false;
    _addMinter(minter);
  }
}

contract AllPurposeCapped is ERC20Capped, AllPurpose {
  /**
   * @dev Constructor that gives msg.sender all of existing tokens.
   */
  constructor(
    string name_,
    string symbol_,
    uint8 decimals_,
    address minter,
    bool isBurnable_,
    uint cap_,
    bool isTransferable_,
    bool isMintable_)
      public
      ERC20Capped(cap_)
      AllPurpose(
          name_,
          symbol_,
          decimals_,
          minter,
          isBurnable_,
          isTransferable_,
          isMintable_)
  {
  }
}