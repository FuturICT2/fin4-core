const AllPurpose = artifacts.require("AllPurpose");
const AllPurposeCapped = artifacts.require("AllPurposeCapped");

module.exports = function(deployer) {
  deployer.deploy(
    AllPurpose, 
    "YESCoin", 
    "YES", 
    "1",
    "0x2EABfB9BfBB12cB16e5f1E521fDBFf79AdA15a6f",
    false, 
    true, 
    false,
    10);
  deployer.deploy(AllPurposeCapped,
    "YESCoin",
    "YES",
    "1",
    "0x2EABfB9BfBB12cB16e5f1E521fDBFf79AdA15a6f",
    false,
    "25",
    true,
    true,
    10);
};
