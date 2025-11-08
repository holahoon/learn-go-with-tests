package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("David")
	want := "hello David"

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
