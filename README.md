# Halving has left the building!

Kaspa 2.0 beyond the Blockchain

The world’s first BlockDAG with an emissions schedule that multiplies once a year using smooth monthly increments by a factor of (2)^(1/12).



# KaspaV2d

KaspaV2d was the reference full node KaspaV2 implementation written in Go (golang).



## What is KaspaV2

KaspaV2 is an attempt at a proof-of-work cryptocurrency with instant confirmations and sub-second block times. It is based on [the PHANTOM protocol](https://eprint.iacr.org/2018/104.pdf), a generalization of Nakamoto consensus.



## Requirements

Go 1.18 or later.

## Installation

#### Build from Source

- Install Go according to the installation instructions here:
  http://golang.org/doc/install

- Ensure Go was installed properly and is a supported version:

```bash
$ go version
```

- Run the following commands to obtain and install kasv2d including all dependencies:

```bash
$ git clone
$ cd kasv2d
$ go install . ./cmd/...
```

- KaspaV2d (and utilities) should now be installed in `$(go env GOPATH)/bin`. If you did
  not already add the bin directory to your system path during Go installation,
  you are encouraged to do so now.



## Getting Started

KaspaV2d has several configuration options available to tweak how it runs, but all
of the basic operations work with zero configuration.

```bash
$ kaspav2d
```



## Discord

Join our Discord server using the following link: https://discord.com/invite/kaspav2

## Telegram

Join our Telegram group using the following link: https://t.me/kaspav2

## X

Follow X for more updates link: https://x.com/Kaspa_v2

## Website

https://kaspa-v2.com/