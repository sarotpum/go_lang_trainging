package main

import (
	"fmt"
	"time"
)

func run1() {
	for i := 0; i < 100; i++ {
		fmt.Println("Run1 something")
	}
}

func run2() {
	for i := 0; i < 100; i++ {
		fmt.Println("Run2 something")
	}
}

func main() {
	//ใส่คำสั่ง go จะเป็นการ run แบบ
	go run1()
	go run2()

	time.Sleep(5 * time.Second)
}
