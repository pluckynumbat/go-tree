package bintreelib

import (
	"testing"
)

func TestNodeString(t *testing.T) {
	var node *Node
	want := "nil"
	got := node.String()

	if got != want {
		t.Errorf("String() returned incorrect results, want: %v, got %v", want, got)
	}

	node = &Node{}
	want = ""
	got = node.String()
	if got != want {
		t.Errorf("String() returned incorrect results, want: %v, got %v", want, got)
	}

	node = &Node{nil, "a", nil}
	want = "a"
	got = node.String()
	if got != want {
		t.Errorf("String() returned incorrect results, want: %v, got %v", want, got)
	}
}
