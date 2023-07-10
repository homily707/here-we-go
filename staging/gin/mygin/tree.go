package mygin

type node struct {
}

type methodTree struct {
	method string
	root   *node
}

type methodTrees []methodTree
