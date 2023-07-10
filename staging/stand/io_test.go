package stand

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"testing"
	"unicode"
)

func Test_IO(t *testing.T) {
	file, err := os.Open("../go.mod")
	if err != nil {

	}
	bytes, _ := ioutil.ReadAll(file)
	bufio.NewReader(file)
	fmt.Println(string(bytes))
}

func Test_Str(t *testing.T) {
	unicode.IsDigit('f')
}
