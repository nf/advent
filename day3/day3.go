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
	grid := make([][]byte, size)
	for y := range grid {
		grid[y] = make([]byte, size)
	}
	origX, origY := size/2, size/2

	for wire, path := range paths {
		wire := uint(wire)
		x, y := origX, origY
		for _, m := range path {
			switch m.dir {
			case 'U':
				for i := 0; i < m.count; i++ {
					y++
					grid[y][x] |= 1 << wire
				}
			case 'D':
				for i := 0; i < m.count; i++ {
					y--
					grid[y][x] |= 1 << wire
				}
			case 'L':
				for i := 0; i < m.count; i++ {
					x--
					grid[y][x] |= 1 << wire
				}
			case 'R':
				for i := 0; i < m.count; i++ {
					x++
					grid[y][x] |= 1 << wire
				}
			}
		}
	}

	var lowX, lowY, lowDist int
	for y := range grid {
		for x := range grid[y] {
			if grid[y][x] == 3 {
				dist := abs(origX-x) + abs(origY-y)
				if lowDist == 0 || lowDist > dist {
					lowX, lowY, lowDist = x, y, dist
				}
			}
		}
	}
	fmt.Printf("x=%d y=%d dist=%d\n", lowX, lowY, lowDist)
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
