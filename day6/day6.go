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

	total := 0
	for name := range nodes {
		total += orbits(nodes, name)
	}
	fmt.Println(total)
}

func orbits(nodes map[string]string, name string) (n int) {
	for name != "COM" {
		name = nodes[name]
		n++
	}
	return
}
