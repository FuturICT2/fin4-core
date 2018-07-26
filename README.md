# fin4-core
Finance 4.0 is a multi-dimensional incentive system to manage complex systems and promote a circular sharing economy to create a high quality of life for more people with less resources.

# Prerequisites
- Golang 1.10 (https://golang.org/doc/install)
- Ganache-cli (https://github.com/trufflesuite/ganache-cli)

# Install
```bash
$ git clone https://github.com/FuturICT2/fin4-core.git
$ cd fin4-core
$ cp .env_example .env
$ go get
```

# Running development server
```bash
$ ganache-cli
$ make server-dev
```

# Running tests
```bash
$ make test
$ make test-cover
$ make cover-view
```
