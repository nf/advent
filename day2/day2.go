package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func main() {
	b, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}
	var pos []int
	for _, s := range strings.Split(string(b), ",") {
		n, err := strconv.Atoi(s)
		if err != nil {
			panic(err)
		}
		pos = append(pos, n)
	}

	for i := 0; i < 100; i++ {
		for j := 0; j < 100; j++ {
			input := append([]int(nil), pos...)
			input[1] = i
			input[2] = j
			n := compute(input)
			if n == 19690720 {
				fmt.Println(i, j, 100*i+j)
				return
			}
		}
	}
	fmt.Println("Failed!")
}

func compute(pos []int) int {
	pc := 0
	for {
		if pc >= len(pos) {
			panic(fmt.Sprintf("overrun pc=%d", pc))
		}
		switch pos[pc] {
		case 1: // add
			pos[pos[pc+3]] = pos[pos[pc+1]] + pos[pos[pc+2]]
			pc += 4
		case 2: // mul
			pos[pos[pc+3]] = pos[pos[pc+1]] * pos[pos[pc+2]]
			pc += 4
		case 99:
			return pos[0]
		default:
			panic(fmt.Sprintf("unknown opcode=%d", pos[pc]))
		}
	}
}
