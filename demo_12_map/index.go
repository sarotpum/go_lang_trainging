package main

import "fmt"

func main() {
	var numbers = map[string]int{"one": 1, "two": 2, "there": 3}
	fmt.Println("", numbers)
	fmt.Println("", numbers["two"])

	var colors = make(map[string]string)
	colors["red"] = "#f00"
	colors["green"] = "#0f0"
	colors["blue"] = "#00f"
	fmt.Println("", colors)
	fmt.Println("", colors["green"])

	var courses = make(map[string]map[string]int)

	courses["Android"] = make(map[string]int)
	courses["Android"]["price"] = 500
	courses["Android"]["price"] = 100
	courses["Android"]["code"] = 1234

	courses["ios"] = make(map[string]int)
	courses["ios"]["price"] = 200
	courses["ios"]["code"] = 444
	fmt.Println(courses)

	fmt.Println(courses["ios"])
	fmt.Println(courses["ios"]["code"])
}
