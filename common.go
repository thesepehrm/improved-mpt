package main

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"log"

	"github.com/mr-tron/base58"
	"golang.org/x/crypto/sha3"
)

const (
	HashLength     = 32
	AddressPrefix  = "G"
	AddressLength  = 36
	CheckSumLength = 4
)

type Hash [HashLength]byte

func (h *Hash) Bytes() []byte {
	return h[:]
}

func (h *Hash) Hex() string {
	return hex.EncodeToString(h[:])
}

func HashFromBytes(b []byte) Hash {
	h := Hash{}
	max := HashLength
	if len(b) < len(h) {
		max = len(b)
	}
	for i := 0; i < max; i++ {
		h[i] = b[i]
	}
	return h
}

func HashFromHex(h string) (Hash, error) {
	if h[:2] == "0x" {
		h = h[2:]
	}
	bytes, err := hex.DecodeString(h)
	return HashFromBytes(bytes), err
}

func (h Hash) String() string {
	return h.Hex()
}

func (h *Hash) GenerateRandom() {
	_, err := rand.Read(h[:])
	if err != nil {
		log.Panic(err)
	}
}

func HashObject(i interface{}) *Hash {
	serialized := fmt.Sprintf("%v", i)
	h := Hash(sha3.Sum256([]byte(serialized)))
	return &h
}

func HashBytes(b []byte) *Hash {
	h := Hash(sha3.Sum256(b))
	return &h
}

type Address [AddressLength]byte

func (a Address) Bytes() []byte {
	return a[:]
}

func (a *Address) SetBytes(in []byte) {
	if len(in) > len(a) {
		in = in[len(in)-AddressLength:]
	}
	copy(a[AddressLength-len(in):], in)
}

func BytesToAddress(b []byte) Address {
	addr := Address{}
	addr.SetBytes(b)
	return addr
}

func (a Address) String() string {
	return AddressPrefix + base58.Encode(a[:])
}

func (a Address) Raw() []byte {
	return a[:AddressLength-CheckSumLength]
}
