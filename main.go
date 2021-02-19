package main

import (
	"fmt"
	"math/rand"
)

func main() {
	max := 100
	stepper := 0
	var Root Candy
	cn := 0
	Root.candyNumber = &cn
	Root.CreateColumn(max, &stepper)
	fmt.Println("Hello!")
}

type Candy struct {
	color          int
	next           *Candy
	deleteNextTime bool
	candyNumber    *int
}

func (c *Candy) CreateColumn(max int, stepper *int) {
	c.color = rand.Intn(10)
	*stepper += 1
	*c.candyNumber++
	if *stepper < max {
		var actCandy Candy
		c.next = &actCandy
		c.next.candyNumber = c.candyNumber
		c.next.CreateColumn(max, stepper)
	}
}

func (c *Candy) DeleteDuplicants(startAgain *bool, firstElement Candy) {
	if c.next != nil {
		if c.next.color == c.color || c.next.deleteNextTime {
			c.DeleteNext()
			c.deleteNextTime = true
			*startAgain = true
			c.DeleteDuplicants(startAgain, firstElement)
		} else {
			c.next.DeleteDuplicants(startAgain, firstElement)
		}
	} else if *startAgain {
		*startAgain = false
		firstElement.DeleteDuplicants(startAgain, firstElement)
	}
}

func (c *Candy) DeleteNext() {
	*c.candyNumber--
	c.next = c.next.next
}
