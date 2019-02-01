package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name string
	Age  int
}

//使用 Golang 解析 JSON  格式数据时，若以 interface{} 接收数据，则会按照下列规则进行解析：
// bool, for JSON booleans

// float64, for JSON numbers

// string, for JSON strings

// []interface{}, for JSON arrays

// map[string]interface{}, for JSON objects

// nil for JSON null

func main() {

	json := `{
		"name":"lizhongweei",
		"age" : 18
	}`
	fmt.Println("json test")
}
