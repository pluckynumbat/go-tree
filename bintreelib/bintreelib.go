// Package bintreelib: Basic Binary Tree Stuff
package bintreelib

import (
	"fmt"

	"github.com/pluckynumbat/go-quez/sgquezlib"
	"github.com/pluckynumbat/go-stax/sgstaxlib"
)

var treeNilError = fmt.Errorf("the binary tree is nil")
var treeEmptyError = fmt.Errorf("the binary tree is empty")

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

// IsNil tells you if this pointer to the Binary Tree is nil
func (bt *BinaryTree) IsNil() bool {
	return bt == nil
}

// IsEmpty checks whether a Binary Tree is empty
func (bt *BinaryTree) IsEmpty() bool {
	return bt.IsNil() || bt.root == nil
}

// Root returns a pointer to the root of the Binary Tree
func (bt *BinaryTree) Root() *Node {
	if bt.IsNil() {
		return nil
	}
	return bt.root
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

// TraverseBFS returns a string that represents the traversal order of nodes using Breadth First Search
func (bt *BinaryTree) TraverseBFS() (string, error) {
	if bt.IsNil() {
		return "", treeNilError
	}

	if bt.IsEmpty() {
		return "", treeEmptyError
	}

	treeStr := ""

	queue := &sgquezlib.SemiGenericQueue[*Node]{}
	err := queue.Enqueue(bt.root)
	if err != nil {
		return "", fmt.Errorf("BFS traversal failed with error: %v", err)
	}

	for !queue.IsEmpty() {
		runner, err2 := queue.Dequeue()
		if err2 != nil {
			return "", fmt.Errorf("BFS traversal failed with error: %v", err2)
		}

		treeStr += fmt.Sprintf("-%v-", runner)

		if runner.left != nil {
			err2 = queue.Enqueue(runner.left)
			if err2 != nil {
				return "", fmt.Errorf("BFS traversal failed with error: %v", err2)
			}
		}

		if runner.right != nil {
			err2 = queue.Enqueue(runner.right)
			if err2 != nil {
				return "", fmt.Errorf("BFS traversal failed with error: %v", err2)
			}
		}
	}

	return treeStr, nil
}

// TraverseDFSPreOrderRecursive returns a string that represents the traversal order of nodes using Depth First Search
// In a pre-order manner (visit a node, then its left child, followed by its right child)
// This method uses recursion
func (bt *BinaryTree) TraverseDFSPreOrderRecursive() (string, error) {
	if bt.IsNil() {
		return "", treeNilError
	}

	if bt.IsEmpty() {
		return "", treeEmptyError
	}

	treeStr := dfsPreOrderRecurse(bt.root)
	return treeStr, nil
}

func dfsPreOrderRecurse(node *Node) string {
	if node == nil {
		return ""
	}

	result := fmt.Sprintf("-%v-", node)
	result += dfsPreOrderRecurse(node.left)
	result += dfsPreOrderRecurse(node.right)
	return result
}
