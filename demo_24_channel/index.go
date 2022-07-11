package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int, 1) // 1 เป็นการจอง ถ้าไม่ใส่ จะตาย
	ch <- 1                 //send

	// step1
	// fmt.Println("step1")
	// ch <- 2 // ถ้าไม่เปลี่ยน 1 เป็น 2 จะตายเพราะไม่ได้จอง
	// fmt.Println("step2")
	// // msg := <-ch จะมองแบบนี้ก็ได้
	// fmt.Println(<-ch)

	// step2
	fmt.Println("step1")
	fmt.Println(<-ch)

	ch <- 2 //send
	fmt.Println("step2")
	fmt.Println(<-ch)

	time.Sleep(1 * time.Second)
}
