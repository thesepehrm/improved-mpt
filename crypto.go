package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"

	"gitlab.com/thesepehrm/galaxy/common"
	"golang.org/x/crypto/sha3"
)

// GenerateKey generates a private/public key pair
func GenerateKey() (*ecdsa.PrivateKey, error) {
	return ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
}

func FromECDSAPub(pub *ecdsa.PublicKey) []byte {
	if pub == nil || pub.X == nil || pub.Y == nil {
		return nil
	}
	return elliptic.Marshal(elliptic.P256(), pub.X, pub.Y)
}

func checkSum(pubKeyHash []byte) []byte {
	checkHash := sha3.Sum256(pubKeyHash)
	return checkHash[:common.CheckSumLength]
}

/* PubKeyToAddress generates the address based on the public key

First, a SHA3-256 hash is generated from the public key, then hash it is hashed again to use its first 4 bytes to provide a checksum.
Then the checksum gets appended to the first hash which is the byte array for the final address

*/
func PubkeyToAddress(p ecdsa.PublicKey) common.Address {
	pubBytes := FromECDSAPub(&p)
	pubKeyHash := sha3.Sum256(pubBytes)
	check := checkSum(pubKeyHash[:])
	hashedAddress := append(pubKeyHash[:], check...)

	return common.BytesToAddress(hashedAddress)
}
