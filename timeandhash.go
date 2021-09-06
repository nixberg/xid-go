package xid

import (
	"encoding/binary"
	"time"

	"github.com/nixberg/xoodyak-go"
)

// TimeAndHash returns a base32-encoded string of
// a 64-bit timestamp followed by a 128-bit hash of input.
func TimeAndHash(time time.Time, input []byte) string {
	bytes := make([]byte, 8, 25)

	offsetBinaryTime := uint64(time.UnixNano()) ^ (1 << 63)
	binary.BigEndian.PutUint64(bytes[:8], offsetBinaryTime)

	xoodyak := xoodyak.New()
	xoodyak.Absorb(input)
	bytes = xoodyak.Squeeze(bytes, 16)

	bytes = append(bytes, idTimeAndHash)

	return encoding.EncodeToString(bytes)[:39]
}

// TimeAndRandomString returns a base32-encoded string of
// a 64-bit timestamp followed by a 128-bit hash of input.
func TimeAndHashString(time time.Time, input string) string {
	return TimeAndHash(time, []byte(input))
}

func IsValidTimeAndHash(xid string) bool {
	return idValid(xid, idTimeAndHash)
}
