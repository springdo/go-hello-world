package main

import "testing"

func TestGetMessage(t *testing.T) {
	m := getMessage()
	if len(m) != 20 {
		t.Errorf("Expected 20 but got %v", len(m))
	}

	if m != "Hi there or whatever" {
		t.Errorf("Expected `Hi there or whatever but got %v`", m)
	}
}
