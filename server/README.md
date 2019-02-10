

# Fin4 Server
This repo contains all files of the finance 4.0 server.


## High level overview
The server currently has three main purposes
- routing API requests from clients to the right destination
- Business logic
  - storing/ caching tokens in a mysql database
  - handling user information
- Communication with the blockchain

### APIs
Communication with the server is done via http. The entry point for all http requests is the routes/routes.go file.

The following methods currently exist


| http Method       | URL           | body  | purpose |
| ------------- |:-------------:| :-----| ------|
| GET | /api/status | - | returns the status of the server| 
| POST      | /wapi/register | "name":String,<br />"email":String, <br />"password":String,<br /> "agreeToTerms":boolean,<br />"isFastSignup":boolean | Registers a user in the fin4 system|
|POST      | /wapi/login      |   "name":String,<br />"email":String, <br />"password":String, | Logs a user in. The returned session cookie needs to be used for all /wapi/... methods, except for the register method|
| GET |  /wapi/session     |  -   | Gets the current seesion|
| GET | /wapi/tokens | - | Fetches all fin4 tokens |


In order to start communicating with the server one needs to perform the following API requests
- 



# Documentation

check out: https://godoc.org/github.com/FuturICT2/fin4-core/server/

<!--
markdown syntax https://help.github.com/articles/page-build-failed-markdown-errors/
-->
