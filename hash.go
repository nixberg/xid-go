package xid

import "github.com/nixberg/xoodyak-go"

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
