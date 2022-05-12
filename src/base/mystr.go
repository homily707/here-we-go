package base

import (
	"bufio"
	"bytes"
	"fmt"
	"strings"
)

func mystr() {

	//str
	s := "一二三1234567"
	fmt.Printf("%c, %x, %U \n", s[0], s[0], s[0])
	fmt.Printf("%c, %x, %U \n", s[3], s[3], s[3])
	fmt.Printf("%c, %x, %U \n", s[6], s[6], s[6])
	for i, v := range s {
		fmt.Printf("%d : %c, %x, %d \n", i, v, v, v)
		// i = 0 3 9 10 11 12 13
		// v = 一 二 三 1 2 3 4 5
	}

	for i := 0; i < len(s); i++ {
		v := s[i]
		// equals to bytes[i]
		fmt.Printf("%d : %c, %x, %d \n", i, v, v, v)
	}

	bytes := ([]byte)(s)
	for i := 0; i < len(bytes); i++ {
		v := bytes[i]
		fmt.Printf("%d : %c, %x, %d \n", i, v, v, v)
	}
}

func mystrings() {

	builder := strings.Builder{}
	builder.Write([]byte("hello,"))
	builder.WriteByte('c')
	builder.WriteRune('d')
	builder.WriteString("f")
	fmt.Println(builder.String())

	reader := strings.NewReader("hello")
	fmt.Println("before read ,Len() : ", reader.Len())
	readByte, err := reader.ReadByte()
	if err == nil {
		fmt.Print(string(readByte))
		fmt.Print(" after read", reader.Len())
	} else {
		fmt.Errorf(" %s", err)
	}
}

func mybytes() {
	buffer := bytes.Buffer{}
	buffer.WriteString("123456789")
	_, _ = buffer.ReadByte()
}

func mybufio() {
	reader := bufio.NewReader(strings.NewReader("hello"))
	reader.ReadByte()
	reader.Peek(1)
}
