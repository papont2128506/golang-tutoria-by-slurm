package main

import "fmt"

func main() {

	// float32 быстро накапливает ошибку
	var f float32 = 16777216 // 1 << 24
	fmt.Println(f == f+1)    // true !!! -

	var f64 float64 = 16777216 // 1 << 24
	fmt.Println(f64 == f64+1)  // false
}
