package main

import "testing"

func TestCode(t *testing.T) {
	cases := []struct {
		in   code
		want bool
	}{
		{code{1, 1, 2, 2, 3, 3}, true},
		{code{1, 2, 3, 4, 4, 4}, false},
		{code{1, 1, 1, 1, 2, 2}, true},
		{code{1, 1, 2, 2, 2, 2}, true},
		{code{2, 2, 2, 2, 2, 2}, false},
		{code{2, 2, 3, 3, 4, 4}, true},
	}

	for _, c := range cases {
		got := c.in.valid()
		if got != c.want {
			t.Errorf("%v.valid() returned %v, want %v", c.in, got, c.want)
		}
	}
}
