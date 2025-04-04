package bstreelib

import (
	"errors"
	"fmt"
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
