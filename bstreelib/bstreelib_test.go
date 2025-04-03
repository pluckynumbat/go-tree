package bstreelib

import (
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

		_, err := n1.Parent()
		if err == nil {
			t.Errorf("calling Parent() on a nil node should have returned an error")
		} else {
			fmt.Println(err)
		}

		n2 = &Node[prInt]{1, nil, nil, nil}
		got, err := n2.Parent()
		if err != nil {
			t.Fatalf("Parent() failed with error: %v", err)
		} else {
			if got != nil {
				t.Errorf("Parent() returned incorrect results, want: nil, got: %v", got)
			}
		}

		n3 = &Node[prInt]{2, n2, nil, nil}
		parent, err := n3.Parent()
		if err != nil {
			t.Fatalf("Parent() failed with error: %v", err)
		} else if parent == nil {
			t.Fatalf("Parent() returned unexpected nil value")
		} else {
			want := "1"
			got := parent.String()

			if got != want {
				t.Errorf("Parent() returned incorrect results, want: %v, got: %v", want, got)
			}
		}
	})
}
