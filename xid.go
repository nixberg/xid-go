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

func Random() string {
	bytes := make([]byte, 24)
	_, err := rand.Read(bytes)
	if err != nil {
		panic(err)
	}
	return encoding.EncodeToString(bytes)
}

func Hash(input []byte) string {
	xoodyak := xoodyak.New()
	xoodyak.Absorb(input)
	digest := xoodyak.Squeeze(nil, 24)
	return encoding.EncodeToString(digest)
}

func HashString(input string) string {
	return Hash([]byte(input))
}
