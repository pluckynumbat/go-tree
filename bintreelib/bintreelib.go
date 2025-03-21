// Package bintreelib: Basic Binary Tree Stuff
package bintreelib

import (
	"fmt"

	"github.com/pluckynumbat/go-quez/sgquezlib"
	"github.com/pluckynumbat/go-stax/sgstaxlib"
)

var nodeNilError error = fmt.Errorf("the node is nil")
var treeNilError = fmt.Errorf("the binary tree is nil")
var treeEmptyError = fmt.Errorf("the binary tree is empty")

type Node struct {
	data   string
	parent *Node
	left   *Node
	right  *Node
}

// Node's implementation of the fmt.Stringer interface
func (node *Node) String() string {
	if node == nil {
		return "nil"
	}
	return node.data
}

// Parent is used to get a pointer to the parent node of a given node
func (node *Node) Parent() (*Node, error) {
	if node == nil {
		return nil, nodeNilError
	}
	return node.parent, nil
}

// LeftChild is used to get a pointer to the left child of a given node
func (node *Node) LeftChild() (*Node, error) {
	if node == nil {
		return nil, nodeNilError
	}
	return node.left, nil
}

// RightChild is used to get a pointer to the right child of a given node
func (node *Node) RightChild() (*Node, error) {
	if node == nil {
		return nil, nodeNilError
	}
	return node.right, nil
}

