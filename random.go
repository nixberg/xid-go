package xid

import "github.com/nixberg/chacha-rng-go"

// Random returns a base32-encoded string of 192 random bits.
func Random() string {
	rng, err := chacha.New8()
	if err != nil {
		panic(err)
	}
	return random(rng)
}

func random(rng *chacha.ChaCha) string {
	bytes := make([]byte, 25)

	rng.FillUint8(bytes[:24])

	bytes[24] = idRandom

	return encoding.EncodeToString(bytes)[:39]
}

func IsValidRandom(xid string) bool {
	return idValid(xid, idRandom)
}
