package bstreelib

import (
	"errors"
	"fmt"
	"github.com/pluckynumbat/go-quez/sgquezlib"
	"testing"
)

type prInt int // printable int
func (p prInt) String() string {
	return fmt.Sprintf("%v", int(p))
}

type prString string // printable string
func (p prString) String() string {
	return fmt.Sprintf("%v", string(p))
}

type prFloat float32

func (p prFloat) String() string {
	return fmt.Sprintf("%v", float32(p))
}

func TestNodeString(t *testing.T) {
	t.Run("test node string: prInt", func(t *testing.T) {
		node := &Node[prInt]{1, nil, nil, nil}

		want := "1"
		got := node.String()

		if got != want {
			t.Errorf("Node's string returned incorrect results, want: %v, got %v", want, got)
		}
	})

	t.Run("test node string: prString", func(t *testing.T) {
		node := &Node[prString]{"a", nil, nil, nil}

		want := "a"
		got := node.String()

		if got != want {
			t.Errorf("Node's string returned incorrect results, want: %v, got %v", want, got)
		}
	})

	t.Run("test node string: prFloat", func(t *testing.T) {
		node := &Node[prFloat]{3.14, nil, nil, nil}

		want := "3.14"
		got := node.String()

		if got != want {
			t.Errorf("Node's string returned incorrect results, want: %v, got %v", want, got)
		}
	})
}

func TestNodeParent(t *testing.T) {
	t.Run("test node parent: prInt", func(t *testing.T) {
		var n1, n2, n3 *Node[prInt]
		n2 = &Node[prInt]{1, nil, nil, nil}
		n3 = &Node[prInt]{2, n2, nil, nil}

		tests := []struct {
			name      string
			node      *Node[prInt]
			expError  error
			parent    *Node[prInt]
			parentStr string
		}{
			{"nil node", n1, nodeNilError, nil, "nil"},
			{"non nil node, nil parent", n2, nil, nil, "nil"},
			{"non nil node, non nil parent", n3, nil, n2, "1"},
		}

		for _, test := range tests {
			t.Run(test.name, func(t *testing.T) {
				parent, err := test.node.Parent()
				if err != nil {
					if !errors.Is(err, test.expError) {
						t.Fatalf("Parent() failed with unexpected error: %v", err)
					}
				} else if parent != test.parent {
					t.Fatalf("Parent() returned incorrect parent pointer, want: %v, got: %v", test.parent, parent)
				} else if parent.String() != test.parentStr {
					t.Errorf("Parent() returned parent pointer with incorrect string, want: %v, got: %v", test.parentStr, parent.String())
				}
			})
		}

		t.Run("test node parent on all nodes of a binary search tree", func(t *testing.T) {
			bst, err := ConstructFromValues[prInt](1, 3, 5, 7, 2, 4, 6)
			if err != nil {
				t.Fatalf("ConstructFromValues() failed with error: %v", err)
			}

			// gather expected parent prInt pointers
			var pr1, pr3, pr5, pr7 prInt = 1, 3, 5, 7
			expParents := []*prInt{nil, &pr1, &pr3, &pr3, &pr5, &pr5, &pr7}

			// construct an expected parent queue from the above pointers
			qParents := sgquezlib.SemiGenericQueue[*prInt]{}
			for _, p := range expParents {
				err = qParents.Enqueue(p)
				if err != nil {
					t.Fatalf("Enqueue() failed with error: %v", err)
				}
			}

			// set up queue for a breadth first search traversal of the binary search tree
			queue := sgquezlib.SemiGenericQueue[*Node[prInt]]{}
			err = queue.Enqueue(bst.root)
			if err != nil {
				t.Fatalf("Enqueue() failed with error: %v", err)
			}

			// do a BFS traversal of the tree, checking each node's parents against the expected parent
			for !queue.IsEmpty() {

				runner, err2 := queue.Dequeue()
				if err2 != nil {
					t.Fatalf("Dequeue() failed with error: %v", err2)
				}

				actualParent, err2 := runner.Parent()
				if err2 != nil {
					t.Fatalf("Parent() failed with error: %v", err2)
				}
				got := "nil"
				if actualParent != nil {
					got = actualParent.String()
				}

				expectedParent, err2 := qParents.Dequeue()
				if err2 != nil {
					t.Fatalf("Dequeue() failed with error: %v", err2)
				}
				want := "nil"
				if expectedParent != nil {
					want = expectedParent.String()
				}

				if got != want {
					t.Errorf("Parent() returned incorrect results, want: %v, got: %v", want, got)
				}

				if runner.left != nil {
					err2 = queue.Enqueue(runner.left)
					if err2 != nil {
						t.Fatalf("Enqueue() failed with error: %v", err2)
					}
				}

				if runner.right != nil {
					err2 = queue.Enqueue(runner.right)
					if err2 != nil {
						t.Fatalf("Enqueue() failed with error: %v", err2)
					}
				}
			}
		})
	})

	t.Run("test node parent: prString", func(t *testing.T) {
		var n1, n2, n3 *Node[prString]
		n2 = &Node[prString]{"a", nil, nil, nil}
		n3 = &Node[prString]{"b", n2, nil, nil}

		tests := []struct {
			name      string
			node      *Node[prString]
			expError  error
			parent    *Node[prString]
			parentStr string
		}{
			{"nil node", n1, nodeNilError, nil, "nil"},
			{"non nil node, nil parent", n2, nil, nil, "nil"},
			{"non nil node, non nil parent", n3, nil, n2, "a"},
		}

		for _, test := range tests {
			t.Run(test.name, func(t *testing.T) {
				parent, err := test.node.Parent()
				if err != nil {
					if !errors.Is(err, test.expError) {
						t.Fatalf("Parent() failed with unexpected error: %v", err)
					}
				} else if parent != test.parent {
					t.Fatalf("Parent() returned incorrect parent pointer, want: %v, got: %v", test.parent, parent)
				} else if parent.String() != test.parentStr {
					t.Errorf("Parent() returned parent pointer with incorrect string, want: %v, got: %v", test.parentStr, parent.String())
				}
			})
		}

		t.Run("test node parent on all nodes of a binary search tree", func(t *testing.T) {
			bst, err := ConstructFromValues[prString]("b", "d", "f", "a", "c", "e", "g")
			if err != nil {
				t.Fatalf("ConstructFromValues() failed with error: %v", err)
			}

			// gather expected parent prString pointers
			var prB, prD, prF prString = "b", "d", "f"
			expParents := []*prString{nil, &prB, &prB, &prD, &prD, &prF, &prF}

			// construct an expected parent queue from the above pointers
			qParents := sgquezlib.SemiGenericQueue[*prString]{}
			for _, p := range expParents {
				err = qParents.Enqueue(p)
				if err != nil {
					t.Fatalf("Enqueue() failed with error: %v", err)
				}
			}

			// set up queue for a breadth first search traversal of the binary search tree
			queue := sgquezlib.SemiGenericQueue[*Node[prString]]{}
			err = queue.Enqueue(bst.root)
			if err != nil {
				t.Fatalf("Enqueue() failed with error: %v", err)
			}

			// do a BFS traversal of the tree, checking each node's parents against the expected parent
			for !queue.IsEmpty() {

				runner, err2 := queue.Dequeue()
				if err2 != nil {
					t.Fatalf("Dequeue() failed with error: %v", err2)
				}

				actualParent, err2 := runner.Parent()
				if err2 != nil {
					t.Fatalf("Parent() failed with error: %v", err2)
				}
				got := "nil"
				if actualParent != nil {
					got = actualParent.String()
				}

				expectedParent, err2 := qParents.Dequeue()
				if err2 != nil {
					t.Fatalf("Dequeue() failed with error: %v", err2)
				}
				want := "nil"
				if expectedParent != nil {
					want = expectedParent.String()
				}

				if got != want {
					t.Errorf("Parent() returned incorrect results, want: %v, got: %v", want, got)
				}

				if runner.left != nil {
					err2 = queue.Enqueue(runner.left)
					if err2 != nil {
						t.Fatalf("Enqueue() failed with error: %v", err2)
					}
				}

				if runner.right != nil {
					err2 = queue.Enqueue(runner.right)
					if err2 != nil {
						t.Fatalf("Enqueue() failed with error: %v", err2)
					}
				}
			}
		})
	})

	t.Run("test node parent: prFloat", func(t *testing.T) {
		var n1, n2, n3 *Node[prFloat]
		n2 = &Node[prFloat]{1.9, nil, nil, nil}
		n3 = &Node[prFloat]{2.1, n2, nil, nil}

		tests := []struct {
			name      string
			node      *Node[prFloat]
			expError  error
			parent    *Node[prFloat]
			parentStr string
		}{
			{"nil node", n1, nodeNilError, nil, "nil"},
			{"non nil node, nil parent", n2, nil, nil, "nil"},
			{"non nil node, non nil parent", n3, nil, n2, "1.9"},
		}

		for _, test := range tests {
			t.Run(test.name, func(t *testing.T) {
				parent, err := test.node.Parent()
				if err != nil {
					if !errors.Is(err, test.expError) {
						t.Fatalf("Parent() failed with unexpected error: %v", err)
					}
				} else if parent != test.parent {
					t.Fatalf("Parent() returned incorrect parent pointer, want: %v, got: %v", test.parent, parent)
				} else if parent.String() != test.parentStr {
					t.Errorf("Parent() returned parent pointer with incorrect string, want: %v, got: %v", test.parentStr, parent.String())
				}
			})
		}

		t.Run("test node parent on all nodes of a binary search tree", func(t *testing.T) {
			bst, err := ConstructFromValues[prFloat](0.1, 0.3, 0.5, 0.7, 0.2, 0.4, 0.6)
			if err != nil {
				t.Fatalf("ConstructFromValues() failed with error: %v", err)
			}

			// gather expected parent prFloat pointers
			var pr1, pr3, pr5, pr7 prFloat = 0.1, 0.3, 0.5, 0.7
			expParents := []*prFloat{nil, &pr1, &pr3, &pr3, &pr5, &pr5, &pr7}

			// construct an expected parent queue from the above pointers
			qParents := sgquezlib.SemiGenericQueue[*prFloat]{}
			for _, p := range expParents {
				err = qParents.Enqueue(p)
				if err != nil {
					t.Fatalf("Enqueue() failed with error: %v", err)
				}
			}

			// set up queue for a breadth first search traversal of the binary search tree
			queue := sgquezlib.SemiGenericQueue[*Node[prFloat]]{}
			err = queue.Enqueue(bst.root)
			if err != nil {
				t.Fatalf("Enqueue() failed with error: %v", err)
			}

			// do a BFS traversal of the tree, checking each node's parents against the expected parent
			for !queue.IsEmpty() {

				runner, err2 := queue.Dequeue()
				if err2 != nil {
					t.Fatalf("Dequeue() failed with error: %v", err2)
				}

				actualParent, err2 := runner.Parent()
				if err2 != nil {
					t.Fatalf("Parent() failed with error: %v", err2)
				}
				got := "nil"
				if actualParent != nil {
					got = actualParent.String()
				}

				expectedParent, err2 := qParents.Dequeue()
				if err2 != nil {
					t.Fatalf("Dequeue() failed with error: %v", err2)
				}
				want := "nil"
				if expectedParent != nil {
					want = expectedParent.String()
				}

				if got != want {
					t.Errorf("Parent() returned incorrect results, want: %v, got: %v", want, got)
				}

				if runner.left != nil {
					err2 = queue.Enqueue(runner.left)
					if err2 != nil {
						t.Fatalf("Enqueue() failed with error: %v", err2)
					}
				}

				if runner.right != nil {
					err2 = queue.Enqueue(runner.right)
					if err2 != nil {
						t.Fatalf("Enqueue() failed with error: %v", err2)
					}
				}
			}
		})
	})
}

