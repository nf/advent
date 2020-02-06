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
	m, err := newMachine(string(b))
	if err != nil {
		panic(err)
	}
	fmt.Printf("result: %d\n", m.run())
}

type mode byte

const (
	positionMode  mode = 0
	immediateMode mode = 1
)

type machine struct {
	pc  int
	mem []int

	// Reset after each call to opcode.
	argN  int
	modes [3]mode

	input  func() int
	output func(int)
}

func newMachine(prog string) (*machine, error) {
	m := machine{
		input: func() int {
			fmt.Printf("input: ")
			var v int
			fmt.Scan(&v)
			return v
		},
		output: func(v int) {
			fmt.Printf("output: %d\n", v)
		},
	}
	for _, s := range strings.Split(prog, ",") {
		n, err := strconv.Atoi(strings.TrimSpace(s))
		if err != nil {
			return nil, err
		}
		m.mem = append(m.mem, n)
	}
	return &m, nil
}

func (m *machine) opcode() int {
	// eg: 02
	opcode := m.mem[m.pc] % 100

	// eg: 11100
	modeFlags := (m.mem[m.pc] - opcode) / 100
	for i := range m.modes {
		m.modes[i] = mode(modeFlags % 10)
		modeFlags /= 10
	}
	m.argN = 0

	//log.Printf("pc=%d mem[pc]=%d opcode=%d modes=%d", pc, mem[pc], opcode, modes)
	m.pc++
	//log.Printf("next instructions %v", mem[pc:pc+5])
	return opcode
}

func (m *machine) load() int {
	defer func() {
		m.argN++
		m.pc++
	}()

	switch m.modes[m.argN] {
	case positionMode:
		return m.mem[m.mem[m.pc]]
	case immediateMode:
		return m.mem[m.pc]
	default:
		panic(fmt.Sprintf("invalid mode %d at pc %d", m.modes[m.argN], m.pc))
	}
}

func (m *machine) stor(v int) {
	defer func() {
		m.argN++
		m.pc++
	}()
	m.mem[m.mem[m.pc]] = v
}

func (m *machine) run() int {
	for {
		switch opcode := m.opcode(); opcode {
		case 1: // add
			m.stor(m.load() + m.load())
		case 2: // mul
			m.stor(m.load() * m.load())
		case 3: // input
			m.stor(m.input())
		case 4: // output
			m.output(m.load())
		case 5: // jump-if-true
			if v, dst := m.load(), m.load(); v != 0 {
				m.pc = dst
			}
		case 6: // jump-if-false
			if v, dst := m.load(), m.load(); v == 0 {
				m.pc = dst
			}
		case 7: // less than
			if m.load() < m.load() {
				m.stor(1)
			} else {
				m.stor(0)
			}
		case 8: // equals
			if m.load() == m.load() {
				m.stor(1)
			} else {
				m.stor(0)
			}
		case 99: // quit
			return m.mem[0]
		default:
			panic(fmt.Sprintf("unknown opcode=%d", opcode))
		}
	}
}
