package testing

import (
	"fmt"
	"testing"
)

/*
test functions are required to start with "Test".
Convention is to name them after the method under test
*/
func TestRandomRangeValue(t *testing.T) {
	value := RandomRangeValue(0, 1)

	if value < 0 || value > 1 {
		t.Errorf("failed")
	}
}

// parameterized tests can be written with custom struct creation and assertions in the test function
func TestSomething(t *testing.T) {
	cases := []struct {
		name         string
		lower, upper int
	}{
		{"0-1", 0, 1},
	}
	for _, testcase := range cases {
		t.Log(fmt.Sprintf("executing '%s'", testcase.name))
	}
}