func TestNodeLeftChild(t *testing.T) {
	t.Run("test node left child: prInt", func(t *testing.T) {
		var n1, n2, n3 *Node[prInt]
		n2 = &Node[prInt]{1, nil, nil, nil}
		n3 = &Node[prInt]{2, n2, nil, nil}
		n2.left = n3

		tests := []struct {
			name         string
			node         *Node[prInt]
			expError     error
			leftChild    *Node[prInt]
			leftChildStr string
		}{
			{"nil node", n1, nodeNilError, nil, "nil"},
			{"non nil node, nil left child", n3, nil, nil, "nil"},
			{"non nil node, non nil left child", n2, nil, n3, "2"},
		}

		for _, test := range tests {
			t.Run(test.name, func(t *testing.T) {
				leftChild, err := test.node.LeftChild()
				if err != nil {
					if !errors.Is(err, test.expError) {
						t.Fatalf("LeftChild() failed with unexpected error: %v", err)
					}
				} else if leftChild != test.leftChild {
					t.Fatalf("LeftChild() returned incorrect left child pointer, want: %v, got: %v", test.leftChild, leftChild)
				} else if leftChild.String() != test.leftChildStr {
					t.Errorf("LeftChild() returned left child pointer with incorrect string, want: %v, got: %v", test.leftChildStr, leftChild.String())
				}
			})
		}

		t.Run("test node left child on all nodes of a binary search tree", func(t *testing.T) {
			bst, err := ConstructFromValues[prInt](1, 3, 5, 7, 2, 4, 6)
			if err != nil {
				t.Fatalf("ConstructFromValues() failed with error: %v", err)
			}

			// gather expected left child prInt pointers
			var pr2, pr4, pr6 prInt = 2, 4, 6
			expLChild := []*prInt{nil, &pr2, nil, &pr4, nil, &pr6, nil}

			// construct an expected left child queue from the above pointers
			qLChild := sgquezlib.SemiGenericQueue[*prInt]{}
			for _, p := range expLChild {
				err = qLChild.Enqueue(p)
				if err != nil {
					t.Fatalf("Enqueue() failed with error: %v", err)
				}
			}

			// set up queue for a breadth first search traversal of the binary search tree
			queue := sgquezlib.SemiGenericQueue[*Node[prInt]]{}
			err = queue.Enqueue(bst.root)
			if err != nil {
				t.Fatalf("Enqueue() failed with error: %v", err)
			}

			// do a BFS traversal of the tree, checking each node's left child against the expected left child
			for !queue.IsEmpty() {

				runner, err2 := queue.Dequeue()
				if err2 != nil {
					t.Fatalf("Dequeue() failed with error: %v", err2)
				}

				actualLeftChild, err2 := runner.LeftChild()
				if err2 != nil {
					t.Fatalf("LeftChild() failed with error: %v", err2)
				}
				got := "nil"
				if actualLeftChild != nil {
					got = actualLeftChild.String()
				}

				expectedLeftChild, err2 := qLChild.Dequeue()
				if err2 != nil {
					t.Fatalf("Dequeue() failed with error: %v", err2)
				}
				want := "nil"
				if expectedLeftChild != nil {
					want = expectedLeftChild.String()
				}

				if got != want {
					t.Errorf("LeftChild() returned incorrect results, want: %v, got: %v", want, got)
				}

				if runner.left != nil {
					err2 = queue.Enqueue(runner.left)
					if err2 != nil {
						t.Fatalf("Enqueue() failed with error: %v", err2)
					}
				}

				if runner.right != nil {
					err2 = queue.Enqueue(runner.right)
					if err2 != nil {
						t.Fatalf("Enqueue() failed with error: %v", err2)
					}
				}
			}
		})
	})

	t.Run("test node left child: prString", func(t *testing.T) {
		var n1, n2, n3 *Node[prString]
		n2 = &Node[prString]{"a", nil, nil, nil}
		n3 = &Node[prString]{"b", n2, nil, nil}
		n2.left = n3

		tests := []struct {
			name         string
			node         *Node[prString]
			expError     error
			leftChild    *Node[prString]
			leftChildStr string
		}{
			{"nil node", n1, nodeNilError, nil, "nil"},
			{"non nil node, nil left child", n3, nil, nil, "nil"},
			{"non nil node, non nil left child", n2, nil, n3, "b"},
		}

		for _, test := range tests {
			t.Run(test.name, func(t *testing.T) {
				leftChild, err := test.node.LeftChild()
				if err != nil {
					if !errors.Is(err, test.expError) {
						t.Fatalf("LeftChild() failed with unexpected error: %v", err)
					}
				} else if leftChild != test.leftChild {
					t.Fatalf("LeftChild() returned incorrect left child pointer, want: %v, got: %v", test.leftChild, leftChild)
				} else if leftChild.String() != test.leftChildStr {
					t.Errorf("LeftChild() returned left child pointer with incorrect string, want: %v, got: %v", test.leftChildStr, leftChild.String())
				}
			})
		}

		t.Run("test node left child on all nodes of a binary search tree", func(t *testing.T) {
			bst, err := ConstructFromValues[prString]("b", "d", "f", "a", "c", "e", "g")
			if err != nil {
				t.Fatalf("ConstructFromValues() failed with error: %v", err)
			}

			// gather expected left child prString pointers
			var prA, prC, prE prString = "a", "c", "e"
			expLChild := []*prString{&prA, nil, &prC, nil, &prE, nil, nil}

			// construct an expected left child queue from the above pointers
			qLChild := sgquezlib.SemiGenericQueue[*prString]{}
			for _, p := range expLChild {
				err = qLChild.Enqueue(p)
				if err != nil {
					t.Fatalf("Enqueue() failed with error: %v", err)
				}
			}

			// set up queue for a breadth first search traversal of the binary search tree
			queue := sgquezlib.SemiGenericQueue[*Node[prString]]{}
			err = queue.Enqueue(bst.root)
			if err != nil {
				t.Fatalf("Enqueue() failed with error: %v", err)
			}

			// do a BFS traversal of the tree, checking each node's left child against the expected left child
			for !queue.IsEmpty() {

				runner, err2 := queue.Dequeue()
				if err2 != nil {
					t.Fatalf("Dequeue() failed with error: %v", err2)
				}

				actualLeftChild, err2 := runner.LeftChild()
				if err2 != nil {
					t.Fatalf("LeftChild() failed with error: %v", err2)
				}
				got := "nil"
				if actualLeftChild != nil {
					got = actualLeftChild.String()
				}

				expectedLeftChild, err2 := qLChild.Dequeue()
				if err2 != nil {
					t.Fatalf("Dequeue() failed with error: %v", err2)
				}
				want := "nil"
				if expectedLeftChild != nil {
					want = expectedLeftChild.String()
				}

				if got != want {
					t.Errorf("LeftChild() returned incorrect results, want: %v, got: %v", want, got)
				}

				if runner.left != nil {
					err2 = queue.Enqueue(runner.left)
					if err2 != nil {
						t.Fatalf("Enqueue() failed with error: %v", err2)
					}
				}

				if runner.right != nil {
					err2 = queue.Enqueue(runner.right)
					if err2 != nil {
						t.Fatalf("Enqueue() failed with error: %v", err2)
					}
				}
			}
		})
	})

	t.Run("test node left child: prFloat", func(t *testing.T) {
		var n1, n2, n3 *Node[prFloat]
		n2 = &Node[prFloat]{1.2, nil, nil, nil}
		n3 = &Node[prFloat]{1.1, n2, nil, nil}
		n2.left = n3

		tests := []struct {
			name         string
			node         *Node[prFloat]
			expError     error
			leftChild    *Node[prFloat]
			leftChildStr string
		}{
			{"nil node", n1, nodeNilError, nil, "nil"},
			{"non nil node, nil left child", n3, nil, nil, "nil"},
			{"non nil node, non nil left child", n2, nil, n3, "1.1"},
		}

		for _, test := range tests {
			t.Run(test.name, func(t *testing.T) {
				leftChild, err := test.node.LeftChild()
				if err != nil {
					if !errors.Is(err, test.expError) {
						t.Fatalf("LeftChild() failed with unexpected error: %v", err)
					}
				} else if leftChild != test.leftChild {
					t.Fatalf("LeftChild() returned incorrect left child pointer, want: %v, got: %v", test.leftChild, leftChild)
				} else if leftChild.String() != test.leftChildStr {
					t.Errorf("LeftChild() returned left child pointer with incorrect string, want: %v, got: %v", test.leftChildStr, leftChild.String())
				}
			})
		}

		t.Run("test node left child on all nodes of a binary search tree", func(t *testing.T) {
			bst, err := ConstructFromValues[prFloat](0.1, 0.3, 0.5, 0.7, 0.2, 0.4, 0.6)
			if err != nil {
				t.Fatalf("ConstructFromValues() failed with error: %v", err)
			}

			// gather expected left child prFloat pointers
			var pr2, pr4, pr6 prFloat = 0.2, 0.4, 0.6
			expLChild := []*prFloat{nil, &pr2, nil, &pr4, nil, &pr6, nil}

			// construct an expected left child queue from the above pointers
			qLChild := sgquezlib.SemiGenericQueue[*prFloat]{}
			for _, p := range expLChild {
				err = qLChild.Enqueue(p)
				if err != nil {
					t.Fatalf("Enqueue() failed with error: %v", err)
				}
			}

			// set up queue for a breadth first search traversal of the binary search tree
			queue := sgquezlib.SemiGenericQueue[*Node[prFloat]]{}
			err = queue.Enqueue(bst.root)
			if err != nil {
				t.Fatalf("Enqueue() failed with error: %v", err)
			}

			// do a BFS traversal of the tree, checking each node's left child against the expected left child
			for !queue.IsEmpty() {

				runner, err2 := queue.Dequeue()
				if err2 != nil {
					t.Fatalf("Dequeue() failed with error: %v", err2)
				}

				actualLeftChild, err2 := runner.LeftChild()
				if err2 != nil {
					t.Fatalf("LeftChild() failed with error: %v", err2)
				}
				got := "nil"
				if actualLeftChild != nil {
					got = actualLeftChild.String()
				}

				expectedLeftChild, err2 := qLChild.Dequeue()
				if err2 != nil {
					t.Fatalf("Dequeue() failed with error: %v", err2)
				}
				want := "nil"
				if expectedLeftChild != nil {
					want = expectedLeftChild.String()
				}

				if got != want {
					t.Errorf("LeftChild() returned incorrect results, want: %v, got: %v", want, got)
				}

				if runner.left != nil {
					err2 = queue.Enqueue(runner.left)
					if err2 != nil {
						t.Fatalf("Enqueue() failed with error: %v", err2)
					}
				}

				if runner.right != nil {
					err2 = queue.Enqueue(runner.right)
					if err2 != nil {
						t.Fatalf("Enqueue() failed with error: %v", err2)
					}
				}
			}
		})
	})
}

