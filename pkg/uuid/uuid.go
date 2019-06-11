package uuid

import (
	"crypto/rand"
	"fmt"
	"io"
)

// UUID represents a simple UUID type.
type UUID string

// Bytes returns a bytes representation of UUID.
func (u UUID) Bytes() []byte {
	var b [16]byte

	n, err := fmt.Sscanf(
		string(u),
		"%02x%02x%02x%02x-%02x%02x-%02x%02x-%02x%02x-%02x%02x%02x%02x%02x%02x",
		&b[0],
		&b[1],
		&b[2],
		&b[3],
		&b[4],
		&b[5],
		&b[6],
		&b[7],
		&b[8],
		&b[9],
		&b[10],
		&b[11],
		&b[12],
		&b[13],
		&b[14],
		&b[15],
	)

	if n != 16 || err != nil {
		return nil
	}

	return b[:]
}

// String returns a string representation of UUID.
func (u UUID) String() string {
	return string(u)
}

// New returns a new random UUID v4.
func New() UUID {
	var b [16]byte

	if _, err := io.ReadFull(rand.Reader, b[:]); err != nil {
		panic("error reading random numbers: " + err.Error())
	}

	b[6] = (b[6] & 0x0f) | (4 << 4)
	b[8] = (b[8] & 0xbf) | 0x80

	return FromBytes(b[:])
}

// FromBytes returns UUID from bytes representation.
func FromBytes(b []byte) UUID {
	return UUID(
		fmt.Sprintf(
			"%08x-%04x-%04x-%04x-%012x",
			b[:4],
			b[4:6],
			b[6:8],
			b[8:10],
			b[10:],
		),
	)
}
