package bintreelib

import (
	"testing"
)

func TestNodeString(t *testing.T) {
	var n1, n2, n3 *Node

	n2 = &Node{}
	n3 = &Node{nil, "a", nil}

	tests := []struct {
		name string
		node *Node
		want string
	}{
		{"nil node", n1, "nil"},
		{"non nil empty node", n2, ""},
		{"non nil non empty node", n3, "a"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := test.node.String()
			want := test.want
			if got != want {
				t.Errorf("String() returned incorrect results, want: %v, got %v", want, got)
			}
		})
	}
}

func TestIsNil(t *testing.T) {
	var bt1 *BinaryTree
	bt2 := &BinaryTree{}

	tests := []struct {
		name string
		bt   *BinaryTree
		want bool
	}{
		{"nil true", bt1, true},
		{"nil false", bt2, false},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := test.bt.IsNil()
			want := test.want
			if got != want {
				t.Errorf("IsNil() returned incorrect results, want: %v, got %v", want, got)
			}
		})
	}
}

func TestIsEmpty(t *testing.T) {
	var bt1 *BinaryTree
	bt2 := &BinaryTree{}

	root := &Node{}
	bt3 := &BinaryTree{root}

	tests := []struct {
		name string
		bt   *BinaryTree
		want bool
	}{
		{"nil binary tree", bt1, true},
		{"non nil, empty binary tree", bt2, true},
		{"non nil, non empty binary tree", bt3, false},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := test.bt.IsEmpty()
			want := test.want
			if got != want {
				t.Errorf("IsEmpty() returned incorrect results, want: %v, got %v", want, got)
			}
		})
	}
}

