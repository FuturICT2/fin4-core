# Dependencies
- Golang 1.10 (https://golang.org/doc/install): on Ubuntu 18.10, just do it via the package manager

export GOROOT=/usr/local/go -> try whereis go, and then see where the root is installed
export GOPATH=$HOME/Projects/goprojects -> folder where fin-4 will be installed in
export PATH=$GOPATH/bin:$GOROOT/bin:$PATH
export GOBIN=$GOPATH/bin


- Ganache-cli (https://github.com/trufflesuite/ganache-cli): npm install -g ganache-cli
- gin (https://github.com/codegangsta/gin): go get github.com/codegangsta/gin
- Elmlang 0.18.0 (you must not use a newer version): npm install -g elm@elm0.18.0
- MySql server: apt-get insall mysql-server


# Install
```bash
$ git clone https://github.com/FuturICT2/fin4-core.git
$ cd fin4-core
$ go get
$ npm install
$ elm package install
```
# Setting up database:

Next, you need to create a new database called fin4. For this, enter mysql and execute the following command:
```SQL
CREATE DATABASE fin4 default charset utf8;
```
This creates a new database. Assign an user full access right to the fin4 database (or leave it with root).

# Env variables
 Copy .env_sample file to .env.
 Then, modify the following entries:
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
