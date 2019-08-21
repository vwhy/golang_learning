package main

import "fmt"
import "os"

func main() {
	f := createfile("/tmp/defer.txt")
	defer closefile(f)
	writefile(f)
}

func createfile(p string) *os.File {
	fmt.Println("creating")
	f, err := os.Create(p)
	if err != nil {
		panic(err)
	}
	return f
}

func writefile(f *os.File) {
	fmt.Println("writing")
	fmt.Fprintln(f, "data")
}

func closefile(f *os.File) {
	fmt.Println("closing")
	err := f.Close()

	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}
