package main

import (
	"fmt"
	"math"
	"time"
)

func main() {
	fmt.Println(math.Exp2(10)) // 1024
	fmt.Println("---------")
	time.Sleep(1 * time.Second)
	fmt.Println("Finished")
}
