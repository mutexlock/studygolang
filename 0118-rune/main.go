package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

//rune 为int32
//byte  int8
func main() {
	//var hh string
	str := "hello 你好"
	fmt.Println(str[0])
	fmt.Println(strings)
	for i, v := range str {
		fmt.Println(i, v, str[i])
	}
	by := []byte(str)
	fmt.Println(len([]rune(str)))
	fmt.Println(len([]byte(str)))
	fmt.Println(by)
	fmt.Println(utf8.RuneCountInString(str))

}
