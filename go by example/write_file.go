package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	d1 := []byte("hello\ngo\n")
	err := ioutil.WriteFile("./dat1", d1, 0644)
	check(err)
	f, err := os.Create("./dat2")
	check(err)
	defer f.Close()
	d2 := []byte{115, 111, 109, 101, 10}
	n2, err := f.Write(d2)
	check(err)
	fmt.Printf("%d bytes wrote\n", n2)

	n3, err := f.WriteString("writes\n")
	fmt.Printf("%d bytes wrote\n", n3)

	f.Sync()

	w := bufio.NewWriter(f)
	n4, err := w.WriteString("bufferd\n")
	fmt.Printf("%d bytes wrote\n", n4)

	w.Flush()
}
