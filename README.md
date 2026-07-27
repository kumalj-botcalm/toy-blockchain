# Toy Blockchain and Ledger Simulator

![Go](https://img.shields.io/badge/Go-1.22+-00ADD8?logo=go)
![Build](https://img.shields.io/badge/build-passing-brightgreen)
![Tests](https://img.shields.io/badge/tests-passing-brightgreen)
![Platform](https://img.shields.io/badge/platform-cross--platform-blue)

A lightweight blockchain simulator written in **Go** that demonstrates the fundamental concepts of blockchain technology.

The project implements deterministic SHA-256 block hashing, Proof-of-Work mining, transaction validation, persistent ledger storage, Merkle Trees, ECDSA digital signatures, concurrent mining, dynamic difficulty adjustment and fork resolution.

This project was developed as part of a **Golang Backend Developer Assessment** and follows idiomatic Go design principles using modular packages, unit testing and a command-line interface.

---

# Features

## Core Features

| Feature | Status |
|---------|--------|
| Genesis Block | ✅ |
| Deterministic SHA-256 Block Hashing | ✅ |
| Proof-of-Work Mining | ✅ |
| Blockchain Validation | ✅ |
| Transaction Ledger | ✅ |
| Account Balances | ✅ |
| Pending Transaction Pool | ✅ |
| Persistent JSON Storage | ✅ |
| Command Line Interface | ✅ |

---

## Advanced Features

| Feature | Status |
|---------|--------|
| Wallet Generation (ECDSA) | ✅ |
| Digital Signatures | ✅ |
| Signature Verification | ✅ |
| Merkle Tree | ✅ |
| Concurrent Mining (Goroutines) | ✅ |
| Automatic Difficulty Retargeting | ✅ |
| Fork Resolution (Longest Valid Chain) | ✅ |
| Configurable Block Size | ✅ |
| Configurable Storage File | ✅ |

---

# Project Architecture

```
                    CLI
                     │
                     ▼
            ┌────────────────┐
            │   Blockchain   │
            └────────────────┘
               │     │     │
               │     │     │
               ▼     ▼     ▼
         Transactions Wallet Miner
               │             │
               ▼             ▼
          Merkle Tree   Proof of Work
               │
               ▼
          SHA-256 Hashing
               │
               ▼
          JSON Persistence
```

---

# Project Structure

```
cmd/
└── blockchain/
    ├── main.go
    ├── init.go
    ├── fund.go
    ├── send.go
    ├── mine.go
    ├── validate.go
    ├── balances.go
    ├── print.go
    └── create_wallet.go

internal/
├── blockchain/
├── crypto/
├── ledger/
├── merkle/
├── miner/
├── storage/
├── transaction/
└── wallet/

data/
└── chain.json

wallets/
```

---

# Requirements

- Go 1.22 or newer

---

# Build

```bash
go build ./...
```

---

# Run All Tests

```bash
go test ./...
```

---

# Static Analysis

```bash
go vet ./...
```

---

# Format Source Code

```bash
go fmt ./...
```

---

# Benchmark

```bash
go test ./internal/miner -bench=.
```

---

# Command Line Usage

## Interactive Menu (Recommended)

The CLI includes an interactive, user-friendly menu by default. Simply run the executable without any arguments to launch it:

```bash
go run ./cmd/blockchain
```

The interactive menu will guide you through all available features and prompt you for any required inputs automatically.

---

## Direct Commands

You can also run commands directly by passing arguments. This is useful for automation or quick execution.

### Initialize Blockchain

```bash
go run ./cmd/blockchain init
```

Using a custom blockchain file:

```bash
go run ./cmd/blockchain --file=mychain.json init
```

---

## Create Wallets

```bash
go run ./cmd/blockchain create-wallet Alice

go run ./cmd/blockchain create-wallet Bob
```

---

## Fund an Account

```bash
go run ./cmd/blockchain fund Alice 100
```

---

## Send Funds

```bash
go run ./cmd/blockchain send Alice Bob 40
```

---

## Mine Pending Transactions

```bash
go run ./cmd/blockchain mine
```

---

## Print Blockchain

```bash
go run ./cmd/blockchain print
```

---

## Show Account Balances

```bash
go run ./cmd/blockchain balances
```

---

## Validate Blockchain

```bash
go run ./cmd/blockchain validate
```

---

# Example Workflow

```
Initialize Blockchain
        │
        ▼
Create Wallets
        │
        ▼
Fund Account
        │
        ▼
Pending Transaction Pool
        │
        ▼
Mine Block
        │
        ▼
Proof-of-Work
        │
        ▼
Merkle Root Generation
        │
        ▼
Block Added
        │
        ▼
Balances Updated
        │
        ▼
Validate Blockchain
```

---

# Example Output

```text
Initializing blockchain...
Blockchain initialized.

Funding transaction added to pending pool.

Mining pending transactions...

Mining completed in 0.532 ms
Current Difficulty : 2

Block mined successfully.

Height : 1
Hash   : 00b0b5df90dbab7b7c9b5e58a9282cb9a8445a74bf6cb61fa3b98d375e277f7c
Nonce  : 537

Blockchain is valid.
```

---

# Design Decisions

## SHA-256 Hashing

SHA-256 is used to produce deterministic and collision-resistant hashes for every block.

---

## Merkle Trees

Transactions inside a block are summarised using a Merkle Root, allowing efficient integrity verification without hashing every transaction individually.

---

## Proof-of-Work

Mining requires discovering a nonce that produces a block hash with a configurable number of leading zeros.

---

## Digital Signatures

Transactions are signed using ECDSA to ensure authenticity and prevent unauthorized spending.

---

## Concurrent Mining

Mining distributes nonce searching across multiple goroutines, improving mining performance while demonstrating Go's concurrency model.

---

## Difficulty Retargeting

Mining difficulty automatically adjusts based on previous mining durations to maintain a consistent mining rate.

---

## Fork Resolution

When multiple blockchain branches exist, the implementation selects the longest valid chain.

---

## Persistent Storage

The blockchain is serialized into JSON, allowing data to persist across program executions.

---

# Configurable Parameters

The following runtime parameters are configurable.

| Parameter | Description |
|-----------|-------------|
| Difficulty | Initial Proof-of-Work difficulty |
| Block Size | Maximum number of transactions per block |
| Storage File | Custom blockchain JSON file using `--file` |

Example:

```bash
go run ./cmd/blockchain --file=mychain.json init
```

---

# Testing

The project includes comprehensive unit tests covering:

- Block Hashing
- Merkle Tree
- Blockchain Validation
- Proof-of-Work Mining
- Digital Signatures
- Wallet Generation
- Ledger Balances
- Persistent Storage
- Fork Resolution
- Difficulty Adjustment

Run all tests:

```bash
go test ./...
```

---

# Optional Features Implemented

- ✅ Digital Signatures
- ✅ Merkle Trees
- ✅ Concurrent Mining
- ✅ Difficulty Retargeting
- ✅ Fork Resolution

All optional features specified in the assessment have been implemented.

---

# Known Limitations

This project is intentionally designed as a **single-node educational blockchain**.

The following production blockchain features are outside the project scope:

- Peer-to-peer networking
- Distributed consensus
- Smart contracts
- Transaction broadcasting
- Network synchronization
- Wallet encryption
- REST API
- Web interface

---

# Future Improvements

Potential future enhancements include:

- REST API
- Web Dashboard
- Peer-to-Peer Networking
- Proof of Stake Consensus
- Multiple Mining Nodes
- Smart Contract Support
- Transaction Broadcasting
- Blockchain Explorer
- Automatic Peer Discovery

---

# Author

**Kumal Jayawardena**



Built using Go.