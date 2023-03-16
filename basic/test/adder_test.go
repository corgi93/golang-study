package integers

import "testing"

func AddTesting(t *testing.T) {
	sum := Add(2, 2)
	expected := 5

	if sum != expected {
		t.Errorf("expected '%d' but got '%d'", expected, sum)
	}
}

func Add(a, b int) int {
	return a + b
}
