package main

import (
	"fmt"
	//	"strings"
	//	"sync"
	"unicode/utf8"
)

//rune 为int32
//byte  int8
func main() {
	//var hh string
	str := "hello 你好"
	fmt.Println(str[0])
	for i, v := range str {
		fmt.Println(i, v, str[i])
	}
	by := []byte(str)
	fmt.Println(len([]rune(str)))
	fmt.Println(len([]byte(str)))
	fmt.Println(by)
	fmt.Println(utf8.RuneCountInString(str))

	var s string

	//string 不能和nil比较
	if s == "" {
		fmt.Println("is Empty", len(s))
	}

}

// type WaitGroupWrapper struct {
// 	sync.WaitGroup
// }

// func (w *WaitGroupWrapper) Wapper(f func()) {
// 	w.Add()
// 	go func() {
// 		f()
// 		w.Done()
// 	}()
// }
