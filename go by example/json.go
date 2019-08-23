package main

import "encoding/json"
import "fmt"
import "os"

var p = fmt.Println

type response1 struct {
	Page   int
	Fruits []string
}

type response2 struct {
	Page   int      `json:"page"`
	Fruits []string `json:"fruits"`
}

func main() {
	bolB, _ := json.Marshal(true)
	p(string(bolB))

	intB, _ := json.Marshal(666)
	p(string(intB))

	fltB, _ := json.Marshal(6.66)
	p(string(fltB))

	strB, _ := json.Marshal("what")
	p(string(strB))

	slcD, _ := json.Marshal([]string{"apple", "peach", "pear"})
	p(string(slcD))

	mapD, _ := json.Marshal(map[string]int{"apple": 5, "banana": 6})
	p(string(mapD))

	res1D := &response1{
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"}}
	res1B, _ := json.Marshal(res1D)
	p(string(res1B))

	res2D := &response2{
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"}}
	res2B, _ := json.Marshal(res2D)
	p(string(res2B))

	byt := []byte(`{"num":6.13,"strs":["a","b"]}`)

	var dat map[string]interface{}

	if err := json.Unmarshal(byt, &dat); err != nil {
		panic(err)
	}
	p(dat)

	num := dat["num"].(float64)
	p(num)
	strs := dat["strs"].([]interface{})
	strs1 := strs[0].(string)
	p(strs1)
	str := []byte(`{"page":1,"fruits":["apple","peach"]}`)
	res := &response2{}
	json.Unmarshal(str, &res)
	p(res)
	p(res.Fruits[1])

	enc := json.NewEncoder(os.Stdout)
	d := map[string]int{"apple": 5, "lettue": 7}
	enc.Encode(d)
}
