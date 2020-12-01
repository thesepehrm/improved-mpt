package main

import (
	"bytes"
	"errors"
)

var ErrKeyNotFound = errors.New("Key not found")

type Trie struct {
	root Node
}

func NewTrie() *Trie {
	trie := Trie{&BranchNode{}}
	return &trie
}

// turns every byte into two 4-bits.
func bytesToNibbles(payload []byte) []byte {
	n := make([]byte, len(payload)*2)
	for i, b := range payload {
		n[i*2] = b / 16
		n[i*2+1] = b % 16
	}
	return n
}

func (t *Trie) Put(key []byte, value []byte) {

	k := bytesToNibbles(key)
	t.root = t.insert(t.root, k, value)
}

func (t *Trie) insert(node Node, key []byte, value []byte) Node {

	switch n := node.(type) {
	case *BranchNode:
		n.Children[key[0]] = t.insert(n.Children[key[0]], key[1:], value)
		return n
	case *LeafNode:
		if bytes.Equal(n.Key, key) {
			return &LeafNode{Key: key, Val: value}
		} else {
			// convert leaf to a branch
			bn := &BranchNode{}

			// insert new key
			bn = t.insert(bn, key, value).(*BranchNode)

			// insert old key
			bn = t.insert(bn, n.Key, n.Val).(*BranchNode)
			return bn
		}
	case nil:
		return &LeafNode{Key: key, Val: value}
	}

	return nil
}

func (t *Trie) Get(key []byte) ([]byte, error) {
	k := bytesToNibbles(key)
	val, err := t.traverse(t.root, k)
	return val.(ValueNode), err
}

func (t *Trie) traverse(node Node, key []byte) (Node, error) {

	if len(key) == 0 {
		return nil, ErrKeyNotFound
	}

	switch n := node.(type) {

	case *BranchNode:
		return t.traverse(n.Children[key[0]], key[1:])
	case *LeafNode:
		if bytes.Equal(key, n.Key) {
			return ValueNode(n.Val), nil
		}
	}

	return nil, ErrKeyNotFound

}
