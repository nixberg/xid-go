package xid

import (
	"strings"
	"testing"
	"time"
)

func TestRandom(t *testing.T) {
	xid := Random()

	if !IsValidRandom(xid) {
		t.Error("Random: invalid xid")
	}
	if IsValidHash(xid) || IsValidTimeAndRandom(xid) || IsValidTimeAndHash(xid) {
		t.Error("Random: wrong type")
	}
}

func TestHash(t *testing.T) {
	xid := HashString("")

	if !IsValidHash(xid) {
		t.Error("Hash: invalid xid")
	}
	if IsValidRandom(xid) || IsValidTimeAndRandom(xid) || IsValidTimeAndHash(xid) {
		t.Error("Hash: wrong type")
	}

	if HashString("hi") != "r90f4wukgbk6eatgc3ygq3f204pzymqkzq12ww1" {
		t.Error("Hash: wrong xid")
	}
}

func TestTimeAndRandom(t *testing.T) {
	xid := TimeAndRandom()

	if !IsValidTimeAndRandom(xid) {
		t.Error("TimeAndRandom: invalid xid")
	}
	if IsValidRandom(xid) || IsValidHash(xid) || IsValidTimeAndHash(xid) {
		t.Error("TimeAndRandom: wrong type")
	}

	if !strings.HasPrefix(TimeAndRandom(), xid[:6]) {
		t.Error("TimeAndRandom: missing expected common prefix")
	}
}

func TestTimeAndHash(t *testing.T) {
	xid := TimeAndHashString(time.Now(), "")

	if !IsValidTimeAndHash(xid) {
		t.Error("TimeAndHash: invalid xid")
	}
	if IsValidRandom(xid) || IsValidHash(xid) || IsValidTimeAndRandom(xid) {
		t.Error("TimeAndHash: wrong type")
	}

	if !strings.HasPrefix(TimeAndHashString(time.Now(), ""), xid[:6]) {
		t.Error("TimeAndHash: missing expected common prefix")
	}
	if !strings.HasSuffix(TimeAndHashString(time.Time{}, "hi"), "gj0y9sq70q6cwnn0r7x1e6y40b") {
		t.Error("Hash: wrong suffix")
	}
}
