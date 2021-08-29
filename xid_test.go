package xid

import (
	"testing"
)

func TestRandom(t *testing.T) {
	if len(Random()) != 39 {
		t.Error("Random: invalid length")
	}
}

func TestHashString(t *testing.T) {
	if HashString("hi") != "r90f4wukgbk6eatgc3ygq3f204pzymqkzq12ww0" {
		t.Error("HashString")
	}
}
