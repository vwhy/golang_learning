package main

import "fmt"
import "io/ioutil"
import (
	"bufio"
	"io"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	dat, err := ioutil.ReadFile("./dat")
	check(err)
	fmt.Println(string(dat))

	f, err := os.Open("./dat")
	check(err)
	defer f.Close()

	b1 := make([]byte, 13)
	n1, err := f.Read(b1)
	check(err)
	fmt.Printf("%d byte: %s\n", n1, string(b1))

	o2, err := f.Seek(9, 0)
	check(err)
	b2 := make([]byte, 4)
	n2, err := f.Read(b2)
	check(err)
	fmt.Printf("%d bytes @ %d: %s\n", n2, o2, string(b2))

	o3, err := f.Seek(9, 0)
	check(err)
	b3 := make([]byte, 6)
	n3, err := io.ReadAtLeast(f, b3, 4)
	check(err)
	fmt.Printf("%d bytes @ %d: %s\n", n3, o3, string(b3))

	_, err = f.Seek(0, 0)
	check(err)

	r4 := bufio.NewReader(f)
	b4, err := r4.Peek(8)
	check(err)
	fmt.Printf("8 byte:%s\n", string(b4))
}
