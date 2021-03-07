package main

import "testing"

// t is a test handler
func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 16 {
		// t.Error("Expect deck length of 20 but got ", len(d))
		t.Errorf("Expect deck length of 20 but got %v", len(d))
	}
}
b