type BinaryTree struct {
	root     *Node
	lastLeaf *Node
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

// LastLeaf returns a pointer to the last leaf of the Binary Tree
func (bt *BinaryTree) LastLeaf() *Node {
	if bt.IsNil() {
		return nil
	}
	return bt.lastLeaf
}

// AddNodeBFS finds the next free position using a breadth first search and adds a node there
func (bt *BinaryTree) AddNodeBFS(val string) error {
	if bt == nil {
		return treeNilError
	}

	node := &Node{val, nil, nil, nil}

	if bt.root == nil {
		//insert as root
		bt.root = node
		bt.lastLeaf = bt.root
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
			node.parent = runner
			runner.left = node
			bt.lastLeaf = runner.left
			return nil
		} else {
			err2 = queue.Enqueue(runner.left)
			if err2 != nil {
				return fmt.Errorf("add node (BFS) failed with error: %v", err2)
			}
		}

		if runner.right == nil {
			//insert as right child
			node.parent = runner
			runner.right = node
			bt.lastLeaf = runner.right
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
// In a pre-order manner (visit a node, then its left subtree, followed by its right subtree)
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

// TraverseDFSPreOrderIterative returns a string that represents the traversal order of nodes using Depth First Search
// In a pre-order manner (visit a node, then its left subtree, followed by its right subtree)
// This method simulates recursion using the semi generic stack
func (bt *BinaryTree) TraverseDFSPreOrderIterative() (string, error) {
	if bt.IsNil() {
		return "", treeNilError
	}

	if bt.IsEmpty() {
		return "", treeEmptyError
	}

	treeStr := ""

	stack := &sgstaxlib.SemiGenericStack[*Node]{}
	err := stack.Push(bt.root)
	if err != nil {
		return "", fmt.Errorf("DFS (pre order) iterative traversal failed with error: %v", err)
	}

	for !stack.IsEmpty() {
		runner, err2 := stack.Pop()

		if err2 != nil {
			return "", fmt.Errorf("DFS (pre order) iterative traversal failed with error: %v", err2)
		}

		treeStr += fmt.Sprintf("-%v-", runner)

		// first right, then left so that they are popped in the correct order

		if runner.right != nil {
			err2 = stack.Push(runner.right)
			if err2 != nil {
				return "", fmt.Errorf("DFS (pre order) iterative traversal failed with error: %v", err2)
			}
		}

		if runner.left != nil {
			err2 = stack.Push(runner.left)
			if err2 != nil {
				return "", fmt.Errorf("DFS (pre order) iterative traversal failed with error: %v", err2)
			}
		}
	}

	return treeStr, nil
}

// TraverseDFSInOrderRecursive returns a string that represents the traversal order of nodes using Depth First Search
// In an in-order manner (visit a node's left subtree, then the node itself, followed by its right subtree)
// This method uses recursion
func (bt *BinaryTree) TraverseDFSInOrderRecursive() (string, error) {
	if bt.IsNil() {
		return "", treeNilError
	}

	if bt.IsEmpty() {
		return "", treeEmptyError
	}

	treeStr := dfsInOrderRecurse(bt.root)
	return treeStr, nil
}

func dfsInOrderRecurse(node *Node) string {
	if node == nil {
		return ""
	}

	result := dfsInOrderRecurse(node.left)
	result += fmt.Sprintf("-%v-", node.data)
	result += dfsInOrderRecurse(node.right)

	return result
}

// TraverseDFSInOrderIterative returns a string that represents the traversal order of nodes using Depth First Search
// In an in-order manner (visit a node's left subtree, then the node itself, followed by its right subtree)
// This method simulates recursion using the semi generic stack
func (bt *BinaryTree) TraverseDFSInOrderIterative() (string, error) {
	if bt.IsNil() {
		return "", treeNilError
	}

	if bt.IsEmpty() {
		return "", treeEmptyError
	}

	runner := bt.root
	stack := sgstaxlib.SemiGenericStack[*Node]{}
	treeStr := ""

	for runner != nil || !stack.IsEmpty() {
		if runner != nil {
			err := stack.Push(runner)
			if err != nil {
				return "", fmt.Errorf("DFS (in order) iterative traversal failed with error: %v", err)
			}
			runner = runner.left

		} else {
			runner, err := stack.Pop()
			if err != nil {
				return "", fmt.Errorf("DFS (in order) iterative traversal failed with error: %v", err)
			}
			treeStr += fmt.Sprintf("-%v-", runner.data) // visit step
			runner = runner.right
		}
	}

	return treeStr, nil
}

// TraverseDFSPostOrderRecursive returns a string that represents the traversal order of nodes using Depth First Search
// In a post-order manner (visit a node's left subtree, then the node's left subtree, finally the node itself)
// This method uses recursion
func (bt *BinaryTree) TraverseDFSPostOrderRecursive() (string, error) {
	if bt.IsNil() {
		return "", treeNilError
	}

	if bt.IsEmpty() {
		return "", treeEmptyError
	}

	treeStr := dfsPostOrderRecurse(bt.root)
	return treeStr, nil
}

func dfsPostOrderRecurse(node *Node) string {
	if node == nil {
		return ""
	}

	result := dfsPostOrderRecurse(node.left)
	result += dfsPostOrderRecurse(node.right)
	result += fmt.Sprintf("-%v-", node.data)

	return result
}

// TraverseDFSPostOrderIterative returns a string that represents the traversal order of nodes using Depth First Search
// In a post-order manner (visit a node's left subtree, then the node's left subtree, finally the node itself)
// This method simulates recursion using the semi generic stack
func (bt *BinaryTree) TraverseDFSPostOrderIterative() (string, error) {
	if bt.IsNil() {
		return "", treeNilError
	}

	if bt.IsEmpty() {
		return "", treeEmptyError
	}

	treeStr := ""

	runner := bt.root
	var lastVisited *Node

	stack := sgstaxlib.SemiGenericStack[*Node]{}

	for runner != nil || !stack.IsEmpty() {
		if runner != nil {
			err := stack.Push(runner)
			if err != nil {
				return "", fmt.Errorf("DFS (post order) iterative traversal failed with error: %v", err)
			}
			runner = runner.left
		} else {
			potentialVisit, err := stack.Peek()
			if err != nil {
				return "", fmt.Errorf("DFS (post order) iterative traversal failed with error: %v", err)
			}

			if potentialVisit.right == nil || potentialVisit.right == lastVisited {
				lastVisited, err = stack.Pop()
				if err != nil {
					return "", fmt.Errorf("DFS (post order) iterative traversal failed with error: %v", err)
				}

				treeStr += fmt.Sprintf("-%v-", lastVisited.data)
			} else {
				runner = potentialVisit.right
			}
		}
	}

	return treeStr, nil
}

// Contains performs a BFS on the binary tree and tells you if the binary tree contains a node for the input string
func (bt *BinaryTree) Contains(val string) (bool, error) {
	if bt.IsNil() {
		return false, treeNilError
	}

	if bt.IsEmpty() {
		return false, treeEmptyError
	}

	queue := sgquezlib.SemiGenericQueue[*Node]{}
	err := queue.Enqueue(bt.root)
	if err != nil {
		return false, fmt.Errorf("method Contains() failed with error: %v", err)
	}

	for !queue.IsEmpty() {
		front, err2 := queue.Dequeue()
		if err2 != nil {
			return false, fmt.Errorf("method Contains() failed with error: %v", err2)
		}

		if front.data == val {
			return true, nil
		}

		if front.left != nil {
			err2 = queue.Enqueue(front.left)
			if err2 != nil {
				return false, fmt.Errorf("method Contains() failed with error: %v", err2)
			}
		}

		if front.right != nil {
			err2 = queue.Enqueue(front.right)
			if err2 != nil {
				return false, fmt.Errorf("method Contains() failed with error: %v", err2)
			}
		}
	}
	return false, nil
}

//RemoveValue will remove the first instance of the input value, if it exists in the binary tree
func (bt *BinaryTree) RemoveValue(val string) error {
	if bt.IsNil() {
		return treeNilError
	}

	if bt.IsEmpty() {
		return treeEmptyError
	}

	return fmt.Errorf("placeholder error")
}
