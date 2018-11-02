package sumOfMultiples

import "testing"

var tests = []struct {
	given    int
	expected int
}{
	{1, 0},
	{10, 23},
}

func TestMultiples(t *testing.T) {
	for _, test := range tests {
		actual := Multiples(test.given)
		if actual != test.expected {
			t.Errorf("Multiples(%d): expected %d, actual %d", test.given, test.expected, actual)
		}
	}
}
