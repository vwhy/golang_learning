package main

import (
	"fmt"
	"time"
)

func main() {
	i := 2
	fmt.Print("write ", i, " as ")
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's weekend !")
	default:
		fmt.Println("Oh, It's time to work")
	}

	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("It's before noon")
	default:
		fmt.Println("It's afternoon")
	}

	whatstype := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("It's a bool")
		case int:
			fmt.Println("It's a int")
		default:
			fmt.Printf("It's a unknow type: %T\n", t)
		}
	}
	whatstype(true)
	whatstype(666)
	whatstype("whattf")

}
