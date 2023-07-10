package base

import (
	"fmt"
	"testing"
)

func Test_myslice(t *testing.T) {
	x := 2
	y := 7
	z := y / x
	fmt.Println("\033[31m[error]\033[0m ", z)
}
