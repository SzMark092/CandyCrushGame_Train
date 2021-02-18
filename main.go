package main

func main() {

}

type Candy struct {
	color          int
	next           *Candy
	deleteNextTime bool
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
	c.next = c.next.next
}
