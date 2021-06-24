package main

import (
	"log"
	assignmentAPI "scg-assignment/api/assignment"
	"scg-assignment/internal/cashier"
	"scg-assignment/internal/search"
)

func main() {

	// init money for cashier
	money := cashier.MoneyMap{
		1000: 10,
		500:  20,
		100:  15,
		50:   20,
		20:   30,
		10:   20,
		5:    20,
		1:    20,
		0.25: 50,
	}

	c := cashier.New(money)

	s := search.New()

	server := assignmentAPI.New(s, c)

	log.Fatal(server.Start(":8080"))
}
