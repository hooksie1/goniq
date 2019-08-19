package main

import "testing"

func TestRead(t *testing.T) {
	lines := readLines("test.txt")
	correct := []string{"line 1", "line 2"}

	if len(lines) != len(correct) {
		t.Errorf("Output was different lengths. Got %v, want: %v", len(lines), len(correct))
	}

	for i, v := range lines {
		if v != correct[i] {
			t.Errorf("Output was incorrect. Expected %v, but got %v", correct[i], lines[i])
		}
	}
}
