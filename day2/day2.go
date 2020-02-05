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

	pos[1] = 12
	pos[2] = 2

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
			fmt.Println(pos[0])
			return
		default:
			panic(fmt.Sprintf("unknown opcode=%d", pos[pc]))
		}
	}
}
