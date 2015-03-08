package main

import "fmt"

type MyStruct struct {
	x string
	i int
	f float64
	m map[string] int64
}

func main() {
	x := "abc"
	i := 3
	fmt.Println(i)
	fmt.Println(x)

	ms := &MyStruct{
		x: "cba",
		i: 10,
		f: 11.10335,
		m: make(map[string]int64),
	}
	ms.m["hello"] = int64(11)
	ms.m["world"] = int64(22)
	fmt.Println(ms)
}
