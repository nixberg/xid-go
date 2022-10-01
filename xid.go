package xid

import (
	"encoding/base32"
)

const (
	idHash          byte = 0b000_00000
	idRandom        byte = 0b001_00000
	idTimeAndHash   byte = 0b010_00000
	idTimeAndRandom byte = 0b011_00000
)

var (
	encoding = base32.NewEncoding(
		"0123456789abcdefghjkmnpqrstuwxyz",
	).WithPadding(base32.NoPadding)
)

func isValidXID(xid string, id byte) bool {
	if len(xid) != 39 {
		return false
	}
	decoded, err := encoding.DecodeString(xid + "0")
	if err != nil {
		return false
	}
	return decoded[24] == id
}
