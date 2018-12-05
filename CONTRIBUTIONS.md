# Dependencies
- Golang 1.10 (https://golang.org/doc/install)
- Ganache-cli (https://github.com/trufflesuite/ganache-cli)
- gin (https://github.com/codegangsta/gin)
- Elmlang 0.18.0 (you must not use a newer version)
- MySql server


# Install
```bash
$ git clone https://github.com/FuturICT2/fin4-core.git
$ cd fin4-core
$ go get
$ npm install
$ elm package install

# Setting up database:
```SQL
CREATE DATABASE fin4 default charset utf8;
```
#Env variables
 copying .env_sample file in .env
```bash
# cp .env_sample .env
```
Fil in required env variables

copy `fin4-core/.env_sample` to `fin4-core/.env`  
modify DATA_SOURCE_NAME by adding USERNAME, PASSWORD and DB_NAME


```

#

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
