package bstreelib

import (
	"fmt"
	"testing"
)

type prInt int // printable int
func (p prInt) String() string {
	return fmt.Sprintf("%v", int(p))
}

func TestNodeString(t *testing.T) {
	node := &Node[prInt]{1, nil, nil, nil}

	want := "1"
	got := node.String()

	if got != want {
		t.Errorf("Node's string returned incorrect results, want: %v, got %v", want, got)
	}
}
