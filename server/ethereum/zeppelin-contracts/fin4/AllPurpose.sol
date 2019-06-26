pragma solidity ^0.4.24;

import "../token/ERC20/ERC20Mintable.sol";
import "../token/ERC20/ERC20Burnable.sol";
import "../token/ERC20/ERC20Capped.sol";
import "../token/ERC20/ERC20Pausable.sol";


/**
 * @title AllPurpose
 */
contract AllPurpose is ERC20Mintable, ERC20Burnable, ERC20Pausable {

  string public  name;
  string public  symbol;
  uint8 public  decimals;

  bool   public  isBurnable;
  bool   public  isTransferable;
  bool   public  isMintable;

  bool   private constructing = true;
  uint256 public  INITIAL_SUPPLY = 10;


  function burn(uint256 value) public {
      require(isBurnable, "Coin not burnable");
      super.burn(value);
  }

  function burnFrom(address from, uint256 value) public {
      require(isBurnable, "Coin not burnable");
      super.burnFrom(from, value);
  }
  
  function mint(
    address to,
    uint256 value
  )
    public
    returns (bool)
  {
    require(isMintable, "Coin not mintable");
    return super.mint(to, value);
  }

  function pause() public{
    require(constructing, "this function can only be run on creation");
    super.pause();
  }

  function unpause() public{
    require(constructing, "this function can only be run on creation");
  }

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
    if(!isTransferable_){
      pause();
    }
    _addMinter(minter);
    _mint(msg.sender, INITIAL_SUPPLY);
    constructing = false;
  }
}

contract AllPurposeCapped is ERC20Capped, AllPurpose {

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