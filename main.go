package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	max := 20
	stepper := 0
	var Root Candy
	cn := 0
	startAgain := false
	Root.candyNumber = &cn
	Root.CreateColumn(max, &stepper)
	Root.PrintCandyCol()
	Root.DeleteDuplicants(&startAgain, &Root)
	Root.PrintCandyCol()
	fmt.Println("Hello!")
}

type Candy struct {
	color          int
	next           *Candy
	deleteNextTime bool
	candyNumber    *int
}

func (c *Candy) CreateColumn(max int, stepper *int) {
	c.color = rand.Intn(19)
	*stepper += 1
	*c.candyNumber++
	if *stepper < max {
		var actCandy Candy
		c.next = &actCandy
		c.next.candyNumber = c.candyNumber
		c.next.CreateColumn(max, stepper)
	}
}

func (c *Candy) PrintCandyCol() {
	fmt.Print(c.color)
	if c.next != nil {
		fmt.Print(",")
		c.next.PrintCandyCol()
	} else {
		fmt.Printf("\n")
	}
}

func (c *Candy) DeleteDuplicants(startAgain *bool, firstElement *Candy) {
	if c.next != nil {
		if c.next.color == c.color {
			c.DeleteNext()
			c.deleteNextTime = true
			*startAgain = true
			c.DeleteDuplicants(startAgain, firstElement)
		} else if c.next.deleteNextTime {
			c.DeleteNext()
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
