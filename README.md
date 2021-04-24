# LifeFile

A general purpose blockchain highly compatible with Ethereum's ecosystem.

This is the first implementation written in golang.

## Table of contents

* [Installation](#installation)
  * [Requirements](#requirements)
  * [Getting the source](#getting-the-source)
  * [Dependency management](#dependency-management)
  * [Building](#building)
* [Running LFT](#running-lft)
  * [Sub-commands](#sub-commands)
* [Acknowledgement](#acknowledgement)
* [Contributing](#contributing)

## Installation

### Requirements

LFT requires `Go` 1.13+ and `C` compiler to build. To install `Go`, follow this [link](https://golang.org/doc/install).

### Getting the source

Clone the LFT repo:

```shell
git clone https://github.com/lifefile/LFT.git
cd LFT
```

### Dependency management

Simply run:

```shell
make dep
```

If you keep getting network errors, it is suggested to use [Go Module Proxy](https://golang.org/cmd/go/#hdr-Module_proxy_protocol). [https://proxy.golang.org/](https://proxy.golang.org/) is one option.

### Building

To build the main app `lft`, just run

```shell
make
```

or build the full suite:

```shell
make all
```

If no errors are reported, all built executable binaries will appear in folder *bin*.

## Running LFT

Connect to LifeFile's mainnet:

```shell
bin/lft --network main
```

Connect to LifeFile's testnet:

```shell
bin/lft --network test
```

or startup a custom network

```shell
bin/lft --network <custom-net-genesis.json>
```

To show usages of all command line options:

```shell
bin/lft -h
```

* `--network value`             the network to join (main|test) or path to genesis file
* `--data-dir value`            directory for block-chain databases
* `--cache value`               megabytes of ram allocated to internal caching (default: 2048)
* `--beneficiary value`         address for block rewards
* `--target-gas-limit value`    target block gas limit (adaptive if set to 0) (default: 0)
* `--api-addr value`            API service listening address (default: "localhost:8669")
* `--api-cors value`            comma separated list of domains from which to accept cross origin requests to API
* `--api-timeout value`         API request timeout value in milliseconds (default: 10000)
* `--api-call-gas-limit value`  limit contract call gas (default: 50000000)
* `--api-backtrace-limit value` limit the distance between 'position' and best block for subscriptions APIs (default: 1000)
* `--verbosity value`           log verbosity (0-9) (default: 3)
* `--max-peers value`           maximum number of P2P network peers (P2P network disabled if set to 0) (default: 25)
* `--p2p-port value`            P2P network listening port (default: 11235)
* `--nat value`                 port mapping mechanism (any|none|upnp|pmp|extip:&lt;IP&gt;) (default: "none")
* `--bootnode value`            comma separated list of bootnode IDs
* `--skip-logs`                 skip writing event|transfer logs (/logs API will be disabled)
* `--pprof`                     turn on go-pprof
* `--disable-pruner`            disable state pruner to keep all history
* `--help, -h`                  show help
* `--version, -v`               print the version

### Sub-commands

* `solo`                client runs in solo mode for test & dev

```shell
# create new block when there is pending transaction
bin/lft solo --on-demand

# save blockchain data to disk(default to memory)
bin/lft solo --persist

# two options can work together
bin/lft solo --persist --on-demand
```

* `master-key`          master key management

```shell
# print the master address
bin/lft master-key

# export master key to keystore
bin/lft master-key --export > keystore.json


# import master key from keystore
cat keystore.json | bin/lft master-key --import
```

## Acknowledgement

A special shout out to following projects:

* [Ethereum](https://github.com/ethereum)
* [Swagger](https://github.com/swagger-api)

## Contributing

Thank you so much for considering to help out with the source code! We welcome contributions from anyone on the internet, and are grateful for even the smallest of fixes!

Please fork, fix, commit and send a pull request for the maintainers to review and merge into the main code base.

### Forking LFT

When you "Fork" the project, GitHub will make a copy of the project that is entirely yours; it lives in your namespace, and you can push to it.

### Getting ready for a pull request

Please check the following:

* Code must be adhere to the official Go Formatting guidelines.
* Get the branch up to date, by merging in any recent changes from the master branch.

### Making the pull request

1. On the GitHub site, go to "Code". Then click the green "Compare and Review" button. Your branch is probably in the "Example Comparisons" list, so click on it. If not, select it for the "compare" branch.
1. Make sure you are comparing your new branch to master. It probably won't be, since the front page is the latest release branch, rather than master now. So click the base branch and change it to master.
1. Press Create Pull Request button.
1. Provide a brief title.
1. Explain the major changes you are asking to be code reviewed. Often it is useful to open a second tab in your browser where you can look through the diff yourself to remind yourself of all the changes you have made.
