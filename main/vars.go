package main

import "fmt"

func main() {
	var v1 int //v1 == 0
	fmt.Println(v1)

	var v2 int = 100
	fmt.Println(v2)

	v3 := 5
	fmt.Println(v3)

	v4 := 5
	//v4 := 5
	v4 = 10
	fmt.Println(v4)

	var v5, v6 = 1, 2
	v5, v6 = 3, 4

	var v7 int
	v5, v6, v7 = 1, 2, 3
	v5, v6 = v6, v5

	var (
		v01 = 1
		v02 = "string"
	)

	_ = v7
	_ = v01
	_ = v02

}
