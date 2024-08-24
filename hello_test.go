package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("DK")
	want := "Hello, DK"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
