package integers

import (
	"fmt"
	"testing"
)

func TestAdder(t *testing.T) {
	sum := Add(2, 2)
	expected := 4

	if sum != expected {
		t.Errorf("Expected '%d' but got '%d'", expected, sum)
	}
}

// Example code that will run as part of documentation and tests
// Make sure to update the expected output if the example changes
func ExampleAdd() {
	sum := Add(1, 5)
	fmt.Println(sum)
	// Output: 6
}
