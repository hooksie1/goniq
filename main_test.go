package main

import "testing"

func TestRead(t *testing.T) {
	lines := readLines("test.txt")
	correct := []string{"line 1", "line 2"}

	if len(lines) != len(correct) {
		t.Errorf("Output was different length. Got %v, want: %v", len(lines), len(correct))
	}

	for i, v := range lines {
		if v != correct[i] {
			t.Errorf("Output was incorrect. Expected %v, but got %v", correct[i], lines[i])
		}
	}
}

func TestUniq(t *testing.T) {
	list := []string{"test", "test2", "test2", "test3", "test3"}

	correct := []string{"test", "test2", "test3"}

	result := uniq(list)

	if len(result) != len(correct) {
		t.Errorf("Output was a different length. Got %v, want: %v", len(result), len(correct))
	}

	for i, v := range result {
		if v != correct[i] {
			t.Errorf("Output was incorrect. Expected %v, but got %v.", correct[i], list[i])
		}
	}
}
