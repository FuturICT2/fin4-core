var assert = require('assert');
const AllPurpose = artifacts.require("AllPurpose");
const AllPurposeCapped = artifacts.require("AllPurposeCapped");
var apMint,apBurn,apNoTransfer;
const initialSupply = 10;
const cap = "30";
const address = "0x2EABfB9BfBB12cB16e5f1E521fDBFf79AdA15a6f";

// TODO to test:
// - can't be minted from non-minter account
// - can't be burnt from non-burner account
// - can't be paused if unpausable
// - can't be unpaused when paused (make transferable after untransferable)

beforeEach(async () => {
  apMint = await AllPurpose.new(
    "YESCoin",  // name
    "YES",      // symbol
    "1",        // decimals
    address,    // minter
    false,      // isBurnable
    true,      // isTransferable
    true,     // isMintable
    initialSupply);      // initialSupply
  apBurn = await AllPurpose.new(
    "YESCoin", 
    "YES", 
    "1", 
    address,
    true,
    true,
    false,
    initialSupply);
  apNoTransfer = await AllPurpose.new(
    "YESCoin", 
    "YES", 
    "1", 
    address,
    false,
    false,
    false,
    initialSupply);
  apCapped = await AllPurposeCapped.new(
    "YESCoin",    // name
    "YES",        // symbol
    "1",          // decimals
    address,      // minter
    false,        // isBurnable
    cap,          // cap
    true,        // isTransferable
    true,        // isMintable
    initialSupply);
  })

contract("AllPurpose: Unmintable, unburnable, transferable, uncapped", accounts => {
  it("should not burn", () => 
    AllPurpose.deployed()
      .then(async instance => {
        assert.rejects( async() => 
          await instance.burn("1")
        );
      }) 
  );
  it("should not mint", () =>
    AllPurpose.deployed()  
      .then(async instance => 
        assert.rejects( async() => 
          await instance.mint(accounts[1], "3")
      )
    )
  );
  it("should be transferable", () =>
    AllPurpose.deployed()
      .then(async instance => {
        await assert.doesNotReject( async () =>
          instance.transfer(accounts[1], "3")
        );   
      })
  );
});


contract("AllPurpose: MINTABLE, unburnable, transferable, uncapped", accounts => {
  it("should not burn", () => 
    AllPurpose.deployed()
      .then(async instance => {
        assert.rejects( async() => 
          await instance.burn("1")
        );
      }) 
  );
  it("should mint", () =>
    AllPurpose.at(apMint.address)  
      .then(async instance => {
        const bAfter = await instance
                              .mint(accounts[1], "3")
                              .then(() => instance
                                            .balanceOf(accounts[1])
                                            .then(result => result.toNumber()));
        await assert.deepStrictEqual(
          bAfter,
          3);
      })
  );
  it("should be transferable", () =>
    AllPurpose.at(apMint.address)  
      .then(async instance => {
        await assert.doesNotReject( async () =>
          instance.transfer(accounts[1], "3")
        );
      })
  );
});

contract("AllPurpose: unmintable, BURNABLE, transferable, uncapped", accounts => {
  it("should burn", () => 
    AllPurpose.at(apBurn.address)  
      .then(async instance => {
        const bAfter = await instance
                              .burn("1")
                              .then(() => instance
                                            .balanceOf(accounts[0])
                                            .then(result => result.toNumber()));
        await assert.deepStrictEqual(
          bAfter,
          initialSupply - 1, 
          "token burnt");
    }) 
  );
  it("should not mint", () =>
    AllPurpose.deployed()  
      .then(async instance => 
        assert.rejects( async() => 
          await instance.mint(accounts[1], "3")
      )
    )
  );
  it("should be transferable", () =>
    AllPurpose.at(apBurn.address)  
      .then(async instance => {
        await assert.doesNotReject( async () =>
          instance.transfer(accounts[1], "3")
        );
      })
  );
});

contract("AllPurpose: Unmintable, unburnable, UNTRANSFERABLE, uncapped", accounts => {
  it("should not burn", () => 
    AllPurpose.deployed()
      .then(async instance => {
        assert.rejects( async() => 
          await instance.burn("1")
        );
      }) 
  );
  it("should not mint", () =>
    AllPurpose.deployed()  
      .then(async instance => 
        assert.rejects( async() => 
          await instance.mint(accounts[1], "3")
      )
    )
  );
  it("should NOT be transferable", () =>
  AllPurpose.at(apNoTransfer.address)  
      .then(async instance => {
      
        await assert.rejects( async () =>
          instance.transfer(accounts[1], "3")
        );          
      })
  );
});

contract("AllPurpose: Unmintable, unburnable, transferable, CAPPED", accounts => {
  it("should not burn", () => 
    AllPurpose.deployed()
      .then(async instance => {
        assert.rejects( async() => 
          await instance.burn("1")
        );
      }) 
  );
  it("should mint", () =>
  AllPurposeCapped.at(apCapped.address)
      .then(async instance => {
        const balance = await instance
                              .mint(accounts[0], "1")
                              .then(() => instance
                                            .balanceOf(accounts[0])
                                            .then(result => result.toNumber()));
        await assert.deepStrictEqual(
          balance,
          initialSupply + 1);
      })
  );
  it("should not mint more than cap", () =>
  AllPurposeCapped.at(apCapped.address)
      .then(async instance => {
        assert.rejects( async() => 
          await instance.mint(accounts[1], cap + 1));                   
      })
  );
  it("should be transferable", () =>
    AllPurposeCapped.at(apCapped.address)
      .then(async instance => {
        await assert.doesNotReject( async () =>
          instance.transfer(accounts[1], "3")
        );   
      })
  );
});
