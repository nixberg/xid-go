package xid

import (
	"crypto/rand"
	"encoding/base32"

	"github.com/nixberg/xoodyak-go"
)

var (
	encoding = base32.NewEncoding(
		"abcdefghjkmnpqrstuwxyz0123456789").WithPadding(base32.NoPadding)
)

// Random returns a random, base32-encoded 192-bit string.
func Random() string {
	bytes := make([]byte, 24)
	_, err := rand.Read(bytes)
	if err != nil {
		panic(err)
	}
	return encoding.EncodeToString(bytes)
}

// Hash returns a base32-encoded 192-bit hash of input.
func Hash(input []byte) string {
	xoodyak := xoodyak.New()
	xoodyak.Absorb(input)
	digest := xoodyak.Squeeze(nil, 24)
	return encoding.EncodeToString(digest)
}

// HashString returns a base32-encoded 192-bit hash of input.
func HashString(input string) string {
	return Hash([]byte(input))
}
