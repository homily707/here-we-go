package main

import "fmt"

type MyStruct struct {
	Field1 string
	Field2 int
}

func (s MyStruct) UpdateField(newVal string) {
	s.UpdateFieldByPointer(newVal)
}

func (s *MyStruct) UpdateFieldByPointer(newVal string) {
	s.Field1 = newVal
}

func main() {
	origin := MyStruct{
		Field1: "origin",
		Field2: 0,
	}
	fmt.Println(origin)
	p := &origin
	p.UpdateField("updated")
	fmt.Println(origin)
	origin.UpdateField("ufdfd")
	fmt.Println(origin)
}