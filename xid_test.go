package xid

import "testing"

func TestRandom(t *testing.T) {
	if len(Random()) != 39 {
		t.Error("Random: invalid length")
	}
}

func TestHashString(t *testing.T) {
	if HashString("hi") != "2kase65xtnxgrm4tpd8t1dscae098y1x91bc66a" {
		t.Error("HashString")
	}
}
