package p34

import (
	"testing"
	"fmt"
)

func TestCompute(t *testing.T) {
	c := toConfigurationFromJSON("./pentadecathlon.json")

	a := c.toGrid()
	b := c.blankGrid()

	a1 := c.toGrid()
	b1 := c.blankGrid()

	a2 := c.toGrid()
	b2 := c.blankGrid()

	for i := 0; i < 10000; i++ {
		compute(a, b, c.Field.W, c.Field.H)
		computeC(a1, b1, c.Field.W, c.Field.H, 1)
		computeC(a2, b2, c.Field.W, c.Field.H, 4)

		s := display(b)
		s1 := display(b1)
		s2 := display(b2)

		fmt.Printf("%s\n%s\n%s\n", s, s1, s2)

		if s != s1 || s != s2 {
			t.Fail()
		}

		b, a = a, b
		b1, a1 = a1, b1
		b2, a2 = a2, b2
	}
}
