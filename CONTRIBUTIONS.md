We assume a fresh Ubuntu 18.10 system for the following guide.

# Dependencies
- Golang 1.10 (https://golang.org/doc/install): install via package manager

- configure env variables:
```bash
export GOROOT=/usr/local/go -> try whereis go, and then see where the root is installed
export GOPATH=$HOME/Projects/goprojects -> folder where fin-4 will be installed in
export PATH=$GOPATH/bin:$GOROOT/bin:$PATH
export GOBIN=$GOPATH/bin
```

- Ganache-cli (https://github.com/trufflesuite/ganache-cli): npm install -g ganache-cli
- gin (https://github.com/codegangsta/gin): go get github.com/codegangsta/gin
- Elmlang 0.18.0 (you must not use a newer version): npm install -g elm@elm0.18.0
- MySql server: apt-get insall mysql-server

# Install Finance 4.0 demonstrator
```bash
$ git clone https://github.com/FuturICT2/fin4-core.git
$ cd fin4-core
$ go get
$ npm install: install server dependencies
$ elm package install: install elm web app dependencies
```
# Setting up database:
sudo apt-get install mysql-server
sudo mysql_secure_installation
set the root password
sudo mysql -u root -p
<enter password>

Inside mysql, create a new database called fin4.
```SQL
CREATE DATABASE fin4 default charset utf8;
```
Assign a user full access right to the fin4 database (or leave it with root).

# Env variables
 Copy .env_sample file to .env.
 Then, modify the following entries:
```bash
cp `fin4-core/.env_sample` `fin4-core/.env`
```
Open the .env file and add the following information:

modify DATA_SOURCE_NAME by adding USERNAME, PASSWORD and DB_NAME

# Running the development-server
```bash
$ ganache-cli # alternatively, you can use the Ganache application
$ make server-dev
```
# Running the development-client (web front-end)
```bash
$ make client-dev
```
# Testing
```bash
$ make test
$ make test-cover
$ make cover-view
```