func TestNodeRightChild(t *testing.T) {
	t.Run("test node right child: prInt", func(t *testing.T) {
		var n1, n2, n3 *Node[prInt]
		n2 = &Node[prInt]{1, nil, nil, nil}
		n3 = &Node[prInt]{2, n2, nil, nil}
		n2.right = n3

		tests := []struct {
			name          string
			node          *Node[prInt]
			expError      error
			rightChild    *Node[prInt]
			rightChildStr string
		}{
			{"nil node", n1, nodeNilError, nil, "nil"},
			{"non nil node, nil right child", n3, nil, nil, "nil"},
			{"non nil node, non nil right child", n2, nil, n3, "2"},
		}

		for _, test := range tests {
			t.Run(test.name, func(t *testing.T) {
				rightChild, err := test.node.RightChild()
				if err != nil {
					if !errors.Is(err, test.expError) {
						t.Fatalf("RightChild() failed with unexpected error: %v", err)
					}
				} else if rightChild != test.rightChild {
					t.Fatalf("RightChild() returned incorrect right child pointer, want: %v, got: %v", test.rightChild, rightChild)
				} else if rightChild.String() != test.rightChildStr {
					t.Errorf("RightChild() returned right child pointer with incorrect string, want: %v, got: %v", test.rightChildStr, rightChild.String())
				}
			})
		}

		t.Run("test node right child on all nodes of a binary search tree", func(t *testing.T) {
			bst, err := ConstructFromValues[prInt](1, 3, 5, 7, 2, 4, 6)
			if err != nil {
				t.Fatalf("ConstructFromValues() failed with error: %v", err)
			}

			// gather expected right child prInt pointers
			var pr3, pr5, pr7 prInt = 3, 5, 7
			expRChild := []*prInt{&pr3, &pr5, nil, &pr7, nil, nil, nil}

			// construct an expected right child queue from the above pointers
			qRChild := sgquezlib.SemiGenericQueue[*prInt]{}
			for _, p := range expRChild {
				err = qRChild.Enqueue(p)
				if err != nil {
					t.Fatalf("Enqueue() failed with error: %v", err)
				}
			}

			// set up queue for a breadth first search traversal of the binary search tree
			queue := sgquezlib.SemiGenericQueue[*Node[prInt]]{}
			err = queue.Enqueue(bst.root)
			if err != nil {
				t.Fatalf("Enqueue() failed with error: %v", err)
			}

			// do a BFS traversal of the tree, checking each node's right child against the expected right child
			for !queue.IsEmpty() {

				runner, err2 := queue.Dequeue()
				if err2 != nil {
					t.Fatalf("Dequeue() failed with error: %v", err2)
				}

				actualRightChild, err2 := runner.RightChild()
				if err2 != nil {
					t.Fatalf("RightChild() failed with error: %v", err2)
				}
				got := "nil"
				if actualRightChild != nil {
					got = actualRightChild.String()
				}

				expectedRightChild, err2 := qRChild.Dequeue()
				if err2 != nil {
					t.Fatalf("Dequeue() failed with error: %v", err2)
				}
				want := "nil"
				if expectedRightChild != nil {
					want = expectedRightChild.String()
				}

				if got != want {
					t.Errorf("RightChild() returned incorrect results, want: %v, got: %v", want, got)
				}

				if runner.left != nil {
					err2 = queue.Enqueue(runner.left)
					if err2 != nil {
						t.Fatalf("Enqueue() failed with error: %v", err2)
					}
				}

				if runner.right != nil {
					err2 = queue.Enqueue(runner.right)
					if err2 != nil {
						t.Fatalf("Enqueue() failed with error: %v", err2)
					}
				}
			}
		})
	})

	t.Run("test node right child: prString", func(t *testing.T) {
		var n1, n2, n3 *Node[prString]
		n2 = &Node[prString]{"a", nil, nil, nil}
		n3 = &Node[prString]{"b", n2, nil, nil}
		n2.right = n3

		tests := []struct {
			name          string
			node          *Node[prString]
			expError      error
			rightChild    *Node[prString]
			rightChildStr string
		}{
			{"nil node", n1, nodeNilError, nil, "nil"},
			{"non nil node, nil right child", n3, nil, nil, "nil"},
			{"non nil node, non nil right child", n2, nil, n3, "b"},
		}

		for _, test := range tests {
			t.Run(test.name, func(t *testing.T) {
				rightChild, err := test.node.RightChild()
				if err != nil {
					if !errors.Is(err, test.expError) {
						t.Fatalf("RightChild() failed with unexpected error: %v", err)
					}
				} else if rightChild != test.rightChild {
					t.Fatalf("RightChild() returned incorrect right child pointer, want: %v, got: %v", test.rightChild, rightChild)
				} else if rightChild.String() != test.rightChildStr {
					t.Errorf("RightChild() returned right child pointer with incorrect string, want: %v, got: %v", test.rightChildStr, rightChild.String())
				}
			})
		}

		t.Run("test node right child on all nodes of a binary search tree", func(t *testing.T) {
			bst, err := ConstructFromValues[prString]("b", "d", "f", "a", "c", "e", "g")
			if err != nil {
				t.Fatalf("ConstructFromValues() failed with error: %v", err)
			}

			// gather expected right child prString pointers
			var prD, prF, prG prString = "d", "f", "g"
			expRChild := []*prString{&prD, nil, &prF, nil, &prG, nil, nil}

			// construct an expected right child queue from the above pointers
			qRChild := sgquezlib.SemiGenericQueue[*prString]{}
			for _, p := range expRChild {
				err = qRChild.Enqueue(p)
				if err != nil {
					t.Fatalf("Enqueue() failed with error: %v", err)
				}
			}

			// set up queue for a breadth first search traversal of the binary search tree
			queue := sgquezlib.SemiGenericQueue[*Node[prString]]{}
			err = queue.Enqueue(bst.root)
			if err != nil {
				t.Fatalf("Enqueue() failed with error: %v", err)
			}

			// do a BFS traversal of the tree, checking each node's right child against the expected right child
			for !queue.IsEmpty() {

				runner, err2 := queue.Dequeue()
				if err2 != nil {
					t.Fatalf("Dequeue() failed with error: %v", err2)
				}

				actualRightChild, err2 := runner.RightChild()
				if err2 != nil {
					t.Fatalf("RightChild() failed with error: %v", err2)
				}
				got := "nil"
				if actualRightChild != nil {
					got = actualRightChild.String()
				}

				expectedRightChild, err2 := qRChild.Dequeue()
				if err2 != nil {
					t.Fatalf("Dequeue() failed with error: %v", err2)
				}
				want := "nil"
				if expectedRightChild != nil {
					want = expectedRightChild.String()
				}

				if got != want {
					t.Errorf("RightChild() returned incorrect results, want: %v, got: %v", want, got)
				}

				if runner.left != nil {
					err2 = queue.Enqueue(runner.left)
					if err2 != nil {
						t.Fatalf("Enqueue() failed with error: %v", err2)
					}
				}

				if runner.right != nil {
					err2 = queue.Enqueue(runner.right)
					if err2 != nil {
						t.Fatalf("Enqueue() failed with error: %v", err2)
					}
				}
			}
		})
	})

	t.Run("test node right child: prFloat", func(t *testing.T) {
		var n1, n2, n3 *Node[prFloat]
		n2 = &Node[prFloat]{1.1, nil, nil, nil}
		n3 = &Node[prFloat]{1.2, n2, nil, nil}
		n2.right = n3

		tests := []struct {
			name          string
			node          *Node[prFloat]
			expError      error
			rightChild    *Node[prFloat]
			rightChildStr string
		}{
			{"nil node", n1, nodeNilError, nil, "nil"},
			{"non nil node, nil right child", n3, nil, nil, "nil"},
			{"non nil node, non nil right child", n2, nil, n3, "1.2"},
		}

		for _, test := range tests {
			t.Run(test.name, func(t *testing.T) {
				rightChild, err := test.node.RightChild()
				if err != nil {
					if !errors.Is(err, test.expError) {
						t.Fatalf("RightChild() failed with unexpected error: %v", err)
					}
				} else if rightChild != test.rightChild {
					t.Fatalf("RightChild() returned incorrect right child pointer, want: %v, got: %v", test.rightChild, rightChild)
				} else if rightChild.String() != test.rightChildStr {
					t.Errorf("RightChild() returned right child pointer with incorrect string, want: %v, got: %v", test.rightChildStr, rightChild.String())
				}
			})
		}

		t.Run("test node right child on all nodes of a binary search tree", func(t *testing.T) {
			bst, err := ConstructFromValues[prFloat](0.1, 0.3, 0.5, 0.7, 0.2, 0.4, 0.6)
			if err != nil {
				t.Fatalf("ConstructFromValues() failed with error: %v", err)
			}

			// gather expected right child prFloat pointers
			var pr3, pr5, pr7 prFloat = 0.3, 0.5, 0.7
			expRChild := []*prFloat{&pr3, &pr5, nil, &pr7, nil, nil, nil}

			// construct an expected right child queue from the above pointers
			qRChild := sgquezlib.SemiGenericQueue[*prFloat]{}
			for _, p := range expRChild {
				err = qRChild.Enqueue(p)
				if err != nil {
					t.Fatalf("Enqueue() failed with error: %v", err)
				}
			}

			// set up queue for a breadth first search traversal of the binary search tree
			queue := sgquezlib.SemiGenericQueue[*Node[prFloat]]{}
			err = queue.Enqueue(bst.root)
			if err != nil {
				t.Fatalf("Enqueue() failed with error: %v", err)
			}

			// do a BFS traversal of the tree, checking each node's right child against the expected right child
			for !queue.IsEmpty() {

				runner, err2 := queue.Dequeue()
				if err2 != nil {
					t.Fatalf("Dequeue() failed with error: %v", err2)
				}

				actualRightChild, err2 := runner.RightChild()
				if err2 != nil {
					t.Fatalf("RightChild() failed with error: %v", err2)
				}
				got := "nil"
				if actualRightChild != nil {
					got = actualRightChild.String()
				}

				expectedRightChild, err2 := qRChild.Dequeue()
				if err2 != nil {
					t.Fatalf("Dequeue() failed with error: %v", err2)
				}
				want := "nil"
				if expectedRightChild != nil {
					want = expectedRightChild.String()
				}

				if got != want {
					t.Errorf("RightChild() returned incorrect results, want: %v, got: %v", want, got)
				}

				if runner.left != nil {
					err2 = queue.Enqueue(runner.left)
					if err2 != nil {
						t.Fatalf("Enqueue() failed with error: %v", err2)
					}
				}

				if runner.right != nil {
					err2 = queue.Enqueue(runner.right)
					if err2 != nil {
						t.Fatalf("Enqueue() failed with error: %v", err2)
					}
				}
			}
		})
	})
}

func TestIsNil(t *testing.T) {
	var bst1, bst2 *BinarySearchTree[prInt]
	bst2 = &BinarySearchTree[prInt]{}

	tests := []struct {
		name string
		bst  *BinarySearchTree[prInt]
		want bool
	}{
		{"nil tree", bst1, true},
		{"non nil tree", bst2, false},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := test.bst.IsNil()
			if got != test.want {
				t.Errorf("IsNil() returned incorrect results, want: %v, got :%v", test.want, got)
			}
		})
	}
}

func TestIsEmpty(t *testing.T) {
	var bst1, bst2, bst3 *BinarySearchTree[prInt]
	bst2 = &BinarySearchTree[prInt]{}
	root := &Node[prInt]{}
	bst3 = &BinarySearchTree[prInt]{root, 1}

	tests := []struct {
		name string
		bst  *BinarySearchTree[prInt]
		want bool
	}{
		{"nil tree", bst1, true},
		{"non nil, empty tree", bst2, true},
		{"non nil, non empty tree", bst3, false},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := test.bst.IsEmpty()
			if got != test.want {
				t.Errorf("IsNil() returned incorrect results, want: %v, got :%v", test.want, got)
			}
		})
	}
}

