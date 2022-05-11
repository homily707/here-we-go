package main

import "fmt"

func myrune() {

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
