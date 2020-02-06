package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	b, err := ioutil.ReadFile("input")
	if err != nil {
		panic(err)
	}
	var mem []int
	for _, s := range strings.Split(string(b), ",") {
		n, err := strconv.Atoi(s)
		if err != nil {
			panic(err)
		}
		mem = append(mem, n)
	}
	fmt.Printf("result: %d\n", compute(mem))
}

type mode byte

const (
	positionMode  mode = 0
	immediateMode mode = 1
)

func compute(mem []int) int {
	pc := 0
	for {
		if pc >= len(mem) {
			panic(fmt.Sprintf("overrun pc=%d", pc))
		}

		// eg: 02
		opcode := mem[pc] % 100

		// eg: 11100
		modeFlags := (mem[pc] - opcode) / 100
		var modes [3]mode
		for i := range modes {
			modes[i] = mode(modeFlags % 10)
			modeFlags /= 10
		}

		load := func(addr, argN int) int {
			switch modes[argN] {
			case positionMode:
				return mem[mem[addr]]
			case immediateMode:
				return mem[addr]
			default:
				panic(fmt.Sprintf("invalid mode %d", modes[argN]))
			}

		}

		switch opcode {
		case 1: // add
			mem[mem[pc+3]] = load(pc+1, 0) + load(pc+2, 1)
			pc += 4
		case 2: // mul
			mem[mem[pc+3]] = load(pc+1, 0) * load(pc+2, 1)
			pc += 4
		case 3: // input
			fmt.Printf("input: ")
			var v int
			fmt.Scan(&v)
			mem[mem[pc+1]] = v
			pc += 2
		case 4: // output
			fmt.Printf("ouput: %d\n", load(pc+1, 0))
			pc += 2
		case 99: // quit
			return mem[0]
		default:
			panic(fmt.Sprintf("unknown opcode=%d", opcode))
		}
	}
}
