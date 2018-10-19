[![CircleCI](https://circleci.com/gh/FuturICT2/fin4-core/tree/master.svg?style=svg&circle-token=fe8beee27987a1dd0a05f68f1fdef4ca17051a14)](https://circleci.com/gh/FuturICT2/fin4-core/tree/master)

# fin4-core
Finance 4.0 is a multi-dimensional incentive system to manage complex systems and promote a circular sharing economy to create a high quality of life for more people with less resources.

# Dependencies
- Golang 1.10 (https://golang.org/doc/install)
- Ganache-cli (https://github.com/trufflesuite/ganache-cli)
- gin (https://github.com/codegangsta/gin)
- Elmlang
- MySql server

### setting up database:
```SQL
CREATE DATABASE fin4 default charset utf8;
```
copy `fin4-core/.env_sample` to `fin4-core/.env`  
modify DATA_SOURCE_NAME by adding USERNAME, PASSWORD and DB_NAME

# Install
```bash
$ git clone https://github.com/FuturICT2/fin4-core.git
$ cd fin4-core
$ cp .env_sample .env
$ go get
$ npm install
$ elm package install
```

# Running development server
```bash
$ ganache-cli
$ make server-dev
```
# Running development web front-end
```bash
$ make client-dev
```

# Running tests
```bash
$ make test
$ make test-cover
$ make cover-view
```
