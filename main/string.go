package main

import "fmt"

func main() {
	var myString string // ""
	fmt.Println(myString)

	var hello string = "\tHello\n"
	var title string = `\tHello\n`
	fmt.Println(hello)
	fmt.Println(title)

	//word := "sssss"

	var b byte = 'c' //99 - байтовой предствление символа
	println(b)

	l := len(hello) // количество байт, а не длина строки
	println(l)

	//str := "12345"
	//str[from:to]
	//newStr := str[2:4] //"34"
	//newStr := str[:4]  //"1234"
	//newStr := str[2:]  //"345"

}
