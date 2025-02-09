package bintreelib

import (
	"fmt"
	"github.com/pluckynumbat/go-quez/sgquezlib"
	"testing"
)

func TestNodeString(t *testing.T) {
	var n1, n2, n3 *Node

	n2 = &Node{}
	n3 = &Node{"a", nil, nil, nil}

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

func TestNodeParent(t *testing.T) {
	var n1, n2, n3 *Node

	n2 = &Node{"a", nil, nil, nil}
	n3 = &Node{"b", n2, nil, nil}

	tests := []struct {
		name      string
		node      *Node
		expError  error
		parent    *Node
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
					t.Fatalf("GetParent() failed with unexpected error: %v", err)
				}
			} else if parent != test.parent {
				t.Errorf("Parent() returned incorrect parent pointer, want: %v, got: %v", test.parent, parent)
			} else if parent.String() != test.parentStr {
				t.Errorf("GetParent() returned parent pointer with incorrect string, want: %v, got: %v", test.parentStr, parent.String())
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
	bt3 := &BinaryTree{root, root}

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

	r1 := &Node{"1", nil, nil, nil}
	bt3 = &BinaryTree{r1, r1}

	n2 := &Node{"b", nil, nil, nil}
	r2 := &Node{"a", nil, n2, nil}
	n2.parent = r2

	bt4 = &BinaryTree{r2, n2}

	n3 := &Node{"l", nil, nil, nil}
	n4 := &Node{"r", nil, nil, nil}
	r3 := &Node{"m", nil, n3, n4}
	n3.parent = r3
	n4.parent = r3

	bt5 = &BinaryTree{r3, n4}

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

func TestLastLeaf(t *testing.T) {
	var bt *BinaryTree

	node := bt.LastLeaf()
	if node != nil {
		t.Fatalf("Last leaf of a nil tree should be nil")
	}

	bt = &BinaryTree{}
	node = bt.LastLeaf()
	if node != nil {
		t.Fatalf("Last leaf of an empty tree should be nil")
	}

	tests := []struct {
		name   string
		addVal string
		want   string
	}{
		{"1 element binary tree", "a", "a"},
		{"2 element binary tree", "b", "b"},
		{"3 element binary tree", "c", "c"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := bt.AddNodeBFS(test.addVal)
			if err != nil {
				t.Fatalf("AddNodeBFS() failed with error: %v", err)
			}

			got := bt.LastLeaf().String()
			if got != test.want {
				t.Errorf("LastLeaf() returned incorrect results, want: %v, got: %v", test.want, got)
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

	bt = &BinaryTree{}
	vals := []string{"a", "b", "c", "d", "e"}
	for _, str := range vals {
		err := bt.AddNodeBFS(str)
		if err != nil {
			t.Fatalf("AddNode() failed with error: %v", err)
		}
	}

	queue := &sgquezlib.SemiGenericQueue[*Node]{}
	err2 := queue.Enqueue(bt.root)
	if err2 != nil {
		t.Fatalf("Enqueue() failed with error: %v", err2)
	}
	expVals := []string{"a", "b", "c", "d", "e"}

	strCnt := 0
	for !queue.IsEmpty() {
		val, err := queue.Dequeue()
		if err != nil {
			t.Fatalf("Dequeue() failed with error: %v", err)
		}

		got := val.String()
		want := expVals[strCnt]
		strCnt++

		// since it is BFS, nodes should be visited in the same order as they were added to the tree
		if got != want {
			t.Errorf("AddNode() has incorrect results, want: %v, got: %v", want, got)
		}

		if val.left != nil {
			err2 := queue.Enqueue(val.left)
			if err2 != nil {
				t.Fatalf("Enqueue() failed with error: %v", err2)
			}
		}

		if val.right != nil {
			err2 := queue.Enqueue(val.right)
			if err2 != nil {
				t.Fatalf("Enqueue() failed with error: %v", err2)
			}
		}
	}
}

func TestConstructFromValues(t *testing.T) {

	vals := []string{"a", "b", "c", "d", "e"}
	bt, err := ConstructFromValues(vals...)
	if err != nil {
		t.Fatalf("ConstructFromValues() failed with error: %v", err)
	}

	queue := &sgquezlib.SemiGenericQueue[*Node]{}
	err2 := queue.Enqueue(bt.root)
	if err2 != nil {
		t.Fatalf("Enqueue() failed with error: %v", err2)
	}
	expVals := []string{"a", "b", "c", "d", "e"}

	strCnt := 0
	for !queue.IsEmpty() {
		val, err := queue.Dequeue()
		if err != nil {
			t.Fatalf("Dequeue() failed with error: %v", err)
		}

		got := val.String()
		want := expVals[strCnt]
		strCnt++

		// since it is BFS, nodes should be visited in the same order as they were added to the tree
		if got != want {
			t.Errorf("ConstructFromValues() has incorrect results, want: %v, got: %v", want, got)
		}

		if val.left != nil {
			err2 := queue.Enqueue(val.left)
			if err2 != nil {
				t.Fatalf("Enqueue() failed with error: %v", err2)
			}
		}

		if val.right != nil {
			err2 := queue.Enqueue(val.right)
			if err2 != nil {
				t.Fatalf("Enqueue() failed with error: %v", err2)
			}
		}
	}
}

func TestTraverseBFS(t *testing.T) {
	var bt *BinaryTree

	_, err := bt.TraverseBFS()
	if err == nil {
		t.Error("TraverseBFS() on a nil tree should return an error")
	} else {
		fmt.Println(err)
	}

	bt = &BinaryTree{}
	_, err = bt.TraverseBFS()
	if err == nil {
		t.Error("TraverseBFS() on an empty tree should return an error")
	} else {
		fmt.Println(err)
	}

	strs := make([]string, 0, 6)
	tests := []struct {
		name   string
		newVal string
		want   string
	}{
		{"1 element tree", "a", "-a-"},
		{"2 element tree", "b", "-a--b-"},
		{"3 element tree", "c", "-a--b--c-"},
		{"4 element tree", "d", "-a--b--c--d-"},
		{"5 element tree", "e", "-a--b--c--d--e-"},
		{"6 element tree", "f", "-a--b--c--d--e--f-"},
		{"7 element tree", "g", "-a--b--c--d--e--f--g-"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			strs = append(strs, test.newVal)
			bt, err := ConstructFromValues(strs...)
			if err != nil {
				t.Errorf("ConstructFromValues() failed with error: %v", err)
			} else {
				got, err2 := bt.TraverseBFS()
				if err2 != nil {
					t.Errorf("TraverseBFS() failed with error: %v", err2)
				} else {
					want := test.want
					if got != want {
						t.Errorf("TraverseBFS() returned incorrect results, want: %v, got: %v", want, got)
					}
				}
			}
		})
	}
}

func TestTraverseDFSPreOrderRecursive(t *testing.T) {
	var bt *BinaryTree

	_, err := bt.TraverseDFSPreOrderRecursive()
	if err == nil {
		t.Error("TraverseDFSPreOrderRecursive() on a nil tree should return an error")
	} else {
		fmt.Println(err)
	}

	bt = &BinaryTree{}
	_, err = bt.TraverseDFSPreOrderRecursive()
	if err == nil {
		t.Error("TraverseDFSPreOrderRecursive() on an empty tree should return an error")
	} else {
		fmt.Println(err)
	}

	strs := make([]string, 0, 6)
	tests := []struct {
		name   string
		newVal string
		want   string
	}{
		{"1 element tree", "a", "-a-"},
		{"2 element tree", "b", "-a--b-"},
		{"3 element tree", "c", "-a--b--c-"},
		{"4 element tree", "d", "-a--b--d--c-"},
		{"5 element tree", "e", "-a--b--d--e--c-"},
		{"6 element tree", "f", "-a--b--d--e--c--f-"},
		{"7 element tree", "g", "-a--b--d--e--c--f--g-"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			strs = append(strs, test.newVal)
			bt, err := ConstructFromValues(strs...)
			if err != nil {
				t.Errorf("ConstructFromValues() failed with error: %v", err)
			} else {
				got, err2 := bt.TraverseDFSPreOrderRecursive()
				if err2 != nil {
					t.Errorf("TraverseDFSPreOrderRecursive() failed with error: %v", err2)
				} else {
					want := test.want
					if got != want {
						t.Errorf("TraverseDFSPreOrderRecursive() returned incorrect results, want: %v, got: %v", want, got)
					}
				}
			}
		})
	}
}

func TestTraverseDFSPreOrderIterative(t *testing.T) {
	var bt *BinaryTree

	_, err := bt.TraverseDFSPreOrderIterative()
	if err == nil {
		t.Error("TraverseDFSPreOrderIterative() on a nil tree should return an error")
	} else {
		fmt.Println(err)
	}

	bt = &BinaryTree{}
	_, err = bt.TraverseDFSPreOrderIterative()
	if err == nil {
		t.Error("TraverseDFSPreOrderIterative() on an empty tree should return an error")
	} else {
		fmt.Println(err)
	}

	strs := make([]string, 0, 6)
	tests := []struct {
		name   string
		newVal string
		want   string
	}{
		{"1 element tree", "a", "-a-"},
		{"2 element tree", "b", "-a--b-"},
		{"3 element tree", "c", "-a--b--c-"},
		{"4 element tree", "d", "-a--b--d--c-"},
		{"5 element tree", "e", "-a--b--d--e--c-"},
		{"6 element tree", "f", "-a--b--d--e--c--f-"},
		{"7 element tree", "g", "-a--b--d--e--c--f--g-"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			strs = append(strs, test.newVal)
			bt, err := ConstructFromValues(strs...)
			if err != nil {
				t.Errorf("ConstructFromValues() failed with error: %v", err)
			} else {
				got, err2 := bt.TraverseDFSPreOrderIterative()
				if err2 != nil {
					t.Errorf("TraverseDFSPreOrderIterative() failed with error: %v", err2)
				} else {
					want := test.want
					if got != want {
						t.Errorf("TraverseDFSPreOrderIterative() returned incorrect results, want: %v, got: %v", want, got)
					}
				}
			}
		})
	}
}

func TestTraverseDFSInOrderRecursive(t *testing.T) {
	var bt *BinaryTree

	_, err := bt.TraverseDFSInOrderRecursive()
	if err == nil {
		t.Error("TraverseDFSPreOrderRecursive() on a nil tree should return an error")
	} else {
		fmt.Println(err)
	}

	bt = &BinaryTree{}
	_, err = bt.TraverseDFSInOrderRecursive()
	if err == nil {
		t.Error("TraverseDFSPreOrderRecursive() on an empty tree should return an error")
	} else {
		fmt.Println(err)
	}

	strs := make([]string, 0, 6)
	tests := []struct {
		name   string
		newVal string
		want   string
	}{
		{"1 element tree", "a", "-a-"},
		{"2 element tree", "b", "-b--a-"},
		{"3 element tree", "c", "-b--a--c-"},
		{"4 element tree", "d", "-d--b--a--c-"},
		{"5 element tree", "e", "-d--b--e--a--c-"},
		{"6 element tree", "f", "-d--b--e--a--f--c-"},
		{"7 element tree", "g", "-d--b--e--a--f--c--g-"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			strs = append(strs, test.newVal)
			bt, err := ConstructFromValues(strs...)
			if err != nil {
				t.Errorf("ConstructFromValues() failed with error: %v", err)
			} else {
				got, err2 := bt.TraverseDFSInOrderRecursive()
				if err2 != nil {
					t.Errorf("TraverseDFSInOrderRecursive() failed with error: %v", err2)
				} else {
					want := test.want
					if got != want {
						t.Errorf("TraverseDFSInOrderRecursive() returned incorrect results, want: %v, got: %v", want, got)
					}
				}
			}
		})
	}
}

func TestTraverseDFSInOrderIterative(t *testing.T) {
	var bt *BinaryTree

	_, err := bt.TraverseDFSInOrderIterative()
	if err == nil {
		t.Error("TraverseDFSInOrderIterative() on a nil tree should return an error")
	} else {
		fmt.Println(err)
	}

	bt = &BinaryTree{}
	_, err = bt.TraverseDFSPreOrderIterative()
	if err == nil {
		t.Error("TraverseDFSInOrderIterative() on an empty tree should return an error")
	} else {
		fmt.Println(err)
	}

	strs := make([]string, 0, 6)
	tests := []struct {
		name   string
		newVal string
		want   string
	}{
		{"1 element tree", "a", "-a-"},
		{"2 element tree", "b", "-a--b-"},
		{"3 element tree", "c", "-a--b--c-"},
		{"4 element tree", "d", "-a--b--d--c-"},
		{"5 element tree", "e", "-a--b--d--e--c-"},
		{"6 element tree", "f", "-a--b--d--e--c--f-"},
		{"7 element tree", "g", "-a--b--d--e--c--f--g-"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			strs = append(strs, test.newVal)
			bt, err := ConstructFromValues(strs...)
			if err != nil {
				t.Errorf("ConstructFromValues() failed with error: %v", err)
			} else {
				got, err2 := bt.TraverseDFSPreOrderIterative()
				if err2 != nil {
					t.Errorf("TraverseDFSInOrderIterative() failed with error: %v", err2)
				} else {
					want := test.want
					if got != want {
						t.Errorf("TraverseDFSInOrderIterative() returned incorrect results, want: %v, got: %v", want, got)
					}
				}
			}
		})
	}
}

