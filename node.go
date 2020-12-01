package main

type (
	Node interface {
	}

	BranchNode struct {
		Children [16]Node
	}

	LeafNode struct {
		Key []byte
		Val []byte
	}

	ValueNode []byte
)
