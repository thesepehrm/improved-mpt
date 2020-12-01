package main

import (
	"bytes"
	"crypto/rand"
	"log"
	"math/big"
	"testing"
	"time"

	"gitlab.com/thesepehrm/galaxy/crypto"
)

type tCase struct {
	key   []byte
	value []byte
}

func testCase() *tCase {
	pk, err := crypto.GenerateKey()
	if err != nil {
		log.Panic(err)
	}

	addr := crypto.PubkeyToAddress(pk.PublicKey)
	val := randValue()

	return &tCase{addr.Raw(), val}
}

func randValue() []byte {

	randInt, err := rand.Int(rand.Reader, big.NewInt(1000))
	if err != nil {
		log.Panic(err)
	}
	timestamp, err := time.Now().MarshalBinary()
	if err != nil {
		log.Panic(err)
	}
	return append(randInt.Bytes(), timestamp...)
}

var testCases = make(map[int]*tCase)

const testCasesNumber = 1500

func TestUpdate(t *testing.T) {
	myTrie := NewTrie()

	// Test Insertion
	for i := 0; i < testCasesNumber; i++ {
		tc := testCase()
		testCases[i] = tc
		myTrie.Put(tc.key, tc.value)
	}

	t.Log("Inserted All")

	for i := 0; i < testCasesNumber; i++ {
		tc := testCases[i]
		value, err := myTrie.Get(tc.key)
		if err != nil {
			t.Logf("\ttest i = %d\n", i)
			t.Fatal(err)
		}

		if !bytes.Equal(tc.value, value) {
			t.Fatalf("Expected same value")
		}
	}

	t.Log("Insertion Completed")

	// Test Update
	for i := 0; i < testCasesNumber; i++ {
		testCases[i].value = randValue()
		myTrie.Put(testCases[i].key, testCases[i].value)
	}

	for i := 0; i < testCasesNumber; i++ {
		tc := testCases[i]
		t.Log(i)
		value, err := myTrie.Get(tc.key)
		if err != nil {
			t.Logf("\ttest i = %d\n", i)
			t.Fatal(err)
		}

		if !bytes.Equal(tc.value, value) {
			t.Fatalf("Expected same value")
		}
	}

}

func BenchmarkMassInsertion(b *testing.B) {
	b.StopTimer()
	b.ResetTimer()

	trie := NewTrie()
	// Test Insertion
	for n := 0; n < b.N; n++ {
		for i := 0; i < 100; i++ {
			tc := testCase()
			testCases[i] = tc
			b.StartTimer()
			trie.Put(tc.key, tc.value)
			b.StopTimer()
		}
	}

}
