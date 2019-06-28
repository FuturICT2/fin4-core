# Contract information

## Generate Go Bindings for contracts

To generate the bindings for ERC20.sol, one would write the following:
```sh
solc --bin --abi -o ./compiled --overwrite --allow-paths . ./zeppelin-contracts/token/ERC20/ERC20.sol
abigen --abi compiled/ERC20.abi --pkg ethereum --type ERC20 --out ERC20.go --bin compiled/ERC20.bin
```

## Where to find the contracts

Tokens written by the Fin4 team can be found in `./zeppelin-contracts/fin4`. The rest of the contracts have been written by OpenZeppelin.

## The contracts

`./zeppelin-contracts/fin4/AllPurpose.sol` contains two contracts: `AllPurpose` and `AllPurposeCapped`. `AllPurpose` is a generic ERC20 token which gives the creator the option to make it mintable, burnable, and/or transferable, or not. `AllPurposeCapped` has the same capabilities as `AllPurpose` and is capped.
