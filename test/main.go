package main

import (
	"encoding/json"
	"fmt"
)

//type Test struct {
//	Name String
//}

func main() {
	byt := []byte(`{"num":6.13,"strs":["a","b"]}`)
	var dat map[string]interface{}
	if err := json.Unmarshal(byt, &dat); err != nil {
		panic(err)
	}

	println(dat)
	fmt.Println(dat)

	p, err := json.MarshalIndent(dat, "", "  ")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(p))

	//	fmt.Println(dat)
	//
	//	num := dat["num"].(float64)
	//	fmt.Println(num)
	//
	//	strs := dat["strs"].([]interface{})
	//	fmt.Println(strs)
	//	fmt.Println(reflect.TypeOf(strs))
	//
	//	str1 := strs[0]
	//	fmt.Println(reflect.TypeOf(str1))
	//
	//	str2 := strs[0].(string)
	//	fmt.Println(reflect.TypeOf(str2))
	//
}
