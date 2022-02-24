package main

func main() {
	str := "abcdef"
	print(str[0] == 'a')
	for _, v := range str {
		print(v == 'a')
	}
}
