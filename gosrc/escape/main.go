package main

import (
	"bytes"
	"io"
)

type UserData struct {
    Name  string
}

// func main() {
//     var info UserData
//     info.Name = "WilburXu"
//     _ = GetUserInfo(info)
// }

// func GetUserInfo(userInfo UserData) *UserData {
//     return &userInfo
// }

func print(w io.Writer, s string) {
	asBytes := []byte(s)
	w.Write(asBytes)
}

func bufWriter(w *bytes.Buffer, s string) {
	asBytes := []byte(s)
	w.Write(asBytes)
}

func main() {
	buf := bytes.Buffer{}
	bufWriter(&buf, "foobar")
}