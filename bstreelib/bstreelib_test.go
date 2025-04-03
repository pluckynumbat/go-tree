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
}