func TestRoot(t *testing.T) {
	var bst1, bst2, bst3, bst4, bst5 *BinarySearchTree[prInt]

	bst2 = &BinarySearchTree[prInt]{}

	r1 := &Node[prInt]{1, nil, nil, nil}
	bst3 = &BinarySearchTree[prInt]{r1, 1}

	r2 := &Node[prInt]{2, nil, nil, nil}
	n2 := &Node[prInt]{1, r2, nil, nil}
	r2.left = n2
	bst4 = &BinarySearchTree[prInt]{r2, 2}

	r3 := &Node[prInt]{0, nil, nil, nil}
	n4 := &Node[prInt]{-1, r3, nil, nil}
	n5 := &Node[prInt]{1, r3, nil, nil}
	r3.left = n4
	r3.right = n5
	bst5 = &BinarySearchTree[prInt]{r3, 3}

	tests := []struct {
		name       string
		bst        *BinarySearchTree[prInt]
		expRootNil bool
		expRootStr string
	}{
		{"nil tree", bst1, true, "nil"},
		{"empty tree", bst2, true, "nil"},
		{"1 element tree", bst3, false, "1"},
		{"2 element tree", bst4, false, "2"},
		{"3 element tree", bst5, false, "0"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			root := test.bst.Root()

			gotRootNil := root == nil
			wantRootNil := test.expRootNil

			if gotRootNil != wantRootNil {
				t.Fatalf("Unexpected Root() nil status, want: %v, got: %v", wantRootNil, gotRootNil)
			}

			got := root.String()
			want := test.expRootStr

			if got != want {
				t.Errorf("Root() returned incorrect results, want: %v, got: %v", want, got)
			}
		})
	}
}

func TestCount(t *testing.T) {
	var bst1 *BinarySearchTree[prInt]

	_, err := bst1.Count()
	if err == nil {
		t.Fatalf("Count() on a nil tree should have failed")
	} else {
		fmt.Println(err)
	}

	bst2 := &BinarySearchTree[prInt]{}

	cnt, err := bst2.Count()
	if err != nil {
		t.Fatalf("Count() failed with unexpected error, %v", err)
	} else {
		if cnt != 0 {
			t.Errorf("Count() returned incorrect results, want: %v, got: %v", 0, cnt)
		}
	}

	t.Run("test insert on Binary Search Tree of prInt nodes", func(t *testing.T) {
		tests := []struct {
			name     string
			input    []prInt
			expError error
			expCount int
		}{
			{"nil tree", nil, treeNilError, 0},
			{"empty tree", []prInt{}, nil, 0},
			{"1 element tree", []prInt{3}, nil, 1},
			{"2 element tree", []prInt{-1, 1}, nil, 2},
			{"3 element tree", []prInt{20, 40, 60}, nil, 3},
			{"4 element tree", []prInt{4, 44, -444, -4444}, nil, 4},
			{"5 element tree", []prInt{5, 4, 3, 2, 1}, nil, 5},
		}

		for _, test := range tests {
			t.Run(test.name, func(t *testing.T) {

				bst, err2 := ConstructFromValues[prInt](test.input...)
				if err2 != nil && !errors.Is(err2, test.expError) {
					t.Fatalf("ConstructFromValues() failed with unexpected error: %v", err2)
				} else if err2 != nil {
					fmt.Println(err2)
				} else {
					cnt2, err3 := bst.Count()
					if err3 != nil {
						t.Fatalf("Count() failed with unexpected error: %v", err2)
					} else {
						want := test.expCount
						got := cnt2
						if got != want {
							t.Errorf("Count() returned incorrect results, want: %v, got: %v", want, got)
						}
					}
				}
			})
		}
	})
}

func TestInsert(t *testing.T) {

	t.Run("test insert on Binary Search Tree of prInt nodes", func(t *testing.T) {
		var bst1, bst2 *BinarySearchTree[prInt]
		bst2 = &BinarySearchTree[prInt]{}

		tests := []struct {
			name               string
			bst                *BinarySearchTree[prInt]
			val                prInt
			expError           error
			expBFStr           string
			expDFSInorderStr   string
			expDFSPreOrderStr  string
			expDFSPostOrderStr string
		}{
			{"nil tree", bst1, 1, treeNilError, "", "", "", ""},
			{"empty tree", bst2, 1, nil, "-(1)-", "-(1)-", "-(1)-", "-(1)-"},
			{"1 element tree", bst2, 4, nil, "-(1)--(4)-", "-(1)--(4)-", "-(1)--(4)-", "-(4)--(1)-"},
			{"2 element tree", bst2, 6, nil, "-(1)--(4)--(6)-", "-(1)--(4)--(6)-", "-(1)--(4)--(6)-", "-(6)--(4)--(1)-"},
			{"3 element tree", bst2, 2, nil, "-(1)--(4)--(2)--(6)-", "-(1)--(2)--(4)--(6)-", "-(1)--(4)--(2)--(6)-", "-(2)--(6)--(4)--(1)-"},
			{"4 element tree", bst2, 5, nil, "-(1)--(4)--(2)--(6)--(5)-", "-(1)--(2)--(4)--(5)--(6)-", "-(1)--(4)--(2)--(6)--(5)-", "-(2)--(5)--(6)--(4)--(1)-"},
			{"5 element tree", bst2, 3, nil, "-(1)--(4)--(2)--(6)--(3)--(5)-", "-(1)--(2)--(3)--(4)--(5)--(6)-", "-(1)--(4)--(2)--(3)--(6)--(5)-", "-(3)--(2)--(5)--(6)--(4)--(1)-"},
		}

		for _, test := range tests {
			t.Run(test.name, func(t *testing.T) {
				err := test.bst.Insert(test.val)
				if err != nil && !errors.Is(err, test.expError) {
					t.Fatalf("Insert() failed with unexpected error: %v", err)
				} else if err != nil {
					fmt.Println(err)
				} else {
					gotBFSstr, err2 := test.bst.TraverseBFS()
					if err2 != nil {
						t.Fatalf("TraverseBFS() failed with unexpected error: %v", err2)
					} else if gotBFSstr != test.expBFStr {
						t.Errorf("Insert() gave incorrect results, want: %v, got: %v", test.expBFStr, gotBFSstr)
					}

					inorderDFSstr, err2 := test.bst.TraverseDFSInOrder()
					if err2 != nil {
						t.Fatalf("TraverseDFSInOrder() failed with unexpected error: %v", err2)
					} else if inorderDFSstr != test.expDFSInorderStr {
						t.Errorf("Insert() gave incorrect results, want: %v, got: %v", test.expDFSInorderStr, inorderDFSstr)
					}

					preOrderDFSstr, err2 := test.bst.TraverseDFSPreOrder()
					if err2 != nil {
						t.Fatalf("TraverseDFSPreOrder() failed with unexpected error: %v", err2)
					} else if preOrderDFSstr != test.expDFSPreOrderStr {
						t.Errorf("Insert() gave incorrect results, want: %v, got: %v", test.expDFSPreOrderStr, preOrderDFSstr)
					}

					postOrderDFSstr, err2 := test.bst.TraverseDFSPostOrder()
					if err2 != nil {
						t.Fatalf("TraverseDFSPostOrder() failed with unexpected error: %v", err2)
					} else if postOrderDFSstr != test.expDFSPostOrderStr {
						t.Errorf("Insert() gave incorrect results, want: %v, got: %v", test.expDFSPostOrderStr, postOrderDFSstr)
					}
				}
			})
		}
	})

	t.Run("test inserting a value already present in a Binary Search Tree of prInt nodes", func(t *testing.T) {
		bst := &BinarySearchTree[prInt]{}

		err := bst.Insert(1)
		if err != nil {
			t.Fatalf("Insert() failed with error: %v", err)
		}

		err = bst.Insert(1)
		if err == nil {
			t.Fatalf("Insert() using a value already present in the tree should have returned an error")
		} else {
			fmt.Println(err)
		}
	})

	t.Run("test insert on Binary Search Tree of prString nodes", func(t *testing.T) {
		var bst1, bst2 *BinarySearchTree[prString]
		bst2 = &BinarySearchTree[prString]{}

		tests := []struct {
			name               string
			bst                *BinarySearchTree[prString]
			val                prString
			expError           error
			expBFStr           string
			expDFSInorderStr   string
			expDFSPreOrderStr  string
			expDFSPostOrderStr string
		}{
			{"nil tree", bst1, "b", treeNilError, "", "", "", ""},
			{"empty tree", bst2, "b", nil, "-(b)-", "-(b)-", "-(b)-", "-(b)-"},
			{"1 element tree", bst2, "d", nil, "-(b)--(d)-", "-(b)--(d)-", "-(b)--(d)-", "-(d)--(b)-"},
			{"2 element tree", bst2, "f", nil, "-(b)--(d)--(f)-", "-(b)--(d)--(f)-", "-(b)--(d)--(f)-", "-(f)--(d)--(b)-"},
			{"3 element tree", bst2, "a", nil, "-(b)--(a)--(d)--(f)-", "-(a)--(b)--(d)--(f)-", "-(b)--(a)--(d)--(f)-", "-(a)--(f)--(d)--(b)-"},
			{"4 element tree", bst2, "c", nil, "-(b)--(a)--(d)--(c)--(f)-", "-(a)--(b)--(c)--(d)--(f)-", "-(b)--(a)--(d)--(c)--(f)-", "-(a)--(c)--(f)--(d)--(b)-"},
			{"5 element tree", bst2, "e", nil, "-(b)--(a)--(d)--(c)--(f)--(e)-", "-(a)--(b)--(c)--(d)--(e)--(f)-", "-(b)--(a)--(d)--(c)--(f)--(e)-", "-(a)--(c)--(e)--(f)--(d)--(b)-"},
			{"6 element tree", bst2, "g", nil, "-(b)--(a)--(d)--(c)--(f)--(e)--(g)-", "-(a)--(b)--(c)--(d)--(e)--(f)--(g)-", "-(b)--(a)--(d)--(c)--(f)--(e)--(g)-", "-(a)--(c)--(e)--(g)--(f)--(d)--(b)-"},
		}
		for _, test := range tests {
			t.Run(test.name, func(t *testing.T) {
				err := test.bst.Insert(test.val)
				if err != nil && !errors.Is(err, test.expError) {
					t.Fatalf("Insert() failed with unexpected error: %v", err)
				} else if err != nil {
					fmt.Println(err)
				} else {
					gotBFSstr, err2 := test.bst.TraverseBFS()
					if err2 != nil {
						t.Fatalf("TraverseBFS() failed with unexpected error: %v", err2)
					} else if gotBFSstr != test.expBFStr {
						t.Errorf("Insert() gave incorrect results, want: %v, got: %v", test.expBFStr, gotBFSstr)
					}

					inorderDFSstr, err2 := test.bst.TraverseDFSInOrder()
					if err2 != nil {
						t.Fatalf("TraverseDFSInOrder() failed with unexpected error: %v", err2)
					} else if inorderDFSstr != test.expDFSInorderStr {
						t.Errorf("Insert() gave incorrect results, want: %v, got: %v", test.expDFSInorderStr, inorderDFSstr)
					}

					preOrderDFSstr, err2 := test.bst.TraverseDFSPreOrder()
					if err2 != nil {
						t.Fatalf("TraverseDFSPreOrder() failed with unexpected error: %v", err2)
					} else if preOrderDFSstr != test.expDFSPreOrderStr {
						t.Errorf("Insert() gave incorrect results, want: %v, got: %v", test.expDFSPreOrderStr, preOrderDFSstr)
					}

					postOrderDFSstr, err2 := test.bst.TraverseDFSPostOrder()
					if err2 != nil {
						t.Fatalf("TraverseDFSPostOrder() failed with unexpected error: %v", err2)
					} else if postOrderDFSstr != test.expDFSPostOrderStr {
						t.Errorf("Insert() gave incorrect results, want: %v, got: %v", test.expDFSPostOrderStr, postOrderDFSstr)
					}
				}
			})
		}
	})
}

