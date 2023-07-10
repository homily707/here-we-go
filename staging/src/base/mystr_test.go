package base

import (
	"strings"
	"testing"
)

func Test_mystr(t *testing.T) {
	//mystrings()
	uniqueEmail()
}

func Test_myos(t *testing.T) {
	myfile()
}

func Test_mynet(t *testing.T) {
	httpserver()
}

func Test_split(t *testing.T) {
	path1 := "/a/b/c/d"
	path2 := "/a/b/c/d/"
	path3 := "a/b/c/d"
	for _, path := range []string{path1, path2, path3} {
		s := strings.SplitAfter(path, "/")
		t.Log(len(s), s)
	}
}
