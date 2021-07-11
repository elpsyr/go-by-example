package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	d1 := []byte("hello\ngo\n")
	err := ioutil.WriteFile("D:\\Projects\\Go\\src\\go-by-example\\59-Writing-Files\\dat1", d1, 0664)
	check(err)

	f, err := os.Create("D:\\Projects\\Go\\src\\go-by-example\\59-Writing-Files\\dat2")
	check(err)

	defer f.Close()

	d2 := []byte{115, 111, 109, 101, 10}
	n2, err := f.Write(d2)
	check(err)
	fmt.Printf("wrote %d bytes\n", n2)

	n3, err := f.WriteString("writes\n")
	fmt.Printf("write %d bytes\n", n3)

	f.Sync()

	w := bufio.NewWriter(f)
	n4, err := w.WriteString("buffered\n")
	fmt.Printf("wrote %d bytes\n", n4)
	w.Flush()

}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
