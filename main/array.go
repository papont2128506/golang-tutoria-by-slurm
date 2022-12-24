package main

import "fmt"

func main() {
	var x1 [5]int // массив 5 целых чисел
	var x2 [0]int
	var x3 [5]string

	var arr [3]int = [3]int{1, 2, 3}
	var arr2 = [3]int{1, 2, 3}
	arr3 := [3]int{1, 2, 3}

	//arr3 = [2]int{1, 2}

	arr5 := [...]int32{1, 2, 3, 4, 5}

	s := len(arr5)

	fmt.Println(arr, arr2, arr3, arr5, x1, x2, x3, s)
}
