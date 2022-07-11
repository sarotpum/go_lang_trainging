package main

import "fmt"

func main() {
	someValue := 10
	if someValue == 10 {
		fmt.Println("== 10")
	} else {
		fmt.Println("!= 10")
	}

	if someValue == 10 || someValue < 20 {
		fmt.Println("someValue > 10 || someValue < 2 ")
	} else {
		fmt.Println("Not someValue > 10 || someValue < 2 ")
	}

	if result := doSomething(); result == "OK" {
		fmt.Println("OK")
	} else {
		fmt.Println("NO OK")
	}

	fnSwitchCase()
}

func doSomething() string {
	// do something
	return "OK"
}

func fnSwitchCase() {
	index := 3
	switch index {
	case 0:
		fmt.Println("0")
	case 1:
		fmt.Println("1")
	case 2:
		fmt.Println("2")
	default:
		fmt.Println("something else")
	}
}