func TestConstructFromValues(t *testing.T) {

	t.Run("test construct from values for type prInt", func(t *testing.T) {

		tests := []struct {
			name                string
			input               []prInt
			shouldConstructFail bool
			expTraverseErr      error
			expBFSStr           string
			expDFSInOrderStr    string
			expDFSPreOrderStr   string
			expDFSPostOrderStr  string
		}{
			{"nil input", nil, false, treeEmptyError, "", "", "", ""},
			{"empty input", []prInt{}, false, treeEmptyError, "", "", "", ""},
			{"2 elements, identical", []prInt{1, 1}, true, nil, "", "", "", ""},
			{"3 elements, -1, 0, 1", []prInt{-1, 0, 1}, false, nil, "-(-1)--(0)--(1)-", "-(-1)--(0)--(1)-", "-(-1)--(0)--(1)-", "-(1)--(0)--(-1)-"},
			{"3 elements, 1, 0, -1", []prInt{1, 0, -1}, false, nil, "-(1)--(0)--(-1)-", "-(-1)--(0)--(1)-", "-(1)--(0)--(-1)-", "-(-1)--(0)--(1)-"},
			{"3 elements, 0, 1, -1", []prInt{0, 1, -1}, false, nil, "-(0)--(-1)--(1)-", "-(-1)--(0)--(1)-", "-(0)--(-1)--(1)-", "-(-1)--(1)--(0)-"},
			{"3 elements, 0, -1, 1", []prInt{0, -1, 1}, false, nil, "-(0)--(-1)--(1)-", "-(-1)--(0)--(1)-", "-(0)--(-1)--(1)-", "-(-1)--(1)--(0)-"},
			{"3 elements, 1, -1, 0", []prInt{1, -1, 0}, false, nil, "-(1)--(-1)--(0)-", "-(-1)--(0)--(1)-", "-(1)--(-1)--(0)-", "-(0)--(-1)--(1)-"},
			{"3 elements, -1, 1, 0", []prInt{-1, 1, 0}, false, nil, "-(-1)--(1)--(0)-", "-(-1)--(0)--(1)-", "-(-1)--(1)--(0)-", "-(0)--(1)--(-1)-"},
			{"7 elements, all positive", []prInt{2, 4, 6, 7, 5, 3, 1}, false, nil, "-(2)--(1)--(4)--(3)--(6)--(5)--(7)-", "-(1)--(2)--(3)--(4)--(5)--(6)--(7)-", "-(2)--(1)--(4)--(3)--(6)--(5)--(7)-", "-(1)--(3)--(5)--(7)--(6)--(4)--(2)-"},
		}

		for _, test := range tests {
			t.Run(test.name, func(t *testing.T) {
				bst, err := ConstructFromValues[prInt](test.input...)

				if err != nil && !test.shouldConstructFail {
					t.Fatalf("ConstructFromValues() failed with unexpected error: %v", err)
				} else if err != nil {
					fmt.Println(err)
				} else {
					gotBFSStr, err2 := bst.TraverseBFS()
					if err2 != nil && !errors.Is(err2, test.expTraverseErr) {
						t.Fatalf("TraverseBFS() failed with unexpected error: %v", err2)
					} else if err2 != nil {
						fmt.Println(err2)
					} else {
						if gotBFSStr != test.expBFSStr {
							t.Fatalf("ConstructFromValues() gave incorrect results, want: %v, got: %v", test.expBFSStr, gotBFSStr)
						}
					}

					gotDFSInOrderStr, err2 := bst.TraverseDFSInOrder()
					if err2 != nil && !errors.Is(err2, test.expTraverseErr) {
						t.Fatalf("TraverseDFSInOrder() failed with unexpected error: %v", err2)
					} else if err2 != nil {
						fmt.Println(err2)
					} else {
						if gotDFSInOrderStr != test.expDFSInOrderStr {
							t.Fatalf("ConstructFromValues() gave incorrect results, want: %v, got: %v", test.expDFSInOrderStr, gotDFSInOrderStr)
						}
					}

					gotDFSPreOrderStr, err2 := bst.TraverseDFSPreOrder()
					if err2 != nil && !errors.Is(err2, test.expTraverseErr) {
						t.Fatalf("TraverseDFSPreOrder() failed with unexpected error: %v", err2)
					} else if err2 != nil {
						fmt.Println(err2)
					} else {
						if gotDFSPreOrderStr != test.expDFSPreOrderStr {
							t.Fatalf("ConstructFromValues() gave incorrect results, want: %v, got: %v", test.expDFSPreOrderStr, gotDFSPreOrderStr)
						}
					}

					gotDFSPostOrderStr, err2 := bst.TraverseDFSPostOrder()
					if err2 != nil && !errors.Is(err2, test.expTraverseErr) {
						t.Fatalf("TraverseDFSPostOrder() failed with unexpected error: %v", err2)
					} else if err2 != nil {
						fmt.Println(err2)
					} else {
						if gotDFSPostOrderStr != test.expDFSPostOrderStr {
							t.Fatalf("ConstructFromValues() gave incorrect results, want: %v, got: %v", test.expDFSPostOrderStr, gotDFSPostOrderStr)
						}
					}
				}
			})
		}
	})

	t.Run("test construct from values for type prString", func(t *testing.T) {

		tests := []struct {
			name                string
			input               []prString
			shouldConstructFail bool
			expTraverseErr      error
			expBFSStr           string
			expDFSInOrderStr    string
			expDFSPreOrderStr   string
			expDFSPostOrderStr  string
		}{
			{"nil input", nil, false, treeEmptyError, "", "", "", ""},
			{"empty input", []prString{}, false, treeEmptyError, "", "", "", ""},
			{"2 elements, identical", []prString{"a", "a"}, true, nil, "", "", "", ""},
			{"3 elements, a, b, c", []prString{"a", "b", "c"}, false, nil, "-(a)--(b)--(c)-", "-(a)--(b)--(c)-", "-(a)--(b)--(c)-", "-(c)--(b)--(a)-"},
			{"3 elements, c, b, a", []prString{"c", "b", "a"}, false, nil, "-(c)--(b)--(a)-", "-(a)--(b)--(c)-", "-(c)--(b)--(a)-", "-(a)--(b)--(c)-"},
			{"3 elements, b, c, a", []prString{"b", "c", "a"}, false, nil, "-(b)--(a)--(c)-", "-(a)--(b)--(c)-", "-(b)--(a)--(c)-", "-(a)--(c)--(b)-"},
			{"3 elements, b, a, c", []prString{"b", "a", "c"}, false, nil, "-(b)--(a)--(c)-", "-(a)--(b)--(c)-", "-(b)--(a)--(c)-", "-(a)--(c)--(b)-"},
			{"3 elements, c, a, b", []prString{"c", "a", "b"}, false, nil, "-(c)--(a)--(b)-", "-(a)--(b)--(c)-", "-(c)--(a)--(b)-", "-(b)--(a)--(c)-"},
			{"3 elements, a, c, b", []prString{"a", "c", "b"}, false, nil, "-(a)--(c)--(b)-", "-(a)--(b)--(c)-", "-(a)--(c)--(b)-", "-(b)--(c)--(a)-"},
			{"7 elements", []prString{"b", "d", "f", "g", "e", "c", "a"}, false, nil, "-(b)--(a)--(d)--(c)--(f)--(e)--(g)-", "-(a)--(b)--(c)--(d)--(e)--(f)--(g)-", "-(b)--(a)--(d)--(c)--(f)--(e)--(g)-", "-(a)--(c)--(e)--(g)--(f)--(d)--(b)-"},
		}

		for _, test := range tests {
			t.Run(test.name, func(t *testing.T) {
				bst, err := ConstructFromValues[prString](test.input...)

				if err != nil && !test.shouldConstructFail {
					t.Fatalf("ConstructFromValues() failed with unexpected error: %v", err)
				} else if err != nil {
					fmt.Println(err)
				} else {
					gotBFSStr, err2 := bst.TraverseBFS()
					if err2 != nil && !errors.Is(err2, test.expTraverseErr) {
						t.Fatalf("TraverseBFS() failed with unexpected error: %v", err2)
					} else if err2 != nil {
						fmt.Println(err2)
					} else {
						if gotBFSStr != test.expBFSStr {
							t.Fatalf("ConstructFromValues() gave incorrect results, want: %v, got: %v", test.expBFSStr, gotBFSStr)
						}
					}

					gotDFSInOrderStr, err2 := bst.TraverseDFSInOrder()
					if err2 != nil && !errors.Is(err2, test.expTraverseErr) {
						t.Fatalf("TraverseDFSInOrder() failed with unexpected error: %v", err2)
					} else if err2 != nil {
						fmt.Println(err2)
					} else {
						if gotDFSInOrderStr != test.expDFSInOrderStr {
							t.Fatalf("ConstructFromValues() gave incorrect results, want: %v, got: %v", test.expDFSInOrderStr, gotDFSInOrderStr)
						}
					}

					gotDFSPreOrderStr, err2 := bst.TraverseDFSPreOrder()
					if err2 != nil && !errors.Is(err2, test.expTraverseErr) {
						t.Fatalf("TraverseDFSPreOrder() failed with unexpected error: %v", err2)
					} else if err2 != nil {
						fmt.Println(err2)
					} else {
						if gotDFSPreOrderStr != test.expDFSPreOrderStr {
							t.Fatalf("ConstructFromValues() gave incorrect results, want: %v, got: %v", test.expDFSPreOrderStr, gotDFSPreOrderStr)
						}
					}

					gotDFSPostOrderStr, err2 := bst.TraverseDFSPostOrder()
					if err2 != nil && !errors.Is(err2, test.expTraverseErr) {
						t.Fatalf("TraverseDFSPostOrder() failed with unexpected error: %v", err2)
					} else if err2 != nil {
						fmt.Println(err2)
					} else {
						if gotDFSPostOrderStr != test.expDFSPostOrderStr {
							t.Fatalf("ConstructFromValues() gave incorrect results, want: %v, got: %v", test.expDFSPostOrderStr, gotDFSPostOrderStr)
						}
					}
				}
			})
		}
	})
}

