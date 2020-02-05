package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Move struct {
	dir   byte
	count int
}

func (m Move) String() string {
	return fmt.Sprintf("%c%d", m.dir, m.count)
}

func main() {
	var paths [][]Move
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		var path []Move
		for _, s := range strings.Split(s.Text(), ",") {
			n, err := strconv.Atoi(s[1:])
			if err != nil {
				panic(err)
			}
			path = append(path, Move{dir: s[0], count: n})
		}
		paths = append(paths, path)
	}
	if len(paths) != 2 {
		log.Fatal("need two wires")
	}

	const size = 40000
	type cell [2]uint32
	grid := make([][]cell, size)
	for y := range grid {
		grid[y] = make([]cell, size)
	}
	origX, origY := size/2, size/2

	for wire, path := range paths {
		step := uint32(0)
		x, y := origX, origY
		inc := func() {
			step++
			if grid[y][x][wire] != 0 {
				return
			}
			if step == 0 {
				fmt.Printf("overflow x=%d y=%d\n", x, y)
			}
			grid[y][x][wire] = step
		}
		for _, m := range path {
			switch m.dir {
			case 'U':
				for i := 0; i < m.count; i++ {
					y++
					inc()
				}
			case 'D':
				for i := 0; i < m.count; i++ {
					y--
					inc()
				}
			case 'L':
				for i := 0; i < m.count; i++ {
					x--
					inc()
				}
			case 'R':
				for i := 0; i < m.count; i++ {
					x++
					inc()
				}
			}
		}
	}

	var lowX, lowY, lowSteps int
	for y := range grid {
		for x := range grid[y] {
			c := grid[y][x]
			if c[0] > 0 && c[1] > 0 {
				steps := c[0] + c[1]
				if lowSteps == 0 || lowSteps > int(steps) {
					lowX, lowY, lowSteps = x, y, int(steps)
				}
			}
		}
	}
	fmt.Printf("x=%d y=%d steps=%d\n", lowX, lowY, lowSteps)
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