func TestTraverseDFSPostOrderRecursive(t *testing.T) {
	var bt *BinaryTree

	_, err := bt.TraverseDFSPostOrderRecursive()
	if err == nil {
		t.Error("TraverseDFSPostOrderRecursive() on a nil tree should return an error")
	} else {
		fmt.Println(err)
	}

	bt = &BinaryTree{}
	_, err = bt.TraverseDFSPostOrderRecursive()
	if err == nil {
		t.Error("TraverseDFSPostOrderRecursive() on an empty tree should return an error")
	} else {
		fmt.Println(err)
	}

	strs := make([]string, 0, 6)
	tests := []struct {
		name   string
		newVal string
		want   string
	}{
		{"1 element tree", "a", "-a-"},
		{"2 element tree", "b", "-b--a-"},
		{"3 element tree", "c", "-b--c--a-"},
		{"4 element tree", "d", "-d--b--c--a-"},
		{"5 element tree", "e", "-d--e--b--c--a-"},
		{"6 element tree", "f", "-d--e--b--f--c--a-"},
		{"7 element tree", "g", "-d--e--b--f--g--c--a-"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			strs = append(strs, test.newVal)
			bt, err := ConstructFromValues(strs...)
			if err != nil {
				t.Errorf("ConstructFromValues() failed with error: %v", err)
			} else {
				got, err2 := bt.TraverseDFSPostOrderRecursive()
				if err2 != nil {
					t.Errorf("TraverseDFSPostOrderRecursive() failed with error: %v", err2)
				} else {
					want := test.want
					if got != want {
						t.Errorf("TraverseDFSPostOrderRecursive() returned incorrect results, want: %v, got: %v", want, got)
					}
				}
			}
		})
	}
}

