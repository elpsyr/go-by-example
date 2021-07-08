package main

import (
	"encoding/json"
	"fmt"
)

//大写字母开头的字段才可以导出
type response1 struct {
	Page   int
	Fruits []string
}
type response2 struct {
	Page   int      `json:"page"`
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

}
