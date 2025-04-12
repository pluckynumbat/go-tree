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

			// gather expected parent prInt pointers
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

				actuaLeftChild, err2 := runner.LeftChild()
				if err2 != nil {
					t.Fatalf("LeftChild() failed with error: %v", err2)
				}
				got := "nil"
				if actuaLeftChild != nil {
					got = actuaLeftChild.String()
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
