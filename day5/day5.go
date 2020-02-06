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
		n, err := strconv.Atoi(strings.TrimSpace(s))
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

		//log.Printf("pc=%d mem[pc]=%d opcode=%d modes=%d", pc, mem[pc], opcode, modes)
		pc++

		//log.Printf("next instructions %v", mem[pc:pc+5])

		argN := 0
		load := func() (v int) {
			defer func() {
				argN++
				pc++
			}()
			switch modes[argN] {
			case positionMode:
				return mem[mem[pc]]
			case immediateMode:
				return mem[pc]
			default:
				panic(fmt.Sprintf("invalid mode %d at pc %d", modes[argN], pc))
			}
		}
		stor := func(v int) {
			defer func() {
				argN++
				pc++
			}()
			mem[mem[pc]] = v
		}

		switch opcode {
		case 1: // add
			stor(load() + load())
		case 2: // mul
			stor(load() * load())
		case 3: // input
			fmt.Printf("input: ")
			var v int
			fmt.Scan(&v)
			stor(v)
		case 4: // output
			fmt.Printf("output: %d\n", load())
		case 5: // jump-if-true
			if v, dst := load(), load(); v != 0 {
				pc = dst
			}
		case 6: // jump-if-false
			if v, dst := load(), load(); v == 0 {
				pc = dst
			}
		case 7: // less than
			if load() < load() {
				stor(1)
			} else {
				stor(0)
			}
		case 8: // equals
			if load() == load() {
				stor(1)
			} else {
				stor(0)
			}
		case 99: // quit
			return mem[0]
		default:
			panic(fmt.Sprintf("unknown opcode=%d", opcode))
		}
	}
}
