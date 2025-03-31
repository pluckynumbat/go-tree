package bstreelib

import "fmt"

type prInt int // printable int
func (p prInt) String() string {
	return fmt.Sprintf("%v", int(p))
}
