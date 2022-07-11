package main

//go mod init demo14
import (
	"demo14/employee"
)

func main() {
	e := employee.Init(
		"Sam",
		"Adolf",
		30,
		20)

	e.LeavesRemaining()

	e = employee.Init(
		"Lek",
		"Adolf",
		30,
		20)

	e.LeavesRemaining()
}
