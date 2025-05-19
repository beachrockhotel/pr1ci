package main

import "testing"

func TestAdd(t *testing.T) {
	got := Add(2, 5)
	want := 7
	if got != want {
		t.Errorf("Add(2, 5) = %d; want %d", got, want)
	}
}

func TestMultiply(t *testing.T) {
	got := Multiply(3, 4)
	want := 12
	if got != want {
		t.Errorf("Multiply(3, 4) = %d; want %d", got, want)
	}
}
