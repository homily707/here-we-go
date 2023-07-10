package base

import (
	"fmt"
	"reflect"
)

type User struct {
	id   int
	id64 int64
	name string
}

func value() {
	user := User{
		id:   0,
		id64: 0,
		name: "",
	}
	typ := reflect.TypeOf(user)
	f1 := typ.Field(1)
	//f2 := typ.FieldByIndex([]int{1, 0})
	fmt.Println(f1.Name)
}
