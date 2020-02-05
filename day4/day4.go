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
	two := false
	runLen := 1
	for i := 1; i < 6; i++ {
		if c[i] == c[i-1] {
			runLen++
		} else {
			if runLen == 2 {
				two = true
			}
			runLen = 1
		}
		if c[i] < c[i-1] {
			return false
		}
	}
	if runLen == 2 {
		return true
	}
	return two
}

func main() {
	c := code{2, 3, 1, 8, 3, 2}
	n := 0
	for i := 0; i < 767346-231832+1; i++ {
		if c.valid() {
			n++
		}
		c.inc()
	}
	fmt.Println(n)
}
