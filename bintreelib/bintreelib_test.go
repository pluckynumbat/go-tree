package bintreelib

import (
	"fmt"
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

func TestRoot(t *testing.T) {

	var bt1, bt2, bt3, bt4, bt5 *BinaryTree
	bt2 = &BinaryTree{}

	r1 := &Node{nil, "1", nil}
	bt3 = &BinaryTree{r1}

	n2 := &Node{nil, "b", nil}
	r2 := &Node{n2, "a", nil}
	bt4 = &BinaryTree{r2}

	n3 := &Node{nil, "l", nil}
	n4 := &Node{nil, "r", nil}
	r3 := &Node{n3, "m", n4}
	bt5 = &BinaryTree{r3}

	tests := []struct {
		name string
		bt   *BinaryTree
		want string
	}{
		{"nil binary tree", bt1, "nil"},
		{"empty binary tree", bt2, "nil"},
		{"1 element binary tree", bt3, "1"},
		{"2 element binary tree", bt4, "a"},
		{"3 element binary tree", bt5, "m"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			want := test.want
			got := test.bt.Root().String()
			if got != want {
				t.Errorf("Root() returned incorrect results, want: %v, got %v", want, got)
			}
		})
	}
}

func TestAddNodeBFS(t *testing.T) {
	var bt *BinaryTree
	err := bt.AddNodeBFS("a")
	if err == nil {
		t.Error("AddNode() on a nil Binary Tree should have returned an error")
	} else {
		fmt.Println(err)
	}
}