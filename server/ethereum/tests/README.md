# Solidity tests

## Set up
Your blockchain settings belong in `truffle-config.js`. The current settings are configured to run on a local blockchain running with ganache (`127.0.0.1:7545`). Please change accordingly for your own setup. More information can be found on [the official documentation](https://www.trufflesuite.com/docs/truffle/reference/configuration).
 
## Run
To run the tests, execute the following:

```
npm install
npm -g install truffle
truffle compile
truffle deploy
truffle test
```

(Make sure your blockchain is running before `truffle deploy` and `truffle test`)

## Info
- `contracts/`: Directory for [Solidity contracts](https://www.trufflesuite.com/docs/truffle/getting-started/interacting-with-your-contracts)
- `migrations/`: Directory for [scriptable deployment files](https://www.trufflesuite.com/docs/truffle/getting-started/running-migrations#migration-files)
- `test/`: Directory for test files for [testing your application and contracts](https://www.trufflesuite.com/docs/truffle/testing/testing-your-contracts)
- `truffle-config.js`: Truffle [configuration file](https://www.trufflesuite.com/docs/truffle/reference/configuration)
- More can be read about the truffle tests on [truffle's js testing documentation](https://www.trufflesuite.com/docs/truffle/testing/writing-tests-in-javascript).

## Known errors
- Strange behaviour: Try deleting your `./build` folder, and compiling and deploying your contracts again. 