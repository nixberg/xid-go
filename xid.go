package xid

import (
	"crypto/rand"
	"encoding/base32"
	"encoding/binary"
	"time"

	"github.com/nixberg/xoodyak-go"
)

const (
	idRandom        byte = 0b000_00000
	idHash          byte = 0b001_00000
	idTimeAndRandom byte = 0b010_00000
	idTimeAndHash   byte = 0b011_00000
)

var (
	encoding = base32.NewEncoding(
		"0123456789abcdefghjkmnpqrstuwxyz").WithPadding(base32.NoPadding)
)

func idValid(xid string, id byte) bool {
	if len(xid) != 39 {
		return false
	}
	decoded, err := encoding.DecodeString(xid + "0")
	if err != nil {
		return false
	}
	return decoded[24] == id
}

// Random returns a base32-encoded string of 192 random bits.
func Random() string {
	bytes := make([]byte, 25)

	_, err := rand.Read(bytes[:24])
	if err != nil {
		panic(err)
	}

	bytes[24] = idRandom
	return encoding.EncodeToString(bytes)[:39]
}

func IsValidRandom(xid string) bool {
	return idValid(xid, idRandom)
}

// Hash returns a base32-encoded 192-bit hash of input.
func Hash(input []byte) string {
	bytes := make([]byte, 0, 25)

	xoodyak := xoodyak.New()
	xoodyak.Absorb(input)
	bytes = xoodyak.Squeeze(bytes, 24)

	bytes = append(bytes, idHash)
	return encoding.EncodeToString(bytes)[:39]
}

// HashString returns a base32-encoded 192-bit hash of input.
func HashString(input string) string {
	return Hash([]byte(input))
}

func IsValidHash(xid string) bool {
	return idValid(xid, idHash)
}

// TimeAndRandom returns a base32-encoded string of
// a 64-bit timestamp followed by 128 random bits.
func TimeAndRandom() string {
	bytes := make([]byte, 25)

	offsetBinaryTime := uint64(time.Now().UnixNano()) ^ (1 << 63)
	binary.BigEndian.PutUint64(bytes[:8], offsetBinaryTime)

	_, err := rand.Read(bytes[8:24])
	if err != nil {
		panic(err)
	}

	bytes[24] = idTimeAndRandom
	return encoding.EncodeToString(bytes)[:39]
}

func IsValidTimeAndRandom(xid string) bool {
	return idValid(xid, idTimeAndRandom)
}

// TimeAndRandom returns a base32-encoded string of
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
