package employee //no package main

import "fmt"

type employee struct {
	FirstName   string
	LastName    string
	TotalLeaves int
	LeavesTaken int
}

func New(
	firstname string,
	lastname string,
	totalLeaves int,
	leavesTaken int) employee {
	e := employee{
		FirstName:   firstname,
		LastName:    lastname,
		TotalLeaves: totalLeaves,
		LeavesTaken: leavesTaken,
	}
	return e
}

func (e employee) LeavesRemaining() {
	sum := e.TotalLeaves - e.LeavesTaken
	fmt.Printf("%s %s has %d leaves remaining\n", e.FirstName, e.LastName, sum)
}
