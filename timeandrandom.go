package xid

import (
	"encoding/binary"
	"time"

	"github.com/nixberg/chacha-rng-go"
)

// TimeAndRandom returns a base32-encoded string of
// a 64-bit timestamp followed by 128 random bits.
func TimeAndRandom(time time.Time) string {
	rng, err := chacha.New8()
	if err != nil {
		panic(err)
	}
	return timeAndRandom(time, rng)
}

func timeAndRandom(time time.Time, rng *chacha.ChaCha) string {
	bytes := make([]byte, 25)

	offsetBinaryTime := uint64(time.UnixNano()) ^ (1 << 63)
	binary.BigEndian.PutUint64(bytes[:8], offsetBinaryTime)

	rng.FillUint8(bytes[8:24])

	bytes[24] = idTimeAndRandom

	return encoding.EncodeToString(bytes)[:39]
}

func IsValidTimeAndRandom(xid string) bool {
	return idValid(xid, idTimeAndRandom)
}