func TestSearch(t *testing.T) {

	t.Run("Search prInt", func(t *testing.T) {

		var bst *BinarySearchTree[prInt]
		_, err := bst.Search(0)
		if err == nil {
			t.Fatalf("Search() on a nil tree should have failed")
		} else {
			fmt.Println(err)
		}

		bst = &BinarySearchTree[prInt]{}
		_, err = bst.Search(0)
		if err == nil {
			t.Fatalf("Search() on an empty tree should have failed")
		} else {
			fmt.Println(err)
		}

		bst, err = ConstructFromValues[prInt](7, 4, 9, 5, 1, 0, 2)

		if err != nil {
			t.Fatalf("ConstructFromValues() failed with error: %v", err)
		}

		tests := []struct {
			name      string
			searchVal prInt
			want      bool
		}{
			{"search for 0", 0, true},
			{"search for 1", 1, true},
			{"search for 2", 2, true},
			{"search for 3", 3, false},
			{"search for 4", 4, true},
			{"search for 5", 5, true},
			{"search for 6", 6, false},
			{"search for 7", 7, true},
			{"search for 8", 8, false},
			{"search for 9", 9, true},
		}

		for _, test := range tests {
			t.Run(test.name, func(t *testing.T) {
				got, err2 := bst.Search(test.searchVal)
				if err2 != nil {
					t.Fatalf("Search() failed with error: %v", err2)
				}

				if got != test.want {
					t.Errorf("Search() returned incorrect results, want: %v, got: %v", test.want, got)
				}
			})
		}
	})

	t.Run("Search prString", func(t *testing.T) {

		var bst *BinarySearchTree[prString]
		_, err := bst.Search("a")
		if err == nil {
			t.Fatalf("Search() on a nil tree should have failed")
		} else {
			fmt.Println(err)
		}

		bst = &BinarySearchTree[prString]{}
		_, err = bst.Search("a")
		if err == nil {
			t.Fatalf("Search() on an empty tree should have failed")
		} else {
			fmt.Println(err)
		}

		bst, err = ConstructFromValues[prString]("u", "a", "e", "o", "i")

		if err != nil {
			t.Fatalf("ConstructFromValues() failed with error: %v", err)
		}

		tests := []struct {
			name      string
			searchVal prString
			want      bool
		}{
			{"search for a", "a", true},
			{"search for b", "b", false},
			{"search for c", "c", false},
			{"search for d", "d", false},
			{"search for e", "e", true},
			{"search for f", "f", false},
			{"search for g", "g", false},
			{"search for h", "h", false},
			{"search for i", "i", true},
			{"search for j", "j", false},
			{"search for k", "k", false},
			{"search for l", "l", false},
			{"search for m", "m", false},
			{"search for n", "n", false},
			{"search for o", "o", true},
			{"search for p", "p", false},
			{"search for q", "q", false},
			{"search for r", "r", false},
			{"search for s", "s", false},
			{"search for t", "t", false},
			{"search for u", "u", true},
			{"search for v", "v", false},
			{"search for w", "w", false},
			{"search for x", "x", false},
			{"search for y", "y", false},
			{"search for z", "z", false},
		}

		for _, test := range tests {
			t.Run(test.name, func(t *testing.T) {
				got, err2 := bst.Search(test.searchVal)
				if err2 != nil {
					t.Fatalf("Search() failed with error: %v", err2)
				}

				if got != test.want {
					t.Errorf("Search() returned incorrect results, want: %v, got: %v", test.want, got)
				}
			})
		}
	})

	t.Run("Search prFloat", func(t *testing.T) {

		var bst *BinarySearchTree[prFloat]
		_, err := bst.Search(0.5)
		if err == nil {
			t.Fatalf("Search() on a nil tree should have failed")
		} else {
			fmt.Println(err)
		}

		bst = &BinarySearchTree[prFloat]{}
		_, err = bst.Search(0.5)
		if err == nil {
			t.Fatalf("Search() on an empty tree should have failed")
		} else {
			fmt.Println(err)
		}

		bst, err = ConstructFromValues[prFloat](0.04, 0.02, 0.00, 0.05, 999, 0.07)

		if err != nil {
			t.Fatalf("ConstructFromValues() failed with error: %v", err)
		}

		tests := []struct {
			name      string
			searchVal prFloat
			want      bool
		}{
			{"search for 0", 0, true},
			{"search for 0.01", 0.01, false},
			{"search for 0.02", 0.02, true},
			{"search for 0.03", 0.03, false},
			{"search for 0.04", 0.04, true},
			{"search for 0.05", 0.05, true},
			{"search for 0.06", 0.06, false},
			{"search for 0.07", 0.07, true},
			{"search for 0.08", 0.08, false},
			{"search for 0.09", 0.09, false},
			{"search for 999", 999, true},
		}

		for _, test := range tests {
			t.Run(test.name, func(t *testing.T) {
				got, err2 := bst.Search(test.searchVal)
				if err2 != nil {
					t.Fatalf("Search() failed with error: %v", err2)
				}

				if got != test.want {
					t.Errorf("Search() returned incorrect results, want: %v, got: %v", test.want, got)
				}
			})
		}
	})
}

func TestConstructOrderedSlice(t *testing.T) {

	t.Run("ConstructOrderedSlice prInt", func(t *testing.T) {

		var nilBst *BinarySearchTree[prInt]
		_, err := nilBst.ConstructOrderedSlice()

		if err == nil {
			t.Fatalf("ConstructOrderedSlice() on a nil tree should have failed")
		} else {
			fmt.Println(err)
		}

		tests := []struct {
			name     string
			input    []prInt
			expLen   int
			expSlice []prInt
		}{
			{"nil input", nil, 0, []prInt{}},
			{"empty input", []prInt{}, 0, []prInt{}},
			{"2 elements", []prInt{99999, 1}, 2, []prInt{1, 99999}},
			{"3 elements", []prInt{3, 2, 1}, 3, []prInt{1, 2, 3}},
			{"4 elements", []prInt{4444, -44, 4, -444}, 4, []prInt{-444, -44, 4, 4444}},

			{"3 elements, -1, 0, 1", []prInt{-1, 0, 1}, 3, []prInt{-1, 0, 1}},
			{"3 elements, 1, 0, -1", []prInt{1, 0, -1}, 3, []prInt{-1, 0, 1}},
			{"3 elements, 0, 1, -1", []prInt{0, 1, -1}, 3, []prInt{-1, 0, 1}},
			{"3 elements, 0, -1, 1", []prInt{0, -1, 1}, 3, []prInt{-1, 0, 1}},
			{"3 elements, 1, -1, 0", []prInt{1, -1, 0}, 3, []prInt{-1, 0, 1}},
			{"3 elements, -1, 1, 0", []prInt{-1, 1, 0}, 3, []prInt{-1, 0, 1}},
		}

		for _, test := range tests {
			t.Run(test.name, func(t *testing.T) {

				bst, err1 := ConstructFromValues[prInt](test.input...)
				if err1 != nil {
					t.Fatalf("ConstructFromValues() failed with unexpected error: %v", err1)
				} else {
					sl, err2 := bst.ConstructOrderedSlice()
					if err2 != nil {
						t.Fatalf("ConstructOrderedSlice() failed with unexpected error: %v", err2)
					} else {

						wantLength := test.expLen
						gotLength := len(sl)

						if gotLength != wantLength {
							t.Fatalf("ConstructOrderedSlice() returned a slice of incorrect length, want: %v, got: %v", wantLength, gotLength)
						}

						for i := 0; i < test.expLen; i++ {
							want := test.expSlice[i]
							got := sl[i]

							if got != want {
								t.Errorf("ConstructOrderedSlice() returned a slice with an incorrect value at index: %v, want: %v, got: %v", i, want, got)
							}
						}
					}
				}
			})
		}
	})

	t.Run("ConstructOrderedSlice prString", func(t *testing.T) {

		var nilBst *BinarySearchTree[prString]
		_, err := nilBst.ConstructOrderedSlice()

		if err == nil {
			t.Fatalf("ConstructOrderedSlice() on a nil tree should have failed")
		} else {
			fmt.Println(err)
		}

		tests := []struct {
			name     string
			input    []prString
			expLen   int
			expSlice []prString
		}{
			{"nil input", nil, 0, []prString{}},
			{"empty input", []prString{}, 0, []prString{}},
			{"2 elements", []prString{"hello", "a"}, 2, []prString{"a", "hello"}},
			{"3 elements", []prString{"z", "y", "x"}, 3, []prString{"x", "y", "z"}},
			{"5 elements", []prString{"bye", "off", "by", "of", "buy"}, 5, []prString{"buy", "by", "bye", "of", "off", "oof"}},

			{"3 elements, their, there, they're", []prString{"their", "there", "they're"}, 3, []prString{"their", "there", "they're"}},
			{"3 elements, their, they're, there", []prString{"their", "they're", "there"}, 3, []prString{"their", "there", "they're"}},
			{"3 elements, there, their, they're", []prString{"there", "their", "they're"}, 3, []prString{"their", "there", "they're"}},
			{"3 elements, there, they're, their", []prString{"there", "they're", "their"}, 3, []prString{"their", "there", "they're"}},
			{"3 elements, they're, there, their", []prString{"they're", "there", "their"}, 3, []prString{"their", "there", "they're"}},
			{"3 elements, they're, their, there", []prString{"they're", "their", "there"}, 3, []prString{"their", "there", "they're"}},
		}

		for _, test := range tests {
			t.Run(test.name, func(t *testing.T) {

				bst, err1 := ConstructFromValues[prString](test.input...)
				if err1 != nil {
					t.Fatalf("ConstructFromValues() failed with unexpected error: %v", err1)
				} else {
					sl, err2 := bst.ConstructOrderedSlice()
					if err2 != nil {
						t.Fatalf("ConstructOrderedSlice() failed with unexpected error: %v", err2)
					} else {

						wantLength := test.expLen
						gotLength := len(sl)

						if gotLength != wantLength {
							t.Fatalf("ConstructOrderedSlice() returned a slice of incorrect length, want: %v, got: %v", wantLength, gotLength)
						}

						for i := 0; i < test.expLen; i++ {
							want := test.expSlice[i]
							got := sl[i]

							if got != want {
								t.Errorf("ConstructOrderedSlice() returned a slice with an incorrect value at index: %v, want: %v, got: %v", i, want, got)
							}
						}
					}
				}
			})
		}
	})

	t.Run("ConstructOrderedSlice prFloat", func(t *testing.T) {

		var nilBst *BinarySearchTree[prFloat]
		_, err := nilBst.ConstructOrderedSlice()

		if err == nil {
			t.Fatalf("ConstructOrderedSlice() on a nil tree should have failed")
		} else {
			fmt.Println(err)
		}

		tests := []struct {
			name     string
			input    []prFloat
			expLen   int
			expSlice []prFloat
		}{
			{"nil input", nil, 0, []prFloat{}},
			{"empty input", []prFloat{}, 0, []prFloat{}},
			{"2 elements", []prFloat{1, .99999}, 2, []prFloat{.99999, 1}},
			{"3 elements", []prFloat{.3, .2, .1}, 3, []prFloat{.1, .2, .3}},
			{"4 elements", []prFloat{.4444, -.44, .4, -.444}, 4, []prFloat{-.444, -.44, .4, .4444}},

			{"3 elements, 0.01, 0.0, 0.1", []prFloat{0.01, 0.0, 0.1}, 3, []prFloat{0.0, 0.01, 0.1}},
			{"3 elements, 0.01, 0.1, 0.0", []prFloat{0.01, 0.1, 0.0}, 3, []prFloat{0.0, 0.01, 0.1}},
			{"3 elements, 0.0, 0.01, 0.1", []prFloat{0.0, 0.01, 0.1}, 3, []prFloat{0.0, 0.01, 0.1}},
			{"3 elements, 0.0, 0.1, 0.01", []prFloat{0.0, 0.1, 0.01}, 3, []prFloat{0.0, 0.01, 0.1}},
			{"3 elements, 0.1, 0.01, 0.0", []prFloat{0.1, 0.01, 0.0}, 3, []prFloat{0.0, 0.01, 0.1}},
			{"3 elements, 0.1, 0.0, 0.01", []prFloat{0.1, 0.0, 0.01}, 3, []prFloat{0.0, 0.01, 0.1}},
		}

		for _, test := range tests {
			t.Run(test.name, func(t *testing.T) {

				bst, err1 := ConstructFromValues[prFloat](test.input...)
				if err1 != nil {
					t.Fatalf("ConstructFromValues() failed with unexpected error: %v", err1)
				} else {
					sl, err2 := bst.ConstructOrderedSlice()
					if err2 != nil {
						t.Fatalf("ConstructOrderedSlice() failed with unexpected error: %v", err2)
					} else {

						wantLength := test.expLen
						gotLength := len(sl)

						if gotLength != wantLength {
							t.Fatalf("ConstructOrderedSlice() returned a slice of incorrect length, want: %v, got: %v", wantLength, gotLength)
						}

						for i := 0; i < test.expLen; i++ {
							want := test.expSlice[i]
							got := sl[i]

							if got != want {
								t.Errorf("ConstructOrderedSlice() returned a slice with an incorrect value at index: %v, want: %v, got: %v", i, want, got)
							}
						}
					}
				}
			})
		}
	})
}

