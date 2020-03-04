package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Hello")
	want := "Hello,wor1d"

	if got != want {
		t.Errorf("got '%q' want '%q'",got,want)
	}

}
