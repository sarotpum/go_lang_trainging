package main

//go mod init demo14
import (
	"demo14/employee"
)

func main() {
	e := employee.New(
		"Sam",
		"Adolf",
		30,
		20)

	e.LeavesRemaining()
}
