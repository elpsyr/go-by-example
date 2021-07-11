package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	//最简单的读取
	dat, err := ioutil.ReadFile("D:\\Projects\\Go\\src\\go-by-example\\57-Base64-Encoding\\data")
	check(err)
	fmt.Println(string(dat))

	//按字节读取
	f, err := os.Open("D:\\Projects\\Go\\src\\go-by-example\\57-Base64-Encoding\\data")
	check(err)
	defer f.Close()

	b1 := make([]byte, 5)
	n1, err := f.Read(b1)
	check(err)
	fmt.Printf("%d bytes: %s\n", n1, string(b1[:n1]))
	//定位读取
	o2, err := f.Seek(6, 0)
	check(err)
	b2 := make([]byte, 2)
	n2, err := f.Read(b2)
	check(err)
	fmt.Printf("%d bytes @ %d: ", n2, o2)
	fmt.Printf("%v\n", string(b2[:n2]))

	//更好的定位
	o3, err := f.Seek(6, 0)
	check(err)
	b3 := make([]byte, 2)
	n3, err := io.ReadAtLeast(f, b3, 2)
	check(err)
	fmt.Printf("%d bytes @ %d: %s\n ", n3, o3, string(b3))

	//倒带
	_, err = f.Seek(0, 0)
	check(err)

	r4 := bufio.NewReader(f)
	b4, err := r4.Peek(5)
	check(err)
	fmt.Printf("5 bytes: %s\n", string(b4))

}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
