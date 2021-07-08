package main

import (
	"bytes"
	"fmt"
	"regexp"
)

func main() {
	match, _ := regexp.MatchString("p([a-z]+)ch", "peach")
	fmt.Println(match)

	r, _ := regexp.Compile("p([a-z]+)ch")

	fmt.Println(r.MatchString("peach"))       //true
	fmt.Println(r.FindString("pea1ch punch")) //punch

	fmt.Println(r.FindStringIndex("peach punch")) //[0 5]

	fmt.Println(r.FindStringSubmatch("peach punch")) //[peach ea]

	fmt.Println(r.FindStringSubmatchIndex("peach punch")) //[0 5 1 3]

	//匹配个数限制
	fmt.Println(r.FindAllString("peach punch pinch", -1))
	fmt.Println(r.FindAllStringSubmatchIndex("peach punch pinch", -1))

	fmt.Println(r.FindAllString("peach punch pinch", 2))

	//String ==> Byte
	fmt.Println(r.Match([]byte("peach")))

	//更安全
	r = regexp.MustCompile("p([a-z]+)ch")
	fmt.Println(r)

	//替换字符串
	fmt.Println(r.ReplaceAllString("I like peach", "<fruit>"))

	in := []byte("a peach")
	out := r.ReplaceAllFunc(in, bytes.ToUpper)
	fmt.Println(string(out))

}
