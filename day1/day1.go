package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var inputs []int
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		n, err := strconv.Atoi(s.Text())
		if err != nil {
			panic(err)
		}
		inputs = append(inputs, n)
	}

	sum := 0
	for _, n := range inputs {
		fuel := n/3 - 2
		for fuel > 0 {
			sum += fuel
			fuel = fuel/3 - 2
		}
	}
	fmt.Println(sum)
}
