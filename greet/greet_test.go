package main

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "DK")

	got := buffer.String()
	want := "Hello, DK"

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
