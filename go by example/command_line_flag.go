package main

import (
	"flag"
	"fmt"
)

func main() {
	wordPtr := flag.String("word", "foo", "a word")
	numPtr := flag.Int("num", 666, "a number")
	boolPtr := flag.Bool("bool", true, "a bool")

	var svar string
	flag.StringVar(&svar, "svar", "bar", "a string var")

	flag.Parse()

	fmt.Println("word:", *wordPtr)
	fmt.Println("num:", *numPtr)
	fmt.Println("bool:", *boolPtr)
	fmt.Println("svar:", svar)

	fmt.Println("tail:", flag.Args())
}
