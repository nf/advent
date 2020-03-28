package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

const (
	width, height       = 25, 6
	black, white, trans = '0', '1', '2'
)

type layer []byte

func (l layer) count(b byte) (n int) {
	for _, c := range l {
		if c == b {
			n++
		}
	}
	return
}

func (l layer) apply(l2 layer) {
	if len(l) != len(l2) {
		panic("length mismatch")
	}
	for i := range l2 {
		switch l2[i] {
		case black, white:
			if l[i] == 0 {
				l[i] = l2[i]
			}
		case trans:
			// do nothing?
		}
	}
}

func (l layer) render() string {
	var s strings.Builder
	for i := range l {
		if i%width == 0 && i > 0 {
			s.WriteByte('\n')
		}
		switch l[i] {
		case white:
			s.WriteByte('@')
		default:
			s.WriteByte(' ')
		}
	}
	return s.String()
}

func main() {
	input, _ := ioutil.ReadAll(os.Stdin)

	var image []layer
	var current layer
	for _, digit := range input {
		if len(current) == width*height {
			image = append(image, current)
			current = nil
		}
		current = append(current, digit)
	}
	if len(current) > 0 {
		image = append(image, current)
	}

	var fewest layer
	for _, l := range image {
		if fewest == nil {
			fewest = l
			continue
		}
		if fewest.count(black) > l.count(black) {
			fewest = l
		}
	}
	fmt.Println(fewest.count(white) * fewest.count(trans))

	output := make(layer, width*height)
	for _, l := range image {
		output.apply(l)
	}
	fmt.Println(output.render())
}
