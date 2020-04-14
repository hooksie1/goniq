package goniq

import (
	"strings"
	"testing"
)

func TestUniq(t *testing.T) {
	words := "test,test2,test3,test2"
	expected := []string{"test", "test2", "test3"}

	list := Uniq(strings.NewReader(words))
	if len(list) != len(expected) {
		t.Errorf("Expected %d items, got %d", len(expected), len(list))
	}

	for i, v := range expected {
		if v != list[i] {
			t.Errorf("Expected %v, got %v", v, list[i])
		}
	}
}
