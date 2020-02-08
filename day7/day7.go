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
	prog, err := parseProg(string(b))
	if err != nil {
		panic(err)
	}

	var highSig int
	for _, phase := range permute(5) {
		sig := compute(prog, phase)
		if sig > highSig {
			highSig = sig
		}
	}
	fmt.Println(highSig)
}

func permute(n int) (out [][]int) {
	if n == 1 {
		return [][]int{{0}}
	}
	base := permute(n - 1)
	for i := range base {
		b := base[i]
		for i := 0; i <= len(b); i++ {
			p := append([]int(nil), b[:i]...)
			p = append(p, n-1)
			p = append(p, b[i:]...)
			out = append(out, p)
		}
	}
	return
}

func compute(prog, phase []int) int {
	in := make(chan int, 2)
	first := in
	var out chan int
	for _, ps := range phase {
		in <- ps
		out = make(chan int, 2)
		go newMachine(prog, in, out).run()
		in = out
	}
	first <- 0 // start computation
	return <-out
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

func newMachine(mem []int, in <-chan int, out chan<- int) *machine {
	return &machine{
		mem:    append([]int(nil), mem...),
		input:  func() int { return <-in },
		output: func(v int) { out <- v },
	}
}

func parseProg(s string) ([]int, error) {
	var mem []int
	for _, s := range strings.Split(s, ",") {
		n, err := strconv.Atoi(strings.TrimSpace(s))
		if err != nil {
			return nil, err
		}
		mem = append(mem, n)
	}
	return mem, nil
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
