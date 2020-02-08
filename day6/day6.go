package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	nodes := make(map[string]string) // [name]parent

	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		// "A)B" - "B orbits around A"
		p := strings.SplitN(s.Text(), ")", 2)
		if _, ok := nodes[p[1]]; ok {
			log.Fatalf("%q already exists in node map", p[1])
		}
		nodes[p[1]] = p[0]
	}
	if err := s.Err(); err != nil {
		panic(err)
	}

	name := "YOU"
	seen := make(map[string]int)
	n := 0
	for name != "COM" {
		name = nodes[name]
		n++
		seen[name] = n
	}

	name = "SAN"
	n = 0
	for name != "COM" {
		name = nodes[name]
		n++
		if n2 := seen[name]; n2 > 0 {
			fmt.Println(n, n2, n+n2-2)
			break
		}
	}
}
