package main

import (
	"fmt"
	"io/ioutil"

	"github.com/BurntSushi/toml"
)

var localConfig map[string]interface{}

func main() {
	content, err := ioutil.ReadFile("config.toml")
	if err != nil {
		panic(err)
	}
	if _, err := toml.Decode(string(content), &localConfig); err != nil {
		panic(err)
	}

	for key, value := range localConfig {
		fmt.Printf("%v, %v\n", key, value)
	}

	fmt.Println(GetLocalIntSliceConfig("new_member_id"))

}

func GetLocalStringConfig(name string) string {
	if v, ok := localConfig[name]; ok {
		return v.(string)
	}
	return ""
}

func GetLocalIntSliceConfig(name string) []int {
	intSlice := make([]int, 0)
	if v, ok := localConfig[name]; ok {
		switch v.(type) {
		case []interface{}:
			for _, i := range v.([]interface{}) {
				switch i.(type) {
				case int64:
					intSlice = append(intSlice, int(i.(int64)))
				default:
					break
				}
			}
		}
	}

	return intSlice
}
