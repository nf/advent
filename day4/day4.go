package main

import "fmt"

type code [6]byte

func (c *code) inc() {
	for i := 5; i >= 0; i-- {
		if c[i] == 9 {
			c[i] = 0
			continue
		}
		c[i]++
		break
	}
}

func (c code) valid() bool {
	prev := c[0]
	two := false
	for i := 1; i < 6; i++ {
		if c[i] == prev {
			two = true
		}
		if c[i] < prev {
			return false
		}
		prev = c[i]
	}
	return two
}

func main() {
	c := code{2, 3, 1, 8, 3, 2}
	n := 0
	for i := 0; i < 767346-231832; i++ {
		c.inc()
		if c.valid() {
			fmt.Println(c)
			n++
		}
	}
	fmt.Println(n)
}
