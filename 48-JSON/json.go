package main

import (
	"encoding/json"
	"fmt"
	"os"
)

//大写字母开头的字段才可以导出
type response1 struct {
	Page   int
	Fruits []string
}
type response2 struct {
	Page   int      `json:"kk"`
	Fruits []string `json:"fruits"`
}

func main() {
	//基本数据类型===>JSON
	bolB, _ := json.Marshal(true)
	fmt.Println(string(bolB))
	fmt.Println(bolB)
	intB, _ := json.Marshal(1)
	fmt.Println(string(intB))
	fltB, _ := json.Marshal(2.34)
	fmt.Println(string(fltB))
	strB, _ := json.Marshal("gopher")
	fmt.Println(string(strB))
	fmt.Println(strB)

	//切片===>数组
	//map===>对象
	slcD := []string{"apple", "peach", "pear"}
	slcB, _ := json.Marshal(slcD)
	fmt.Println(string(slcB))

	mapD := map[string]int{"apple": 5, "lettuce": 7}
	mapB, _ := json.Marshal(mapD)
	fmt.Println(string(mapB))

	//自定义类型
	res1D := &response1{
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"}}
	res1B, _ := json.Marshal(res1D)
	fmt.Println(string(res1B))

	res2D := &response2{
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"}}
	res2B, _ := json.Marshal(res2D)
	fmt.Println(string(res2B))

	//JSON  to GO
	byt := []byte(`{"num":6.13,"strs":["a","b"]}`)
	//解码以后存的地方
	var dat map[string]interface{}

	if err := json.Unmarshal(byt, &dat); err != nil {
		panic(err)
	}
	fmt.Println(dat)
	num := dat["num"].(float64)
	fmt.Println(num)
	strs := dat["strs"].([]interface{})
	str1 := strs[0].(string)
	fmt.Println(str1)

	str := `{"page":1,"fruits":["apple","peach"]}`
	res := response2{} //匹配一下类型
	json.Unmarshal([]byte(str), &res)
	fmt.Println(res.Fruits[0])

	enc := json.NewEncoder(os.Stdout)
	d := map[string]int{"apple": 5, "lettuce": 7}
	enc.Encode(d)

}
