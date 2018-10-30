# Generate Go Bindings for contracts

```sh
solc --bin --abi -o ./compiled --allow-paths . ./contracts/token/ERC20/ERC20.sol
abigen --abi compiled/ERC20.abi --pkg ethereum --type ERC20 --out ERC20.go --bin compiled/ERC20.bin
```
