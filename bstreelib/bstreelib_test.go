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
	bst3 = &BinarySearchTree[prInt]{root}

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
	bst3 = &BinarySearchTree[prInt]{r1}

	r2 := &Node[prInt]{2, nil, nil, nil}
	n2 := &Node[prInt]{1, r2, nil, nil}
	r2.left = n2
	bst4 = &BinarySearchTree[prInt]{r2}

	r3 := &Node[prInt]{0, nil, nil, nil}
	n4 := &Node[prInt]{-1, r3, nil, nil}
	n5 := &Node[prInt]{1, r3, nil, nil}
	r3.left = n4
	r3.right = n5
	bst5 = &BinarySearchTree[prInt]{r3}

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

func TestInsert(t *testing.T) {

	t.Run("test insert on Binary Search Tree of prInt nodes", func(t *testing.T) {
		var bst1, bst2 *BinarySearchTree[prInt]
		bst2 = &BinarySearchTree[prInt]{}

		tests := []struct {
			name             string
			bst              *BinarySearchTree[prInt]
			val              prInt
			expError         error
			expBFStr         string
			expDFSInorderStr string
		}{
			{"nil tree", bst1, 1, treeNilError, "", ""},
			{"empty tree", bst2, 1, nil, "-(1)-", "-(1)-"},
			{"1 element tree", bst2, 4, nil, "-(1)--(4)-", "-(1)--(4)-"},
			{"2 element tree", bst2, 6, nil, "-(1)--(4)--(6)-", "-(1)--(4)--(6)-"},
			{"3 element tree", bst2, 2, nil, "-(1)--(4)--(2)--(6)-", "-(1)--(2)--(4)--(6)-"},
			{"4 element tree", bst2, 5, nil, "-(1)--(4)--(2)--(6)--(5)-", "-(1)--(2)--(4)--(5)--(6)-"},
			{"5 element tree", bst2, 3, nil, "-(1)--(4)--(2)--(6)--(3)--(5)-", "-(1)--(2)--(3)--(4)--(5)--(6)-"},
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
			name             string
			bst              *BinarySearchTree[prString]
			val              prString
			expError         error
			expBFStr         string
			expDFSInorderStr string
		}{
			{"nil tree", bst1, "b", treeNilError, "", ""},
			{"empty tree", bst2, "b", nil, "-(b)-", "-(b)-"},
			{"1 element tree", bst2, "d", nil, "-(b)--(d)-", "-(b)--(d)-"},
			{"2 element tree", bst2, "f", nil, "-(b)--(d)--(f)-", "-(b)--(d)--(f)-"},
			{"3 element tree", bst2, "a", nil, "-(b)--(a)--(d)--(f)-", "-(a)--(b)--(d)--(f)-"},
			{"4 element tree", bst2, "c", nil, "-(b)--(a)--(d)--(c)--(f)-", "-(a)--(b)--(c)--(d)--(f)-"},
			{"5 element tree", bst2, "e", nil, "-(b)--(a)--(d)--(c)--(f)--(e)-", "-(a)--(b)--(c)--(d)--(e)--(f)-"},
			{"6 element tree", bst2, "g", nil, "-(b)--(a)--(d)--(c)--(f)--(e)--(g)-", "-(a)--(b)--(c)--(d)--(e)--(f)--(g)-"},
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
		}{
			{"nil input", nil, false, treeEmptyError, "", ""},
			{"empty input", []prInt{}, false, treeEmptyError, "", ""},
			{"2 elements, identical", []prInt{1, 1}, true, nil, "", ""},
			{"3 elements, -1, 0, 1", []prInt{-1, 0, 1}, false, nil, "-(-1)--(0)--(1)-", "-(-1)--(0)--(1)-"},
			{"3 elements, 1, 0, -1", []prInt{1, 0, -1}, false, nil, "-(1)--(0)--(-1)-", "-(-1)--(0)--(1)-"},
			{"3 elements, 0, 1, -1", []prInt{0, 1, -1}, false, nil, "-(0)--(-1)--(1)-", "-(-1)--(0)--(1)-"},
			{"3 elements, 0, -1, 1", []prInt{0, -1, 1}, false, nil, "-(0)--(-1)--(1)-", "-(-1)--(0)--(1)-"},
			{"3 elements, 1, -1, 0", []prInt{1, -1, 0}, false, nil, "-(1)--(-1)--(0)-", "-(-1)--(0)--(1)-"},
			{"3 elements, -1, 1, 0", []prInt{-1, 1, 0}, false, nil, "-(-1)--(1)--(0)-", "-(-1)--(0)--(1)-"},
			{"7 elements, all positive", []prInt{2, 4, 6, 7, 5, 3, 1}, false, nil, "-(2)--(1)--(4)--(3)--(6)--(5)--(7)-", "-(1)--(2)--(3)--(4)--(5)--(6)--(7)-"},
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
		}{
			{"nil input", nil, false, treeEmptyError, "", ""},
			{"empty input", []prString{}, false, treeEmptyError, "", ""},
			{"2 elements, identical", []prString{"a", "a"}, true, nil, "", ""},
			{"3 elements, a, b, c", []prString{"a", "b", "c"}, false, nil, "-(a)--(b)--(c)-", "-(a)--(b)--(c)-"},
			{"3 elements, c, b, a", []prString{"c", "b", "a"}, false, nil, "-(c)--(b)--(a)-", "-(a)--(b)--(c)-"},
			{"3 elements, b, c, a", []prString{"b", "c", "a"}, false, nil, "-(b)--(a)--(c)-", "-(a)--(b)--(c)-"},
			{"3 elements, b, a, c", []prString{"b", "a", "c"}, false, nil, "-(b)--(a)--(c)-", "-(a)--(b)--(c)-"},
			{"3 elements, c, a, b", []prString{"c", "a", "b"}, false, nil, "-(c)--(a)--(b)-", "-(a)--(b)--(c)-"},
			{"3 elements, a, c, b", []prString{"a", "c", "b"}, false, nil, "-(a)--(c)--(b)-", "-(a)--(b)--(c)-"},
			{"7 elements", []prString{"b", "d", "f", "g", "e", "c", "a"}, false, nil, "-(b)--(a)--(d)--(c)--(f)--(e)--(g)-", "-(a)--(b)--(c)--(d)--(e)--(f)--(g)-"},
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
	})
}
