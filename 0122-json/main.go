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
	j := `{
  "enable": 1,
  "enabled_province": [
    "110000",
    "650000",
    "610000",
    "130000",
    "430000",
    "230000",
    "520000",
    "640000",
    "530000",
    "340000",
    "440000",
    "210000",
    "440100",
    "630000",
    "620000",
    "140000",
    "410000",
    "360000",
    "220000",
    "310000",
    "120000",
    "150000",
    "320000",
    "330000",
    "450000",
    "460000",
    "510000",
    "420000",
    "370000",
    "350000",
    "500000",
    "540000"
  ],
  "white_list": {
    "100009": [
      "110000",
      "650000",
      "610000",
      "130000",
      "430000",
      "230000",
      "520000",
      "640000",
      "530000",
      "340000",
      "440000",
      "210000",
      "440100",
      "630000",
      "620000",
      "140000",
      "410000",
      "360000",
      "220000",
      "310000",
      "120000",
      "150000",
      "320000",
      "330000",
      "450000",
      "460000",
      "510000",
      "420000",
      "370000",
      "350000",
      "500000",
      "540000"
    ],
    "348742": [
      "110000",
      "650000",
      "610000",
      "130000",
      "430000",
      "230000",
      "520000",
      "640000",
      "530000",
      "340000",
      "440000",
      "210000",
      "440100",
      "630000",
      "620000",
      "140000",
      "410000",
      "360000",
      "220000",
      "120000",
      "150000",
      "320000",
      "330000",
      "450000",
      "460000",
      "510000",
      "420000",
      "370000",
      "350000",
      "500000"
    ],
    "9288868": [
      "110000",
      "650000",
      "130000",
      "610000",
      "430000",
      "230000",
      "520000",
      "640000",
      "530000",
      "340000",
      "440000",
      "210000",
      "440100",
      "630000",
      "620000",
      "140000",
      "410000",
      "360000",
      "220000",
      "310000",
      "120000",
      "150000",
      "320000",
      "330000",
      "450000",
      "460000",
      "510000",
      "420000",
      "370000",
      "350000",
      "500000"
    ],
    "10684583": [
      "110000",
      "650000",
      "610000",
      "130000",
      "430000",
      "230000",
      "520000",
      "640000",
      "530000",
      "340000",
      "440000",
      "210000",
      "440100",
      "630000",
      "620000",
      "140000",
      "410000",
      "360000",
      "220000",
      "310000",
      "120000",
      "150000",
      "320000",
      "330000",
      "450000",
      "460000",
      "510000",
      "420000",
      "370000",
      "350000",
      "500000"
    ],
    "13970678": [
      "110000",
      "650000",
      "610000",
      "130000",
      "430000",
      "230000",
      "520000",
      "640000",
      "530000",
      "340000",
      "440000",
      "210000",
      "440100",
      "630000",
      "620000",
      "140000",
      "410000",
      "360000",
      "220000",
      "310000",
      "120000",
      "150000",
      "320000",
      "330000",
      "450000",
      "460000",
      "510000",
      "420000",
      "370000",
      "350000",
      "500000"
    ],
    "15236502": [
      "110000",
      "650000",
      "610000",
      "130000",
      "430000",
      "230000",
      "520000",
      "640000",
      "530000",
      "340000",
      "440000",
      "210000",
      "440100",
      "630000",
      "620000",
      "140000",
      "410000",
      "360000",
      "220000",
      "310000",
      "120000",
      "150000",
      "320000",
      "330000",
      "450000",
      "460000",
      "510000",
      "420000",
      "370000",
      "350000",
      "500000"
    ],
    "16062055": [
      "110000",
      "650000",
      "610000",
      "130000",
      "430000",
      "230000",
      "520000",
      "640000",
      "530000",
      "340000",
      "440000",
      "210000",
      "440100",
      "630000",
      "620000",
      "140000",
      "410000",
      "360000",
      "220000",
      "120000",
      "150000",
      "320000",
      "330000",
      "450000",
      "460000",
      "510000",
      "420000",
      "370000",
      "350000",
      "500000",
      "540000"
    ],
    "16667285": [
      "650000",
      "610000",
      "130000",
      "430000",
      "230000",
      "520000",
      "640000",
      "530000",
      "340000",
      "440000",
      "210000",
      "440100",
      "630000",
      "620000",
      "140000",
      "410000",
      "360000",
      "220000",
      "310000",
      "120000",
      "150000",
      "320000",
      "330000",
      "450000",
      "460000",
      "510000",
      "420000",
      "370000",
      "350000",
      "500000",
      "540000"
    ],
    "16981187": [
      "650000",
      "610000",
      "130000",
      "430000",
      "230000",
      "520000",
      "640000",
      "530000",
      "340000",
      "440000",
      "210000",
      "440100",
      "630000",
      "620000",
      "140000",
      "410000",
      "360000",
      "220000",
      "120000",
      "150000",
      "310000",
      "320000",
      "330000",
      "450000",
      "460000",
      "510000",
      "420000",
      "370000",
      "350000",
      "500000",
      "540000"
    ],
    "17946345": [
      "110000",
      "650000",
      "610000",
      "130000",
      "430000",
      "230000",
      "520000",
      "640000",
      "530000",
      "340000",
      "440000",
      "210000",
      "440100",
      "630000",
      "620000",
      "140000",
      "410000",
      "360000",
      "220000",
      "310000",
      "120000",
      "150000",
      "320000",
      "330000",
      "450000",
      "460000",
      "510000",
      "420000",
      "370000",
      "350000",
      "500000",
      "540000"
    ],
    "50753119": [
      "110000",
      "650000",
      "610000",
      "130000",
      "430000",
      "230000",
      "520000",
      "640000",
      "530000",
      "340000",
      "440000",
      "210000",
      "440100",
      "630000",
      "620000",
      "140000",
      "410000",
      "360000",
      "220000",
      "310000",
      "120000",
      "150000",
      "320000",
      "330000",
      "450000",
      "460000",
      "510000",
      "420000",
      "370000",
      "350000",
      "500000",
      "540000"
    ],
    "50812556": [
      "110000",
      "650000",
      "610000",
      "130000",
      "430000",
      "230000",
      "520000",
      "640000",
      "530000",
      "340000",
      "440000",
      "210000",
      "440100",
      "630000",
      "620000",
      "140000",
      "410000",
      "360000",
      "220000",
      "310000",
      "120000",
      "150000",
      "320000",
      "330000",
      "450000",
      "460000",
      "510000",
      "420000",
      "370000",
      "350000",
      "500000"
    ],
    "51266935": [
      "110000",
      "650000",
      "610000",
      "130000",
      "430000",
      "230000",
      "520000",
      "640000",
      "530000",
      "340000",
      "440000",
      "210000",
      "440100",
      "630000",
      "620000",
      "140000",
      "410000",
      "360000",
      "220000",
      "310000",
      "120000",
      "150000",
      "320000",
      "330000",
      "450000",
      "460000",
      "510000",
      "420000",
      "370000",
      "350000",
      "500000"
    ],
    "51428320": [
      "110000",
      "650000",
      "610000",
      "130000",
      "430000",
      "230000",
      "520000",
      "640000",
      "530000",
      "340000",
      "440000",
      "210000",
      "440100",
      "630000",
      "620000",
      "140000",
      "410000",
      "360000",
      "220000",
      "310000",
      "120000",
      "150000",
      "320000",
      "330000",
      "450000",
      "460000",
      "510000",
      "420000",
      "370000",
      "350000",
      "500000"
    ],
    "51536677": [
      "110000",
      "650000",
      "610000",
      "130000",
      "430000",
      "230000",
      "520000",
      "640000",
      "530000",
      "340000",
      "440000",
      "210000",
      "440100",
      "630000",
      "620000",
      "140000",
      "410000",
      "360000",
      "220000",
      "310000",
      "120000",
      "150000",
      "320000",
      "330000",
      "450000",
      "460000",
      "510000",
      "420000",
      "370000",
      "350000",
      "500000"
    ],
    "51657680": [
      "110000",
      "650000",
      "610000",
      "130000",
      "430000",
      "230000",
      "520000",
      "640000",
      "530000",
      "340000",
      "440000",
      "210000",
      "440100",
      "630000",
      "620000",
      "140000",
      "410000",
      "360000",
      "220000",
      "310000",
      "120000",
      "150000",
      "320000",
      "330000",
      "450000",
      "460000",
      "510000",
      "420000",
      "370000",
      "350000",
      "500000"
    ],
    "52315982": [
      "110000",
      "650000",
      "610000",
      "130000",
      "430000",
      "230000",
      "520000",
      "640000",
      "530000",
      "340000",
      "440000",
      "210000",
      "440100",
      "630000",
      "620000",
      "140000",
      "410000",
      "360000",
      "220000",
      "310000",
      "120000",
      "150000",
      "320000",
      "330000",
      "450000",
      "460000",
      "510000",
      "420000",
      "370000",
      "350000",
      "500000"
    ],
    "52407341": [
      "110000",
      "650000",
      "610000",
      "130000",
      "430000",
      "230000",
      "520000",
      "640000",
      "530000",
      "340000",
      "440000",
      "210000",
      "440100",
      "630000",
      "620000",
      "140000",
      "410000",
      "360000",
      "220000",
      "310000",
      "120000",
      "150000",
      "320000",
      "330000",
      "450000",
      "460000",
      "510000",
      "420000",
      "370000",
      "350000",
      "500000"
    ],
    "52670386": [
      "110000",
      "650000",
      "610000",
      "130000",
      "430000",
      "230000",
      "520000",
      "640000",
      "530000",
      "340000",
      "440000",
      "210000",
      "440100",
      "630000",
      "620000",
      "140000",
      "410000",
      "360000",
      "220000",
      "310000",
      "120000",
      "150000",
      "320000",
      "330000",
      "450000",
      "460000",
      "510000",
      "420000",
      "370000",
      "350000",
      "500000"
    ],
    "53804564": [
      "110000",
      "650000",
      "610000",
      "130000",
      "430000",
      "230000",
      "520000",
      "640000",
      "530000",
      "340000",
      "440000",
      "210000",
      "440100",
      "630000",
      "620000",
      "140000",
      "410000",
      "360000",
      "220000",
      "310000",
      "120000",
      "150000",
      "320000",
      "330000",
      "450000",
      "460000",
      "510000",
      "420000",
      "370000",
      "350000",
      "500000"
    ],
    "55352431": [
      "110000",
      "610000",
      "650000",
      "130000",
      "430000",
      "230000",
      "520000",
      "640000",
      "530000",
      "340000",
      "440000",
      "210000",
      "440100",
      "630000",
      "620000",
      "140000",
      "410000",
      "360000",
      "220000",
      "310000",
      "120000",
      "150000",
      "320000",
      "330000",
      "450000",
      "460000",
      "510000",
      "420000",
      "370000",
      "350000",
      "500000",
      "540000"
    ],
    "58579983": [
      "110000",
      "650000",
      "610000",
      "130000",
      "430000",
      "230000",
      "520000",
      "640000",
      "530000",
      "340000",
      "440000",
      "210000",
      "440100",
      "630000",
      "620000",
      "140000",
      "410000",
      "360000",
      "220000",
      "310000",
      "120000",
      "150000",
      "320000",
      "330000",
      "450000",
      "460000",
      "510000",
      "420000",
      "370000",
      "350000",
      "500000"
    ],
    "59988391": [
      "110000",
      "650000",
      "610000",
      "130000",
      "430000",
      "230000",
      "520000",
      "640000",
      "530000",
      "340000",
      "440000",
      "210000",
      "440100",
      "630000",
      "620000",
      "140000",
      "410000",
      "360000",
      "220000",
      "310000",
      "120000",
      "150000",
      "320000",
      "330000",
      "450000",
      "460000",
      "510000",
      "420000",
      "370000",
      "350000",
      "500000"
    ],
    "60478051": [
      "110000",
      "650000",
      "610000",
      "130000",
      "430000",
      "230000",
      "520000",
      "640000",
      "530000",
      "340000",
      "440000",
      "210000",
      "440100",
      "630000",
      "620000",
      "140000",
      "410000",
      "360000",
      "220000",
      "310000",
      "120000",
      "150000",
      "320000",
      "330000",
      "450000",
      "460000",
      "510000",
      "420000",
      "370000",
      "350000",
      "500000",
      "540000"
    ],
    "77824984": [
      "650000",
      "610000",
      "130000",
      "430000",
      "230000",
      "520000",
      "640000",
      "530000",
      "340000",
      "440000",
      "210000",
      "440100",
      "630000",
      "620000",
      "140000",
      "410000",
      "360000",
      "220000",
      "310000",
      "120000",
      "150000",
      "320000",
      "330000",
      "450000",
      "460000",
      "510000",
      "420000",
      "370000",
      "350000",
      "500000",
      "540000"
    ],
    "77938181": [
      "610000",
      "650000",
      "130000",
      "430000",
      "230000",
      "520000",
      "640000",
      "530000",
      "340000",
      "440000",
      "210000",
      "440100",
      "630000",
      "620000",
      "140000",
      "410000",
      "360000",
      "220000",
      "310000",
      "120000",
      "150000",
      "320000",
      "330000",
      "450000",
      "460000",
      "510000",
      "420000",
      "370000",
      "350000",
      "500000",
      "540000"
    ],
    "97714622": [
      "110000",
      "650000",
      "610000",
      "130000",
      "430000",
      "230000",
      "520000",
      "640000",
      "530000",
      "340000",
      "440000",
      "210000",
      "440100",
      "630000",
      "620000",
      "140000",
      "410000",
      "360000",
      "220000",
      "310000",
      "120000",
      "150000",
      "320000",
      "330000",
      "450000",
      "460000",
      "510000",
      "420000",
      "370000",
      "350000",
      "500000",
      "540000"
    ],
    "101843542": [
      "110000",
      "650000",
      "610000",
      "130000",
      "430000",
      "230000",
      "520000",
      "640000",
      "530000",
      "340000",
      "440000",
      "210000",
      "440100",
      "630000",
      "620000",
      "140000",
      "410000",
      "360000",
      "220000",
      "310000",
      "120000",
      "150000",
      "320000",
      "330000",
      "450000",
      "460000",
      "510000",
      "420000",
      "370000",
      "350000",
      "500000",
      "540000"
    ],
    "105856837": [
      "650000",
      "610000",
      "130000",
      "430000",
      "230000",
      "520000",
      "640000",
      "530000",
      "340000",
      "440000",
      "210000",
      "440100",
      "630000",
      "620000",
      "140000",
      "410000",
      "360000",
      "220000",
      "120000",
      "150000",
      "310000",
      "320000",
      "330000",
      "450000",
      "460000",
      "510000",
      "420000",
      "370000",
      "350000",
      "500000"
    ],
    "113882029": [
      "650000",
      "610000",
      "130000",
      "430000",
      "230000",
      "520000",
      "640000",
      "530000",
      "340000",
      "440000",
      "210000",
      "440100",
      "630000",
      "620000",
      "140000",
      "410000",
      "360000",
      "220000",
      "310000",
      "120000",
      "150000",
      "320000",
      "330000",
      "450000",
      "460000",
      "510000",
      "420000",
      "370000",
      "350000",
      "500000",
      "540000"
    ],
    "225591856": [
      "650000",
      "610000",
      "130000",
      "430000",
      "230000",
      "520000",
      "640000",
      "530000",
      "340000",
      "440000",
      "210000",
      "440100",
      "630000",
      "620000",
      "140000",
      "410000",
      "360000",
      "220000",
      "310000",
      "120000",
      "150000",
      "320000",
      "330000",
      "450000",
      "460000",
      "510000",
      "420000",
      "370000",
      "350000",
      "500000",
      "540000"
    ],
    "227185893": [
      "110000",
      "650000",
      "130000",
      "610000",
      "430000",
      "230000",
      "520000",
      "640000",
      "530000",
      "340000",
      "440000",
      "210000",
      "440100",
      "630000",
      "620000",
      "140000",
      "410000",
      "360000",
      "220000",
      "310000",
      "120000",
      "150000",
      "320000",
      "330000",
      "450000",
      "460000",
      "510000",
      "420000",
      "370000",
      "350000",
      "500000",
      "540000"
    ],
    "229203163": [
      "650000",
      "610000",
      "130000",
      "430000",
      "230000",
      "520000",
      "640000",
      "530000",
      "340000",
      "440000",
      "210000",
      "440100",
      "630000",
      "620000",
      "140000",
      "410000",
      "360000",
      "220000",
      "310000",
      "120000",
      "150000",
      "320000",
      "330000",
      "450000",
      "460000",
      "510000",
      "420000",
      "370000",
      "350000",
      "500000",
      "540000"
    ],
    "429050470": [
      "110000",
      "650000",
      "610000",
      "130000",
      "430000",
      "230000",
      "520000",
      "640000",
      "530000",
      "340000",
      "440000",
      "210000",
      "440100",
      "630000",
      "620000",
      "140000",
      "410000",
      "360000",
      "220000",
      "120000",
      "150000",
      "320000",
      "330000",
      "450000",
      "460000",
      "510000",
      "420000",
      "370000",
      "350000",
      "500000",
      "540000"
    ],
    "434133360": [
      "110000",
      "650000",
      "610000",
      "130000",
      "430000",
      "230000",
      "520000",
      "640000",
      "530000",
      "340000",
      "440000",
      "210000",
      "440100",
      "630000",
      "620000",
      "140000",
      "410000",
      "360000",
      "220000",
      "310000",
      "120000",
      "150000",
      "320000",
      "330000",
      "450000",
      "460000",
      "510000",
      "420000",
      "370000",
      "350000",
      "500000",
      "540000"
    ]
  }
}`
	json := `{
		"name":"lizhongweei",
		"age" : 18
	}`
	fmt.Println("json test")
}
