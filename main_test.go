package main

import "testing"

func TestAdd(t *testing.T) {
	got := Add(2, 5)
	want := 7

	if got != want {
		t.Errorf("Add(2, 5) = %d; want %d", got, want)
	}
}
