package xid

import "github.com/nixberg/xoodyak-go"

// Hash returns a base32-encoded 192-bit hash of inputs.
func Hash(inputs ...[]byte) string {
	bytes := make([]byte, 0, 25)

	xoodyak := xoodyak.New()
	for _, input := range inputs {
		xoodyak.Absorb(input)
	}
	bytes = xoodyak.Squeeze(bytes, 24)

	bytes = append(bytes, idHash)

	return encoding.EncodeToString(bytes)[:39]
}

// HashString returns a base32-encoded 192-bit hash of strings.
func HashStrings(strings ...string) string {
	inputs := [][]byte{}
	for _, s := range strings {
		inputs = append(inputs, []byte(s))
	}
	return Hash(inputs...)
}

func IsValidHash(xid string) bool {
	return isValidXID(xid, idHash)
}
