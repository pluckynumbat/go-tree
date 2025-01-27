// Package bintreelib: Basic Binary Tree Stuff
package bintreelib

import (
	"fmt"

	"github.com/pluckynumbat/go-quez/sgquezlib"
)

var treeNilError = fmt.Errorf("the binary tree is nil")
var rootNilError = fmt.Errorf("the root is nil")

type Node struct {
	left  *Node
	data  string
	right *Node
}

// Node's implementation of the fmt.Stringer interface
func (node *Node) String() string {
	if node == nil {
		return "nil"
	}
	return node.data
}

type BinaryTree struct {
	root *Node
}

// AddNodeBFS finds the next free position using a breadth first search and adds a node there
func (bt *BinaryTree) AddNodeBFS(val string) error {
	if bt == nil {
		return treeNilError
	}

	node := &Node{nil, val, nil}

	if bt.root == nil {
		//insert as root
		bt.root = node
		return nil
	}

	queue := &sgquezlib.SemiGenericQueue[*Node]{}
	err := queue.Enqueue(bt.root)
	if err != nil {
		return fmt.Errorf("add node (BFS) failed with error: %v", err)
	}

	for !queue.IsEmpty() {
		runner, err2 := queue.Dequeue()
		if err2 != nil {
			return fmt.Errorf("add node (BFS) failed with error: %v", err2)
		}

		if runner.left == nil {
			//insert as left child
			runner.left = node
			return nil
		} else {
			err2 = queue.Enqueue(runner.left)
			if err2 != nil {
				return fmt.Errorf("add node (BFS) failed with error: %v", err2)
			}
		}

		if runner.right == nil {
			//insert as right child
			runner.right = node
			return nil
		} else {
			err2 = queue.Enqueue(runner.right)
			if err2 != nil {
				return fmt.Errorf("add node (BFS) failed with error: %v", err2)
			}
		}
	}
	return nil
}

// ConstructFromValues is a helper function to add all values from the given slice to a tree
func ConstructFromValues(values ...string) (*BinaryTree, error) {
	binTree := &BinaryTree{}

	for _, val := range values {
		err := binTree.AddNodeBFS(val)
		if err != nil {
			return nil, fmt.Errorf("construct from values failed with error: %v", err)
		}
	}
	return binTree, nil
}

