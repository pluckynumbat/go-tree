// Package bstreelib: Basic Binary Search Tree Stuff
package bstreelib

import (
	"cmp"
	"fmt"
	"github.com/pluckynumbat/go-quez/sgquezlib"
	"slices"
)

const invalidCount = -1

var nodeNilError = fmt.Errorf("the node is nil")
var treeNilError = fmt.Errorf("the binary search tree is nil")
var treeEmptyError = fmt.Errorf("the binary search tree is empty")
var noValuesError = fmt.Errorf("there are no values in the input")

// duplicateElementError is a custom error raised when an element already present in the tree is attempted to be inserted
type duplicateElementError[T BinarySearchTreeElement] struct {
	value T
}

// duplicateElementError's implementation of the Error interface
func (err duplicateElementError[T]) Error() string {
	return fmt.Sprintf("the binary search tree already has the value attempting to be inserted: %v", err.value)
}

// BinarySearchTreeElement is a custom interface that combines the constraints of the Ordered and Stringer interfaces
type BinarySearchTreeElement interface {
	cmp.Ordered
	fmt.Stringer
}

// Node is the basic unit of the binary search tree, and contains data which can be anything that implements the BinarySearchTreeElement interface
type Node[T BinarySearchTreeElement] struct {
	data   T
	parent *Node[T]
	left   *Node[T]
	right  *Node[T]
}

// Node's implementation of the fmt.Stringer interface
func (node *Node[T]) String() string {
	if node == nil {
		return "nil"
	}
	return node.data.String()
}

// Parent is used to get a pointer to the parent node of a given node
func (node *Node[T]) Parent() (*Node[T], error) {
	if node == nil {
		return nil, nodeNilError
	}
	return node.parent, nil
}

// LeftChild is used to get a pointer to the left child of a given node
func (node *Node[T]) LeftChild() (*Node[T], error) {
	if node == nil {
		return nil, nodeNilError
	}
	return node.left, nil
}

// RightChild is used to get a pointer to the right child of a given node
func (node *Node[T]) RightChild() (*Node[T], error) {
	if node == nil {
		return nil, nodeNilError
	}
	return node.right, nil
}

// BinarySearchTree struct will hold the core functionality of this library
type BinarySearchTree[T BinarySearchTreeElement] struct {
	root  *Node[T]
	count int
}

// IsNil tells you if the pointer to the binary search tree is nil
func (bst *BinarySearchTree[T]) IsNil() bool {
	return bst == nil
}

// IsEmpty tells you if the binary search tree is empty
func (bst *BinarySearchTree[T]) IsEmpty() bool {
	return bst.IsNil() || bst.root == nil
}

// Root returns a pointer to the root of a binary search tree
func (bst *BinarySearchTree[T]) Root() *Node[T] {
	if bst.IsNil() {
		return nil
	}
	return bst.root
}

// Count returns the number of elements in a binary search tree
func (bst *BinarySearchTree[T]) Count() (int, error) {
	if bst.IsNil() {
		return invalidCount, treeNilError
	}
	return bst.count, nil
}

// Insert will add a new value to the binary search tree at the correct position
func (bst *BinarySearchTree[T]) Insert(value T) error {
	if bst.IsNil() {
		return treeNilError
	}

	node := &Node[T]{value, nil, nil, nil}

	// empty tree
	if bst.root == nil {
		bst.root = node
		bst.count = 1
		return nil
	}

	runner := bst.root

	for runner != nil {
		if runner.data == value { // the value is already present
			return duplicateElementError[T]{value}
		}

		if runner.data > value {
			if runner.left == nil { // insert as left child
				runner.left = node
				node.parent = runner
				bst.count += 1
				return nil
			}
			runner = runner.left // check left subtree
			continue
		}

		if runner.data < value {
			if runner.right == nil { // insert as right child
				runner.right = node
				node.parent = runner
				bst.count += 1
				return nil
			}
			runner = runner.right // check right subtree
			continue
		}
	}
	return nil
}

// ConstructFromValues is a helper function to insert all the given values (in the order that they are provided) into a binary search tree
func ConstructFromValues[T BinarySearchTreeElement](values ...T) (*BinarySearchTree[T], error) {
	bstree := &BinarySearchTree[T]{}

	for _, val := range values {
		err := bstree.Insert(val)
		if err != nil {
			return nil, fmt.Errorf("construct from values failed with error: %v", err)
		}
	}

	return bstree, nil
}

