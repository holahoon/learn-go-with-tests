package iteration

import "strings"

const repeatCount = 5

// func Repeat(character string) string {
// 	var repeated string
// 	for i := 0; i < repeatCount; i++ {
// 		repeated += character
// 	}
// 	return repeated
// }

// Strings in Go are immutable, meaning every concatenation, involves copying memory to accommodate the new string.
// This impacts performance, particularly during heavy string concatenation in loops.

// The standard library provides the strigs.Builder type which minimizes memory copying.
// It implements a `WriteString`` method which we can use to efficiently build our string.

// You can test the performance improvement by running the benchmark test before and after this change.
// go test <directory> -bench=.

func Repeat(character string) string {
	var repeated strings.Builder
	for i := 0; i < repeatCount; i++ {
		repeated.WriteString(character)
	}
	return repeated.String()
}
