package base

import "os"

func myfile() {
	file, _ := os.Create("myfile")
	file.WriteString("hello")
	file.Close()
}