func TestBalanceTree(t *testing.T) {

	t.Run("BalanceTree prInt", func(t *testing.T) {
		var bst1 *BinarySearchTree[prInt]

		err1 := bst1.BalanceTree()
		if err1 == nil {
			t.Fatalf("BalanceTree() on a nil tree should have failed")
		} else {
			fmt.Println(err1)
		}

		bst1 = &BinarySearchTree[prInt]{}
		_, expErr := bst1.TraverseBFS()
		if expErr == nil {
			t.Fatalf("TraverseBFS() on an empty tree should have failed")
		}

		err1 = bst1.BalanceTree()
		if err1 != nil {
			t.Fatalf("BalanceTree() failed with unexpected error: %v", err1)
		} else {
			_, gotErr := bst1.TraverseBFS()
			if gotErr == nil {
				t.Fatalf("TraverseBFS() on an empty tree should have failed")
			} else if !errors.Is(gotErr, expErr) {
				t.Fatalf("TraverseBFS() failed with an unexpected error, want: %v, got : %v", expErr, gotErr)
			} else {
				fmt.Println(gotErr)
			}
		}

		tests := []struct {
			name    string
			input   []prInt
			wantBFS string
		}{
			{"3 element tree", []prInt{1, 2, 3}, "-(2)--(1)--(3)-"},
			{"4 element tree", []prInt{1, 2, 3, 4}, "-(2)--(1)--(3)--(4)-"},
			{"5 element tree", []prInt{1, 2, 3, 4, 5}, "-(3)--(1)--(4)--(2)--(5)-"},
			{"6 element tree", []prInt{1, 2, 3, 4, 5, 6}, "-(3)--(1)--(5)--(2)--(4)--(6)-"},
			{"7 element tree", []prInt{1, 2, 3, 4, 5, 6, 7}, "-(4)--(2)--(6)--(1)--(3)--(5)--(7)-"},
			{"8 element tree", []prInt{1, 2, 3, 4, 5, 6, 7, 8}, "-(4)--(2)--(6)--(1)--(3)--(5)--(7)--(8)-"},
			{"9 element tree", []prInt{1, 2, 3, 4, 5, 6, 7, 8, 9}, "-(5)--(2)--(7)--(1)--(3)--(6)--(8)--(4)--(9)-"},
			{"10 element tree", []prInt{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, "-(5)--(2)--(8)--(1)--(3)--(6)--(9)--(4)--(7)--(10)-"},

			{"3 element tree, reversed", []prInt{3, 2, 1}, "-(2)--(1)--(3)-"},
			{"4 element tree, reversed", []prInt{4, 3, 2, 1}, "-(2)--(1)--(3)--(4)-"},
			{"5 element tree, reversed", []prInt{5, 4, 3, 2, 1}, "-(3)--(1)--(4)--(2)--(5)-"},
			{"6 element tree, reversed", []prInt{6, 5, 4, 3, 2, 1}, "-(3)--(1)--(5)--(2)--(4)--(6)-"},
			{"7 element tree, reversed", []prInt{7, 6, 5, 4, 3, 2, 1}, "-(4)--(2)--(6)--(1)--(3)--(5)--(7)-"},
			{"8 element tree, reversed", []prInt{8, 7, 6, 5, 4, 3, 2, 1}, "-(4)--(2)--(6)--(1)--(3)--(5)--(7)--(8)-"},
			{"9 element tree, reversed", []prInt{9, 8, 7, 6, 5, 4, 3, 2, 1}, "-(5)--(2)--(7)--(1)--(3)--(6)--(8)--(4)--(9)-"},
			{"10 element tree, reversed", []prInt{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}, "-(5)--(2)--(8)--(1)--(3)--(6)--(9)--(4)--(7)--(10)-"},

			{"15 element tree", []prInt{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}, "-(8)--(4)--(12)--(2)--(6)--(10)--(14)--(1)--(3)--(5)--(7)--(9)--(11)--(13)--(15)-"},
			{"15 element tree, reversed", []prInt{15, 14, 13, 12, 11, 10, 9, 8, 7, 6, 5, 4, 3, 2, 1}, "-(8)--(4)--(12)--(2)--(6)--(10)--(14)--(1)--(3)--(5)--(7)--(9)--(11)--(13)--(15)-"},
		}

		for _, test := range tests {
			t.Run(test.name, func(t *testing.T) {
				bst, err := ConstructFromValues[prInt](test.input...)
				if err != nil {
					t.Fatalf("ConstructFromValues() failed with unexpected error: %v", err)
				}

				unbalancedDFSInOrder, err := bst.TraverseDFSInOrder()
				if err != nil {
					t.Fatalf("TraverseBFS() failed with an unexpected error, %v", err)
				}

				err = bst.BalanceTree()
				if err != nil {
					t.Fatalf("ConstructFromValues() failed with unexpected error: %v", err)
				}

				wantBFS := test.wantBFS
				gotBFS, err := bst.TraverseBFS()
				if err != nil {
					t.Fatalf("TraverseBFS() failed with an unexpected error, %v", err)
				} else if gotBFS != wantBFS {
					t.Errorf("Post balance BFS tree traversal results are incorrect, want: %v, got %v", wantBFS, gotBFS)
				}

				gotDFSInOrder, err := bst.TraverseDFSInOrder()
				if err != nil {
					t.Fatalf("TraverseDFSInOrder() failed with an unexpected error, %v", err)
				} else if gotDFSInOrder != unbalancedDFSInOrder {
					t.Errorf("DFS inorder tree traversal should return same results before and after balancing, want: %v, got %v", unbalancedDFSInOrder, gotDFSInOrder)
				}
			})
		}
	})

	t.Run("BalanceTree prFloat", func(t *testing.T) {

		var bst1 *BinarySearchTree[prFloat]

		err1 := bst1.BalanceTree()
		if err1 == nil {
			t.Fatalf("BalanceTree() on a nil tree should have failed")
		} else {
			fmt.Println(err1)
		}

		bst1 = &BinarySearchTree[prFloat]{}
		_, expErr := bst1.TraverseBFS()
		if expErr == nil {
			t.Fatalf("TraverseBFS() on an empty tree should have failed")
		}

		err1 = bst1.BalanceTree()
		if err1 != nil {
			t.Fatalf("BalanceTree() failed with unexpected error: %v", err1)
		} else {
			_, gotErr := bst1.TraverseBFS()
			if gotErr == nil {
				t.Fatalf("TraverseBFS() on an empty tree should have failed")
			} else if !errors.Is(gotErr, expErr) {
				t.Fatalf("TraverseBFS() failed with an unexpected error, want: %v, got : %v", expErr, gotErr)
			} else {
				fmt.Println(gotErr)
			}
		}

		tests := []struct {
			name    string
			input   []prFloat
			wantBFS string
		}{
			{"3 elements", []prFloat{0.99999, 1.00001, 1}, "-(1)--(0.99999)--(1.00001)-"},

			{"3 elements 0.01, 0.1, 1", []prFloat{0.01, 0.1, 1}, "-(0.1)--(0.01)--(1)-"},
			{"3 elements 0.01, 1, 0.1", []prFloat{0.01, 1, 0.1}, "-(0.1)--(0.01)--(1)-"},
			{"3 elements 0.1, 0.01, 1", []prFloat{0.1, 0.01, 1}, "-(0.1)--(0.01)--(1)-"},
			{"3 elements 0.1, 1, 0.01", []prFloat{0.1, 1, 0.01}, "-(0.1)--(0.01)--(1)-"},
			{"3 elements 1, 0.01, 0.1", []prFloat{1, 0.01, 0.1}, "-(0.1)--(0.01)--(1)-"},
			{"3 elements 1, 0.1, 0.01", []prFloat{1, 0.1, 0.01}, "-(0.1)--(0.01)--(1)-"},

			{"3 element tree", []prFloat{0.01, 0.02, 0.03}, "-(0.02)--(0.01)--(0.03)-"},
			{"4 element tree", []prFloat{0.01, 0.02, 0.03, 0.04}, "-(0.02)--(0.01)--(0.03)--(0.04)-"},
			{"5 element tree", []prFloat{0.01, 0.02, 0.03, 0.04, 0.05}, "-(0.03)--(0.01)--(0.04)--(0.02)--(0.05)-"},
			{"6 element tree", []prFloat{0.01, 0.02, 0.03, 0.04, 0.05, 0.06}, "-(0.03)--(0.01)--(0.05)--(0.02)--(0.04)--(0.06)-"},
			{"7 element tree", []prFloat{0.01, 0.02, 0.03, 0.04, 0.05, 0.06, 0.07}, "-(0.04)--(0.02)--(0.06)--(0.01)--(0.03)--(0.05)--(0.07)-"},
			{"8 element tree", []prFloat{0.01, 0.02, 0.03, 0.04, 0.05, 0.06, 0.07, 0.08}, "-(0.04)--(0.02)--(0.06)--(0.01)--(0.03)--(0.05)--(0.07)--(0.08)-"},
			{"9 element tree", []prFloat{0.01, 0.02, 0.03, 0.04, 0.05, 0.06, 0.07, 0.08, 0.09}, "-(0.05)--(0.02)--(0.07)--(0.01)--(0.03)--(0.06)--(0.08)--(0.04)--(0.09)-"},
			{"10 element tree", []prFloat{0.01, 0.02, 0.03, 0.04, 0.05, 0.06, 0.07, 0.08, 0.09, 0.1}, "-(0.05)--(0.02)--(0.08)--(0.01)--(0.03)--(0.06)--(0.09)--(0.04)--(0.07)--(0.1)-"},

			{"3 element tree, reversed", []prFloat{0.03, 0.02, 0.01}, "-(0.02)--(0.01)--(0.03)-"},
			{"4 element tree, reversed", []prFloat{0.04, 0.03, 0.02, 0.01}, "-(0.02)--(0.01)--(0.03)--(0.04)-"},
			{"5 element tree, reversed", []prFloat{0.05, 0.04, 0.03, 0.02, 0.01}, "-(0.03)--(0.01)--(0.04)--(0.02)--(0.05)-"},
			{"6 element tree, reversed", []prFloat{0.06, 0.05, 0.04, 0.03, 0.02, 0.01}, "-(0.03)--(0.01)--(0.05)--(0.02)--(0.04)--(0.06)-"},
			{"7 element tree, reversed", []prFloat{0.07, 0.06, 0.05, 0.04, 0.03, 0.02, 0.01}, "-(0.04)--(0.02)--(0.06)--(0.01)--(0.03)--(0.05)--(0.07)-"},
			{"8 element tree, reversed", []prFloat{0.08, 0.07, 0.06, 0.05, 0.04, 0.03, 0.02, 0.01}, "-(0.04)--(0.02)--(0.06)--(0.01)--(0.03)--(0.05)--(0.07)--(0.08)-"},
			{"9 element tree, reversed", []prFloat{0.09, 0.08, 0.07, 0.06, 0.05, 0.04, 0.03, 0.02, 0.01}, "-(0.05)--(0.02)--(0.07)--(0.01)--(0.03)--(0.06)--(0.08)--(0.04)--(0.09)-"},
			{"10 element tree, reversed", []prFloat{0.1, 0.09, 0.08, 0.07, 0.06, 0.05, 0.04, 0.03, 0.02, 0.01}, "-(0.05)--(0.02)--(0.08)--(0.01)--(0.03)--(0.06)--(0.09)--(0.04)--(0.07)--(0.1)-"},

			{"15 element tree", []prFloat{0.1, 0.2, 0.3, 0.4, 0.5, 0.6, 0.7, 0.8, 0.9, 1, 1.1, 1.2, 1.3, 1.4, 1.5}, "-(0.8)--(0.4)--(1.2)--(0.2)--(0.6)--(1)--(1.4)--(0.1)--(0.3)--(0.5)--(0.7)--(0.9)--(1.1)--(1.3)--(1.5)-"},
			{"15 element tree, reversed", []prFloat{1.5, 1.4, 1.3, 1.2, 1.1, 1, 0.9, 0.8, 0.7, 0.6, 0.5, 0.4, 0.3, 0.2, 0.1}, "-(0.8)--(0.4)--(1.2)--(0.2)--(0.6)--(1)--(1.4)--(0.1)--(0.3)--(0.5)--(0.7)--(0.9)--(1.1)--(1.3)--(1.5)-"},
		}

		for _, test := range tests {
			t.Run(test.name, func(t *testing.T) {
				bst, err := ConstructFromValues[prFloat](test.input...)
				if err != nil {
					t.Fatalf("ConstructFromValues() failed with unexpected error: %v", err)
				}

				unbalancedDFSInOrder, err := bst.TraverseDFSInOrder()
				if err != nil {
					t.Fatalf("TraverseBFS() failed with an unexpected error, %v", err)
				}

				err = bst.BalanceTree()
				if err != nil {
					t.Fatalf("ConstructFromValues() failed with unexpected error: %v", err)
				}

				wantBFS := test.wantBFS
				gotBFS, err := bst.TraverseBFS()
				if err != nil {
					t.Fatalf("TraverseBFS() failed with an unexpected error, %v", err)
				} else if gotBFS != wantBFS {
					t.Errorf("Post balance BFS tree traversal results are incorrect, want: %v, got %v", wantBFS, gotBFS)
				}

				gotDFSInOrder, err := bst.TraverseDFSInOrder()
				if err != nil {
					t.Fatalf("TraverseDFSInOrder() failed with an unexpected error, %v", err)
				} else if gotDFSInOrder != unbalancedDFSInOrder {
					t.Errorf("DFS inorder tree traversal should return same results before and after balancing, want: %v, got %v", unbalancedDFSInOrder, gotDFSInOrder)
				}
			})
		}
	})

	t.Run("BalanceTree prString", func(t *testing.T) {
		var bst1 *BinarySearchTree[prString]

		err1 := bst1.BalanceTree()
		if err1 == nil {
			t.Fatalf("BalanceTree() on a nil tree should have failed")
		} else {
			fmt.Println(err1)
		}

		bst1 = &BinarySearchTree[prString]{}
		_, expErr := bst1.TraverseBFS()
		if expErr == nil {
			t.Fatalf("TraverseBFS() on an empty tree should have failed")
		}

		err1 = bst1.BalanceTree()
		if err1 != nil {
			t.Fatalf("BalanceTree() failed with unexpected error: %v", err1)
		} else {
			_, gotErr := bst1.TraverseBFS()
			if gotErr == nil {
				t.Fatalf("TraverseBFS() on an empty tree should have failed")
			} else if !errors.Is(gotErr, expErr) {
				t.Fatalf("TraverseBFS() failed with an unexpected error, want: %v, got : %v", expErr, gotErr)
			} else {
				fmt.Println(gotErr)
			}
		}

		tests := []struct {
			name    string
			input   []prString
			wantBFS string
		}{
			{"basic test case", []prString{"a", "b", "c"}, "-(b)--(a)--(c)-"},

			{"3 elements: I, me, myself", []prString{"I", "me", "myself"}, "-(me)--(I)--(myself)-"},
			{"3 elements: I, myself, me", []prString{"I", "myself", "me"}, "-(me)--(I)--(myself)-"},
			{"3 elements: me, I, myself", []prString{"me", "I", "myself"}, "-(me)--(I)--(myself)-"},
			{"3 elements: me, myself, I", []prString{"me", "myself", "I"}, "-(me)--(I)--(myself)-"},
			{"3 elements: myself, I, me", []prString{"myself", "I", "me"}, "-(me)--(I)--(myself)-"},
			{"3 elements: myself, me, I", []prString{"myself", "me", "I"}, "-(me)--(I)--(myself)-"},

			{"3 element tree", []prString{"a", "an", "any"}, "-(an)--(a)--(any)-"},
			{"4 element tree", []prString{"a", "an", "any", "ain't"}, "-(ain't)--(a)--(an)--(any)-"},
			{"5 element tree", []prString{"a", "an", "any", "ain't", "aren't"}, "-(an)--(a)--(any)--(ain't)--(aren't)-"},
			{"6 element tree", []prString{"a", "an", "any", "ain't", "aren't", "are not"}, "-(an)--(a)--(are not)--(ain't)--(any)--(aren't)-"},
			{"7 element tree", []prString{"a", "an", "any", "ain't", "aren't", "are not", "at least"}, "-(any)--(ain't)--(aren't)--(a)--(an)--(are not)--(at least)-"},
			{"8 element tree", []prString{"a", "an", "any", "ain't", "aren't", "are not", "at least", "although"}, "-(an)--(ain't)--(are not)--(a)--(although)--(any)--(aren't)--(at least)-"},
			{"9 element tree", []prString{"a", "an", "any", "ain't", "aren't", "are not", "at least", "although", "along with"}, "-(an)--(ain't)--(are not)--(a)--(along with)--(any)--(aren't)--(although)--(at least)-"},
			{"10 element tree", []prString{"a", "an", "any", "ain't", "aren't", "are not", "at least", "although", "along with", "altogether"}, "-(altogether)--(ain't)--(are not)--(a)--(along with)--(an)--(aren't)--(although)--(any)--(at least)-"},

			{"3 element tree, reversed", []prString{"any", "an", "a"}, "-(an)--(a)--(any)-"},
			{"4 element tree, reversed", []prString{"ain't", "any", "an", "a"}, "-(ain't)--(a)--(an)--(any)-"},
			{"5 element tree, reversed", []prString{"aren't", "ain't", "any", "an", "a"}, "-(an)--(a)--(any)--(ain't)--(aren't)-"},
			{"6 element tree, reversed", []prString{"are not", "aren't", "ain't", "any", "an", "a"}, "-(an)--(a)--(are not)--(ain't)--(any)--(aren't)-"},
			{"7 element tree, reversed", []prString{"at least", "are not", "aren't", "ain't", "any", "an", "a"}, "-(any)--(ain't)--(aren't)--(a)--(an)--(are not)--(at least)-"},
			{"8 element tree, reversed", []prString{"although", "at least", "are not", "aren't", "ain't", "any", "an", "a"}, "-(an)--(ain't)--(are not)--(a)--(although)--(any)--(aren't)--(at least)-"},
			{"9 element tree, reversed", []prString{"along with", "although", "at least", "are not", "aren't", "ain't", "any", "an", "a"}, "-(an)--(ain't)--(are not)--(a)--(along with)--(any)--(aren't)--(although)--(at least)-"},
			{"10 element tree, reversed", []prString{"altogether", "along with", "although", "at least", "are not", "aren't", "ain't", "any", "an", "a"}, "-(altogether)--(ain't)--(are not)--(a)--(along with)--(an)--(aren't)--(although)--(any)--(at least)-"},

			{"15 element tree", []prString{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o"}, "-(h)--(d)--(l)--(b)--(f)--(j)--(n)--(a)--(c)--(e)--(g)--(i)--(k)--(m)--(o)-"},
			{"15 element tree", []prString{"o", "n", "m", "l", "k", "j", "i", "h", "g", "f", "e", "d", "c", "b", "a"}, "-(h)--(d)--(l)--(b)--(f)--(j)--(n)--(a)--(c)--(e)--(g)--(i)--(k)--(m)--(o)-"},
		}

		for _, test := range tests {
			t.Run(test.name, func(t *testing.T) {
				bst, err := ConstructFromValues[prString](test.input...)
				if err != nil {
					t.Fatalf("ConstructFromValues() failed with unexpected error: %v", err)
				}

				unbalancedDFSInOrder, err := bst.TraverseDFSInOrder()
				if err != nil {
					t.Fatalf("TraverseBFS() failed with an unexpected error, %v", err)
				}

				err = bst.BalanceTree()
				if err != nil {
					t.Fatalf("ConstructFromValues() failed with unexpected error: %v", err)
				}

				wantBFS := test.wantBFS
				gotBFS, err := bst.TraverseBFS()
				if err != nil {
					t.Fatalf("TraverseBFS() failed with an unexpected error, %v", err)
				} else if gotBFS != wantBFS {
					t.Errorf("Post balance BFS tree traversal results are incorrect, want: %v, got %v", wantBFS, gotBFS)
				}

				gotDFSInOrder, err := bst.TraverseDFSInOrder()
				if err != nil {
					t.Fatalf("TraverseDFSInOrder() failed with an unexpected error, %v", err)
				} else if gotDFSInOrder != unbalancedDFSInOrder {
					t.Errorf("DFS inorder tree traversal should return same results before and after balancing, want: %v, got %v", unbalancedDFSInOrder, gotDFSInOrder)
				}
			})
		}
	})
}

func TestConstructBalancedTree(t *testing.T) {

	_, err1 := ConstructBalancedTree[prInt]()
	if err1 == nil {
		t.Fatalf("ConstructBalancedTree() should have failed on empty input")
	} else {
		fmt.Println(err1)
	}

	bst1, err1 := ConstructBalancedTree[prInt](1, 2, 3)
	if err1 != nil {
		t.Fatalf("ConstructBalancedTree() failed with an unexpected error, %v", err1)
	} else {
		wantBFS := "-(2)--(1)--(3)-"
		gotBFS, err2 := bst1.TraverseBFS()
		if err2 != nil {
			t.Fatalf("TraverseDBFS() failed with an unexpected error, %v", err2)
		} else if gotBFS != wantBFS {
			t.Errorf("BFS tree traversal results are incorrect, want: %v, got %v", wantBFS, gotBFS)
		}

		wantDFSInOrder := "-(1)--(2)--(3)-"
		gotDFSInOrder, err2 := bst1.TraverseDFSInOrder()
		if err2 != nil {
			t.Fatalf("TraverseDFSInOrder() failed with an unexpected error, %v", err2)
		} else if gotDFSInOrder != wantDFSInOrder {
			t.Errorf("DFS In Order traversal results are incorrect, want: %v, got %v", wantDFSInOrder, gotDFSInOrder)
		}

		wantDFSPreOrder := "-(2)--(1)--(3)-"
		gotDFSPreOrder, err2 := bst1.TraverseDFSPreOrder()
		if err2 != nil {
			t.Fatalf("TraverseDFSPreOrder() failed with an unexpected error, %v", err2)
		} else if gotDFSPreOrder != wantDFSPreOrder {
			t.Errorf("DFS Pre Order tree traversal results are incorrect, want: %v, got %v", wantDFSPreOrder, gotDFSPreOrder)
		}

		wantDFSPostOrder := "-(1)--(3)--(2)-"
		gotDFSPostOrder, err2 := bst1.TraverseDFSPostOrder()
		if err2 != nil {
			t.Fatalf("TraverseDFSPostOrder() failed with an unexpected error, %v", err2)
		} else if gotDFSPostOrder != wantDFSPostOrder {
			t.Errorf("DFS Post Order tree traversal results are incorrect, want: %v, got %v", wantDFSPostOrder, gotDFSPostOrder)
		}
	}

	t.Run("ConstructBalancedTree() prInt", func(t *testing.T) {

	})
}
