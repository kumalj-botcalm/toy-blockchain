package wallet

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"	
)

const WalletDirectory = "wallets"

// Wallet represents a blockchain wallet.
type Wallet struct {
	Owner      string `json:"owner"`
	PublicKey  string `json:"public_key"`
	PrivateKey string `json:"private_key"`
}

// Generate creates a new wallet.
func Generate(owner string) (*Wallet, error) {

	if owner == "" {
		return nil, fmt.Errorf("owner cannot be empty")
	}

	privateKey, err := ecdsa.GenerateKey(
		elliptic.P256(),
		rand.Reader,
	)

	if err != nil {
		return nil, err
	}

	privateBytes, err := x509.MarshalECPrivateKey(privateKey)
	if err != nil {
		return nil, err
	}

	publicBytes, err := x509.MarshalPKIXPublicKey(&privateKey.PublicKey)
	if err != nil {
		return nil, err
	}

	return &Wallet{
		Owner: owner,
		PublicKey: hex.EncodeToString(publicBytes),
		PrivateKey: hex.EncodeToString(privateBytes),
	}, nil
}

// Save writes the wallet to disk.
func (w *Wallet) Save() error {

	err := os.MkdirAll(WalletDirectory, 0755)
	if err != nil {
		return err
	}

	filename := filepath.Join(
		WalletDirectory,
		w.Owner+".json",
	)

	data, err := json.MarshalIndent(w, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(filename, data, 0644)
}

// Load reads a wallet from disk.
func Load(owner string) (*Wallet, error) {

	filename := filepath.Join(
		WalletDirectory,
		owner+".json",
	)

	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var wallet Wallet

	err = json.Unmarshal(data, &wallet)
	if err != nil {
		return nil, err
	}

	return &wallet, nil
}