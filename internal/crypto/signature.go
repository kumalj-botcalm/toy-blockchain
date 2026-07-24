package crypto

import (
	"crypto/ecdsa"
	"crypto/rand"
	"crypto/sha256"
	"crypto/x509"
	"encoding/asn1"
	"encoding/hex"
	"fmt"
	"math/big"
)

type ecdsaSignature struct {
	R, S *big.Int
}

// Sign signs arbitrary data using a hex encoded private key.
func Sign(privateKeyHex string, data []byte) (string, error) {

	privateBytes, err := hex.DecodeString(privateKeyHex)
	if err != nil {
		return "", err
	}

	privateKey, err := x509.ParseECPrivateKey(privateBytes)
	if err != nil {
		return "", err
	}

	hash := sha256.Sum256(data)

	r, s, err := ecdsa.Sign(rand.Reader, privateKey, hash[:])
	if err != nil {
		return "", err
	}

	signature, err := asn1.Marshal(ecdsaSignature{
		R: r,
		S: s,
	})

	if err != nil {
		return "", err
	}

	return hex.EncodeToString(signature), nil
}

// Verify verifies a signature using a hex encoded public key.
func Verify(publicKeyHex string, data []byte, signatureHex string) (bool, error) {

	publicBytes, err := hex.DecodeString(publicKeyHex)
	if err != nil {
		return false, err
	}

	pub, err := x509.ParsePKIXPublicKey(publicBytes)
	if err != nil {
		return false, err
	}

	publicKey, ok := pub.(*ecdsa.PublicKey)
	if !ok {
		return false, fmt.Errorf("invalid public key")
	}

	signatureBytes, err := hex.DecodeString(signatureHex)
	if err != nil {
		return false, err
	}

	var sig ecdsaSignature

	_, err = asn1.Unmarshal(signatureBytes, &sig)
	if err != nil {
		return false, err
	}

	hash := sha256.Sum256(data)

	valid := ecdsa.Verify(
		publicKey,
		hash[:],
		sig.R,
		sig.S,
	)

	return valid, nil
}
