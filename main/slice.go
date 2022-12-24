package main

import "fmt"

func main() {
	var list []int64
	l := len(list) //0
	c := cap(list) //0
	fmt.Println(list, l, c)

	list = []int64{1, 2, 3, 4, 5} //[1, 2, 3, 4, 5]
	l = len(list)                 //5
	c = cap(list)                 //5
	fmt.Println(list, l, c)

	list = make([]int64, 0, 5) //[]
	l = len(list)
	c = cap(list)
	fmt.Println(list, l, c)

	list = make([]int64, 5, 5) //[0 0 0 0 0]
	l = len(list)
	c = cap(list)
	fmt.Println(list, l, c)

	list = append(list, 5) //[0 0 0 0 0 5]
	l = len(list)
	c = cap(list)
	fmt.Println(list, l, c)

	list = make([]int64, 0, 3) // []           len: 0, cap: 3
	list = append(list, 1)     // [1]          len: 1, cap: 3
	list = append(list, 2)     // [1, 2]       len: 2, cap: 3
	list = append(list, 3)     // [1, 2, 3]    len: 3, cap: 3
	list = append(list, 4)     // [1, 2, 3, 4] len: 4, cap: 6

}
