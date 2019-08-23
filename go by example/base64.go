package main

import "fmt"
import bs64 "encoding/base64"

func main() {
	data := "this is a string for base64 encode : 8r78r43&((%&$&^)^"
	sEnc := bs64.StdEncoding.EncodeToString([]byte(data))
	fmt.Println(sEnc)

	sDec, _ := bs64.StdEncoding.DecodeString(sEnc)
	fmt.Println(string(sDec))
	fmt.Println()

	uEnc := bs64.URLEncoding.EncodeToString([]byte(data))
	fmt.Println(uEnc)
	uDec, _ := bs64.URLEncoding.DecodeString(uEnc)
	fmt.Println(string(uDec))
}
