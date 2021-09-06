package xid

import (
	"testing"
	"time"

	"github.com/nixberg/chacha-rng-go"
)

func TestHash(t *testing.T) {
	xid := HashString("")

	if xid != "x8ajyat7qkh4xyu6rhwx9bfhff9j9p06x1fzeqg" {
		t.Error("Hash: wrong value")
	}

	if !IsValidHash(xid) {
		t.Error("Hash: not valid")
	}
	if IsValidRandom(xid) ||
		IsValidTimeAndRandom(xid) ||
		IsValidTimeAndHash(xid) {
		t.Error("Hash: not invalid")
	}
}

func TestRandom(t *testing.T) {
	rng := chacha.Zero8(0)

	if random(rng) != "7r0eybw9bx0dcztuq3m1y2d5m4p883p3std7yes" {
		t.Error("Random: wrong value")
	}
	if random(rng) != "30dy327fe4d1x62cw5sbj8bf86fm8mu78npnc69" {
		t.Error("Random: wrong value")
	}

	if !IsValidRandom(Random()) {
		t.Error("Random: not valid")
	}
	if IsValidHash(Random()) ||
		IsValidTimeAndRandom(Random()) ||
		IsValidTimeAndHash(Random()) {
		t.Error("Random: not invalid")
	}
}

func TestTimeAndHash(t *testing.T) {
	if TimeAndHashString(time.Unix(-9223372036, -854775808), "") !=
		"0000000000001tgn5wnmff729uxpdh3stjpz2yt" {
		t.Error("TimeAndHash: wrong value")
	}
	if TimeAndHashString(time.Unix(0, 0), "") !=
		"g000000000001tgn5wnmff729uxpdh3stjpz2yt" {
		t.Error("TimeAndHash: wrong value")
	}
	if TimeAndHashString(time.Unix(+9223372036, +854775807), "") !=
		"zzzzzzzzzzzzztgn5wnmff729uxpdh3stjpz2yt" {
		t.Error("TimeAndHash: wrong value")
	}

	if !IsValidTimeAndHash(TimeAndHashString(time.Now(), "")) {
		t.Error("TimeAndHash: not valid")
	}
	if IsValidRandom(TimeAndHashString(time.Now(), "")) ||
		IsValidHash(TimeAndHashString(time.Now(), "")) ||
		IsValidTimeAndRandom(TimeAndHashString(time.Now(), "")) {
		t.Error("TimeAndHash: not invalid ")
	}
}

func TestTimeAndRandom(t *testing.T) {
	rng := chacha.Zero8(0)

	if timeAndRandom(time.Unix(-9223372036, -854775808), rng) !=
		"0000000000000fg0xwqrjqt0tsznqe783w4tb8b" {
		t.Error("TimeAndRandom: wrong value")
	}
	if timeAndRandom(time.Unix(0, 0), rng) !=
		"g000000000000b441u1wx6kz7cc1qrc8xxrhm7k" {
		t.Error("TimeAndRandom: wrong value")
	}
	if timeAndRandom(time.Unix(+9223372036, +854775807), rng) !=
		"zzzzzzzzzzzzz62cw5sbj8bf86fm8mu78npnc6b" {
		t.Error("TimeAndRandom: wrong value")
	}

	if !IsValidTimeAndRandom(TimeAndRandom(time.Now())) {
		t.Error("TimeAndRandom: not valid")
	}
	if IsValidRandom(TimeAndRandom(time.Now())) ||
		IsValidHash(TimeAndRandom(time.Now())) ||
		IsValidTimeAndHash(TimeAndRandom(time.Now())) {
		t.Error("TimeAndRandom: not invalid")
	}
}
