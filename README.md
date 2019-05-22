[![CircleCI](https://circleci.com/gh/FuturICT2/fin4-core/tree/master.svg?style=svg&circle-token=fe8beee27987a1dd0a05f68f1fdef4ca17051a14)](https://circleci.com/gh/FuturICT2/fin4-core/tree/master)


# Finance 4.0
Finance 4.0 aims for a multi-dimensional incentive system to motivate sustainable behavior. Communities can creat cryptographic tokens to incentivize certain behavior, e.g. collecting waste, avoiding C02, helping others, etc. Users can obtain these tokens by proving that they performed such actions. The number of different token types is unlimited, leading over time to a multi-dimensional system of different incentives -- smart coordination through "sustainable money". In the same spirit, the system uses tokens for democratic governance. 
A live instance can be found here: www.finfour.net .
Main concepts of the Finance 4.0 system are explained in [Decrypting Distributed Ledger Systems](https://arxiv.org/abs/1811.03419). This document is a good starter for everyone who wants to know
- What is a cryptoeconomic token?
- What is the difference between Ethereum, Bitcoin and Hyperledger?
- What are actions in a distributed ledger technology (DLT) system?
- What are claims?
- How do cryptoeconomic tokens have value?


To learn more about the finance 4.0 system:

Short video: [Youtube - Distributed Sustainability](https://www.youtube.com/watch?v=oNlKdHjvExo)

Social media: [Twitter](https://twitter.com/futurict2)

Concept paper: [Finance4.0](https://futurict2.eu/finance-4-0-concept-wp3-interim-report-m12-february-2018/)

Slack: todo

# Content

- [ Fin4 Demonstrator. ](#demonstrator)
- [ Installation. ](#installation)
- [ Documentation.](#documentation)
- [ Contribution. ](#contribution)


<a name="demonstrator"></a>
# Fin4 Demonstrator | This repository
This repository contains a backend server (written in Golang) and a web front-end (written in Elm) of the Finance 4.0 demonstrator. The ELM web app enables a user to 

* Create a new Type of cryptoeconomic Token, which is deployed on the server and Ethereum Blockchain
* Obtain units of created tokens via performing an action, such as planting a tree
* Performing Oracle tasks in order to approve actions

The server provides
* APIs for clients, such as the ELM app
* Communication functionalities with the Ethereum Blockchain
* Smart contract deployer
* Functionality to store actions

The smart contracts are written in Solidity for the Ethereum platform.

In the following Section we explain how the web app and server can be set up.

<a name="installation"></a>
# Install Finance 4.0 demonstrator
We assume a fresh Ubuntu 18.10 system for the following.

## Dependencies
- Golang 1.10 (https://golang.org/doc/install): install via package manager

- configure env variables:
```bash
export GOROOT=/usr/local/go -> try whereis go, and then see where the root is installed
export GOPATH=$HOME/Projects/goprojects -> folder where fin-4 will be installed in
export PATH=$GOPATH/bin:$GOROOT/bin:$PATH
export GOBIN=$GOPATH/bin
```
Best is to store this in ~/.bashrc, so that this automatically set whenever you start a new terminal session.

- Ganache-cli (https://github.com/trufflesuite/ganache-cli): npm install -g ganache-cli
-- gin is a simple command line utility for live-reloading Go web applications.
- gin (https://github.com/codegangsta/gin): go get github.com/codegangsta/gin
- Elmlang 0.18.0 (you must not use a newer version): npm install -g elm@elm0.18.0
- MySql server: apt-get insall mysql-server

## Set up
```bash
$ git clone https://github.com/FuturICT2/fin4-core.git
$ cd fin4-core
$ go get
$ npm install: install server dependencies
$ elm package install: install elm web app dependencies
```
## Setting up database:
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

```SQL
CREATE USER 'newuser'@'localhost' IDENTIFIED BY 'password1';
GRANT ALL PRIVILEGES ON * . * TO 'newuser'@'localhost';
FLUSH PRIVILEGES;
```

## Env variables
 Copy .env_sample file to .env.
 Then, modify the following entries:
```bash
cp `fin4-core/.env_sample` `fin4-core/.env`
```
Open the .env file and add the following information:

modify DATA_SOURCE_NAME by adding USERNAME, PASSWORD and DB_NAME. Assuming a root user for the fin4 database and the password "password1" the expression should look as follow:
```bash
$ export DATA_SOURCE_NAME="root:password1@/fin4?charset=utf8mb4,utf8&parseTime=true"
```
## Running the development-server
```bash
$ export .env # load the environment variables - Do not forget this whenever you open a new terminal
$ ganache-cli # simulates the ethereum blockchain. Alternatively, you can use the Ganache application
$ make server-dev # deploy the server
```
## Running the development-client (web front-end)
```bash
$ make client-dev
```
After the build, open https://localhost:8888/ inside of your favourite browser. 
## Testing
```bash
$ make test
$ make test-cover
$ make cover-view
```

<a name="documentation"></a>
# Documentation
In the following we document the basic execution of the server and the elm app are explained. Detailed information can/ will be found in the respective packages.

## fin4 Server
The main functioning of the server is explained in the [server](https://github.com/FuturICT2/fin4-core/tree/master/server) package.
### Entry point - main.go
When executing
```bash
make server-dev
```
A new server instance is exportet to the port 8181, which listens to port 3000 (hence all http request will need to be send to this port). You dont know what ports are? Find brief example explanation [ here ](https://www.quora.com/What-are-ports-in-computer-programming-the-source-and-destination-ports-why-do-we-need-them ).
In particular, main.go is executed which
- loads the necessary environment variables into the program
- connects to the local database (which you specified in the .env file)
- initiates the token and user services
- connects with AWS (is used in our live instance, which is deployed to www.finfour.net)
- connects with an instance of Ethereum
- start listening to http request (which internally uses the gin library)

## fin4 webapp


<a name="contribution"></a>
# Contribute!
We are looking for collaboration from the Open Source community. The project needs enhancements regarding code quality, documentation, and testing. To contribute to the project please take a look at [open issues](https://github.com/FuturICT2/fin4-core/issues).
<!--
[Here](CONTRIBUTIONS.md) you can find more details on the architecture of the code and how to run the development environment on your local machine.
-->
<!--
markdown syntax https://help.github.com/articles/page-build-failed-markdown-errors/
-->
