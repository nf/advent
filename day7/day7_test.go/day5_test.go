package main

import "testing"

func TestMachine(t *testing.T) {
	const prog = `3,21,1008,21,8,20,1005,20,22,107,8,21,20,1006,20,31,1106,0,36,98,0,0,1002,21,125,20,4,20,1105,1,46,104,999,1105,1,46,1101,1000,1,20,4,20,1105,1,46,98,99`

	cases := []struct {
		in   int
		want int
	}{
		{7, 999},
		{8, 1000},
		{9, 1001},
	}

	for _, c := range cases {
		m, err := newMachine(prog)
		if err != nil {
			t.Error(err)
		}
		m.input = func() int { return c.in }
		var got int
		m.output = func(v int) { got = v }
		m.run()
		if got != c.want {
			t.Errorf("input %d returned %d, want %d", c.in, got, c.want)
		}

	}
}
