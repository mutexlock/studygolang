package main

import (
	"encoding/json"
	"fmt"
	. "studygolang/0201-rpc-client/user_fields"

	//	"time"
	simplejson "github.com/bitly/go-simplejson"

	"github.com/hprose/hprose-golang/rpc"
)

// type HelloService struct {
// 	GetFieldsByJson func(GetFieldsByJsonReq) GetFieldsByJsonResp
// }

// type GetFieldsByJsonReq struct {
// 	MemberId int64
// 	Fields   []string
// }

// type GetFieldsByJsonResp struct {
// 	MemberId int64
// 	JsonVal  string
// }

type Info struct {
	Version_name string `json:"version_name"`
	Channle      string `json:"channel"`
}
type Info11 struct {
	Info1 Info `json:"info"`
}

func GetInfoField(member map[string]string, field string) string {
	data := ""
	if value, ok := member["info"]; ok {
		valueJson, err := simplejson.NewJson([]byte(value))
		if err != nil {
			return data
		}
		data = GetRawMapField(valueJson.MustMap(), field, "").(string)
	}
	return data
}

func GetRawMapField(aMap map[string]interface{}, field string, defaultValue interface{}) interface{} {
	if item, ok := aMap[field]; ok {
		return item
	}
	return defaultValue
}

func stringTointerface() interface{} {
	strMap := "sdf"
	return strMap
}

func main() {
	client := rpc.NewHTTPClient("http://test1-passport.qttcs3.cn/user-fields")
	client.Header.Set("X-AppId", "PASSPORT_QUKAN_APPID")

	//client.SetTimeout(time.Duration(200) * time.Millisecond)

	var helloService *UserFieldsClientService

	client.UseService(&helloService)
	//	getFieldsByJsonReq := GetFieldsByJsonReq{MemberId: 2147700139, Fields: []string{"telephone", "nickname", "avatar", "info"}}
	getFieldsByJsonReq := GetFieldsByJsonReq{MemberId: 2147700139, Fields: []string{"info", "telephone"}}

	aa, err := helloService.GetFieldsByJson(&getFieldsByJsonReq)

	if err != nil {
		fmt.Println(err)

	} else {
		//	fmt.Println(aa.JsonVal)
		mmm := make(map[string]interface{}, 0)

		err = json.Unmarshal([]byte(aa.JsonVal), &mmm)

		if err != nil {
			fmt.Println(err)
		}
		fmt.Printf("%+v\n", mmm["info"])
		mapInterface := mmm["info"].(string)

		mmm11 := make(map[string]interface{}, 0)
		json.Unmarshal([]byte(mapInterface), &mmm11)

		fmt.Println(mmm11["Channle"].(string))
		//		fmt.Println(GetInfoField(mmm, "channel"))

		//mm := Info{}

	}

	// aaa := Info{Version_name: "3.0.17.000.718.1741", Channle: "23,123"}
	// in := &Info11{Info1: aaa}
	// b, _ := json.Marshal(in)

	// fmt.Println(string(b))

	// jsonTemp := `{"info":{"channel":"%s"}}`

	// jsonValue := fmt.Sprintf(jsonTemp, "1,2,3")

	// fmt.Println(jsonValue)

	// updateFieldsByJsonReq := UpdateFieldsByJsonReq{MemberId: 2147700139, JsonVal: jsonValue}
	// bb, err := helloService.UpdateFieldsByJson(&updateFieldsByJsonReq)

	// if err != nil {
	// 	fmt.Println(err)

	// } else {
	// 	fmt.Println(bb)
	// }
}