// TraverseBFS returns a string that represents the traversal order of nodes using Breadth First Search
func (bst *BinarySearchTree[T]) TraverseBFS() (string, error) {
	if bst.IsNil() {
		return "", treeNilError
	}

	if bst.IsEmpty() {
		return "", treeEmptyError
	}

	treeStr := ""
	queue := &sgquezlib.SemiGenericQueue[*Node[T]]{}
	err := queue.Enqueue(bst.root)
	if err != nil {
		return "", fmt.Errorf("BFS traversal failed with error: %v", err)
	}

	for !queue.IsEmpty() {
		runner, err2 := queue.Dequeue()

		if err2 != nil {
			return "", fmt.Errorf("BFS traversal failed with error: %v", err2)
		}

		treeStr += fmt.Sprintf("-(%v)-", runner.data)

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

// TraverseDFSInOrder returns a string that represents the traversal order of nodes using Depth First Search
// In an in-order manner (visit a node's left subtree, then the node itself, followed by its right subtree)
// This method uses recursion
func (bst *BinarySearchTree[T]) TraverseDFSInOrder() (string, error) {
	if bst.IsNil() {
		return "", treeNilError
	}

	if bst.IsEmpty() {
		return "", treeEmptyError
	}

	return recurseDFSInOrder(bst.root), nil
}

func recurseDFSInOrder[T BinarySearchTreeElement](node *Node[T]) string {
	if node == nil {
		return ""
	}

	result := recurseDFSInOrder(node.left)
	result += fmt.Sprintf("-(%v)-", node.data)
	result += recurseDFSInOrder(node.right)

	return result
}

func (bst *BinarySearchTree[T]) TraverseDFSPreOrder() (string, error) {
	if bst.IsNil() {
		return "", treeNilError
	}

	if bst.IsEmpty() {
		return "", treeEmptyError
	}

	return recurseDFSPreOrder(bst.root), nil
}

func recurseDFSPreOrder[T BinarySearchTreeElement](node *Node[T]) string {
	if node == nil {
		return ""
	}

	result := fmt.Sprintf("-(%v)-", node.data)
	result += recurseDFSPreOrder(node.left)
	result += recurseDFSPreOrder(node.right)

	return result
}

func (bst *BinarySearchTree[T]) TraverseDFSPostOrder() (string, error) {
	if bst.IsNil() {
		return "", treeNilError
	}

	if bst.IsEmpty() {
		return "", treeEmptyError
	}

	return recurseDFSPostOrder(bst.root), nil
}

func recurseDFSPostOrder[T BinarySearchTreeElement](node *Node[T]) string {
	if node == nil {
		return ""
	}

	result := recurseDFSPostOrder(node.left)
	result += recurseDFSPostOrder(node.right)
	result += fmt.Sprintf("-(%v)-", node.data)

	return result
}

// Search looks for a given value the binary search tree, and tell you whether that value is present in the tree or not
func (bst *BinarySearchTree[T]) Search(val T) (bool, error) {

	if bst.IsNil() {
		return false, treeNilError
	}

	if bst.IsEmpty() {
		return false, treeEmptyError
	}

	runner := bst.root

	for runner != nil {

		if runner.data == val {
			return true, nil
		}

		if runner.data > val {
			runner = runner.left
		} else {
			runner = runner.right
		}
	}

	return false, nil
}

// ConstructOrderedSlice collects all the elements in the binary search tree in an ordered manner, and returns them in a slice
func (bst *BinarySearchTree[T]) ConstructOrderedSlice() ([]T, error) {

	if bst.IsNil() {
		return nil, treeNilError
	}

	cnt, cntErr := bst.Count()
	if cntErr != nil {
		return nil, cntErr
	}

	result := make([]T, 0, cnt)
	recurseCollectInOrder(&result, bst.root)

	return result, nil
}

func recurseCollectInOrder[T BinarySearchTreeElement](slicePtr *[]T, node *Node[T]) {

	if node == nil {
		return
	}

	recurseCollectInOrder(slicePtr, node.left)
	*slicePtr = append(*slicePtr, node.data)
	recurseCollectInOrder(slicePtr, node.right)
}

// BalanceTree will re-arrange the binary tree into its most balanced form
func (bst *BinarySearchTree[T]) BalanceTree() error {

	if bst.IsNil() {
		return treeNilError
	}

	cnt, cntErr := bst.Count()
	if cntErr != nil {
		return cntErr
	}

	// no need to balance if there are less than 2 nodes in the tree
	if cnt < 2 {
		return nil
	}

	sl, sliceErr := bst.ConstructOrderedSlice()
	if sliceErr != nil {
		return sliceErr
	}

	// re-initialize the tree to an empty one
	bst.root = nil
	bst.count = 0

	// recursively insert elements from the slice into the binary search tree
	insertErr := recurseInsertNode(bst, sl, 0, cnt-1)
	if insertErr != nil {
		return insertErr
	}

	return nil
}

// recurseInsertNode is used to insert the element present in the middle of the given range into the binary search tree
// and then recursively do this on the 2 new ranges created on each side of the middle element
func recurseInsertNode[T BinarySearchTreeElement](bst *BinarySearchTree[T], slice []T, min, max int) error {

	if max < min {
		return nil
	}

	// insert the element present in the middle of the given range
	mid := (min + max) / 2

	err := bst.Insert(slice[mid])
	if err != nil {
		return err
	}

	// recursive calls
	err = recurseInsertNode[T](bst, slice, min, mid-1)
	if err != nil {
		return err
	}

	err = recurseInsertNode[T](bst, slice, mid+1, max)
	if err != nil {
		return err
	}

	return nil
}

// ConstructBalancedTree is a helper function to insert all the given values into a binary search tree, in a manner which creates a balanced tree
func ConstructBalancedTree[T BinarySearchTreeElement](values ...T) (*BinarySearchTree[T], error) {

	cnt := len(values)
	if cnt == 0 {
		return nil, noValuesError
	}

	slices.Sort(values)

	bst := &BinarySearchTree[T]{}

	// recursively insert elements from the input values into the binary search tree
	insertErr := recurseInsertNode(bst, values, 0, cnt-1)
	if insertErr != nil {
		return nil, insertErr
	}

	return bst, nil
}
