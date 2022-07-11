package main

import (
	"fmt"
)

var count int = 0

func main() {
	fmt.Println("Begin")

	// Explicit Declaration
	var tmp1 int = 0
	tmp1 = 10
	var tmp2 string = "hello"
	var tmp3 bool = true

	/*
		const tmp4 int = 0 //เปลี่ยนแปลงไม่ได้
		tmp4 = 10
	*/

	// Implicit Declaration
	// var tmp5 int = 0
	tmp5 := 0
	tmp6 := "codeGolang"

	fmt.Println(tmp1)
	fmt.Println(tmp2)
	fmt.Println(tmp3)
	// fmt.Println(tmp4)
	fmt.Println(tmp5)
	fmt.Println(tmp6)

	count++
	fmt.Println("count => ", count)
	count++
	fmt.Println("count => ", count)
	run()
}

func run() {
	count++
	fmt.Println("count => ", count)
}