func TestTraverseDFSPostOrderIterative(t *testing.T) {
	var bt *BinaryTree

	_, err := bt.TraverseDFSPostOrderIterative()
	if err == nil {
		t.Error("TraverseDFSPostOrderIterative() on a nil tree should return an error")
	} else {
		fmt.Println(err)
	}

	bt = &BinaryTree{}
	_, err = bt.TraverseDFSPostOrderIterative()
	if err == nil {
		t.Error("TraverseDFSPostOrderIterative() on an empty tree should return an error")
	} else {
		fmt.Println(err)
	}

	strs := make([]string, 0, 6)
	tests := []struct {
		name   string
		newVal string
		want   string
	}{
		{"1 element tree", "a", "-a-"},
		{"2 element tree", "b", "-b--a-"},
		{"3 element tree", "c", "-b--c--a-"},
		{"4 element tree", "d", "-d--b--c--a-"},
		{"5 element tree", "e", "-d--e--b--c--a-"},
		{"6 element tree", "f", "-d--e--b--f--c--a-"},
		{"7 element tree", "g", "-d--e--b--f--g--c--a-"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			strs = append(strs, test.newVal)
			bt, err := ConstructFromValues(strs...)
			if err != nil {
				t.Errorf("ConstructFromValues() failed with error: %v", err)
			} else {
				got, err2 := bt.TraverseDFSPostOrderIterative()
				if err2 != nil {
					t.Errorf("TraverseDFSPostOrderIterative() failed with error: %v", err2)
				} else {
					want := test.want
					if got != want {
						t.Errorf("TraverseDFSPostOrderIterative() returned incorrect results, want: %v, got: %v", want, got)
					}
				}
			}
		})
	}
}
