package main

import "fmt"

func main() {
	var m1 map[int32]bool
	var m2 map[string]string

	m3 := make(map[int]int)

	ages := map[string]int{
		"Андрей":    30,
		"Анастасия": 25,
	}

	age := ages["Андрей"]
	ages["Наталья"] = 31

	fmt.Println(ages)
	fmt.Println(m1, m2, m3, age)
}
