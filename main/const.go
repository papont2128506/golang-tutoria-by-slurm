package main

import "fmt"

func main() {
	//const earthRadiusInMeter int = 6371000
	const earthRadiusInMeter = 6371000

	_ = earthRadiusInMeter
	//
	//const (
	//	earthRadiusInMeter int = 6371000
	//	year                   = 2022
	//)

	x := 1
	ptr := &x
	*ptr = 5

	fmt.Println(x)
}
