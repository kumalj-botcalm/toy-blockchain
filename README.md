# Toy Blockchain and Ledger Simulator

A command-line blockchain implementation written entirely in Go.

This project was developed as part of the Golang Backend Developer Assessment. It demonstrates the core concepts of blockchain technology including proof-of-work mining, transaction validation, ledger management, persistent storage, blockchain validation, digital signatures, Merkle trees, concurrent mining and fork resolution.

---

# Features

## Core Features

- Genesis Block
- Deterministic SHA-256 Block Hashing
- Proof-of-Work Mining
- Blockchain Validation
- Transaction Ledger
- Account Balances
- Pending Transaction Pool
- Persistent JSON Storage
- Command Line Interface

## Advanced Features

- ECDSA Wallet Generation
- Digital Transaction Signatures
- Signature Verification
- Merkle Tree Support
- Concurrent Mining using Goroutines
- Automatic Difficulty Retargeting
- Longest Valid Chain (Fork Resolution)

---

# Project Structure

```
cmd/
    blockchain/

internal/
    blockchain/
    crypto/
    ledger/
    merkle/
    miner/
    storage/
    transaction/
    wallet/

data/
    chain.json

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

# Run Tests

```bash
go test ./...
```

---

# Initialize Blockchain

```bash
go run ./cmd/blockchain init
```

---

# Create Wallet

```bash
go run ./cmd/blockchain create-wallet Alice

go run ./cmd/blockchain create-wallet Bob
```

---

# Fund an Account

```bash
go run ./cmd/blockchain fund Alice 100
```

---

# Mine Pending Transactions

```bash
go run ./cmd/blockchain mine
```

---

# Send Funds

```bash
go run ./cmd/blockchain send Alice Bob 40
```

---

# Print Blockchain

```bash
go run ./cmd/blockchain print
```

---

# Show Balances

```bash
go run ./cmd/blockchain balances
```

---

# Validate Blockchain

```bash
go run ./cmd/blockchain validate
```

---

# Example Workflow

```text
Initialize blockchain

↓

Create wallets

↓

Fund Alice

↓

Mine

↓

Alice sends funds to Bob

↓

Mine

↓

Print blockchain

↓

Validate blockchain

↓

View balances
```

---

# Example Output

```text
Initializing blockchain...
Blockchain initialized.

Funding transaction added to pending pool.

Mining pending transactions...

Block mined successfully.

Height : 1

Hash : 00b0...

Nonce : 537
```

---

# Design Decisions

- SHA-256 is used for deterministic block hashing.
- Transactions are summarised using a Merkle Root.
- Proof-of-Work uses configurable mining difficulty.
- Blockchain validation verifies:
  - Previous hash links
  - Block hashes
  - Proof-of-Work
  - Merkle Root
  - Block order
  - Timestamp consistency
  - Transaction signatures
- Persistent storage uses JSON serialization.
- Concurrent mining distributes nonce searching across multiple goroutines.

---

# Optional Features Implemented

- Digital Signatures
- Merkle Root
- Concurrent Mining
- Difficulty Retargeting
- Fork Resolution

---

# Known Limitations

This project is intentionally a single-node educational blockchain.

The following production blockchain features are outside the project scope:

- Peer-to-peer networking
- Distributed consensus
- Transaction mempool synchronization
- Smart contracts
- Block broadcasting
- Real cryptocurrency economics

---

# Future Improvements

Possible future enhancements include:

- REST API
- Web UI
- P2P Networking
- Proof of Stake
- Multiple Mining Nodes
- Smart Contracts
- Automatic Peer Discovery

---

# Author

Kumal Jayawardena

Backend Engineering Internship Assessment

Built using Go.