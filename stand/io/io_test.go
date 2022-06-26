package io

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"
)

func TestIO(t *testing.T) {
	file, err := os.Open("../go.mod")
	if err != nil {

	}
	bytes, _ := ioutil.ReadAll(file)
	fmt.Println(string(bytes))

}
