package uuid

import (
	"bytes"
	"testing"
)

const (
	id = UUID("5154e02d-57cd-4da9-9194-d199946681a6")
)

func TestBytes(t *testing.T) {
	good := []byte{0x51, 0x54, 0xe0, 0x2d, 0x57, 0xcd, 0x4d, 0xa9, 0x91, 0x94, 0xd1, 0x99, 0x94, 0x66, 0x81, 0xa6}
	b := id.Bytes()

	if b == nil {
		t.Fatalf("returned nil")
	}

	if !bytes.Equal(good, b) {
		t.Fatalf("wrong result: expected %x, got %x", good, b)
	}
}

func TestBadBytes(t *testing.T) {
	bad := "ZZCAFA12-F257-11E3-B24E-B7BCECBA0006"
	b := UUID(bad).Bytes()

	if b != nil {
		t.Fatalf("converted bad uuid (%s): %x", bad, b)
	}
}

func TestFromBytes(t *testing.T) {
	b := id.Bytes()
	s := FromBytes(b)

	if id != s {
		t.Fatalf("wrong result: expected %s, got %s", id, s)
	}
}

func TestNew(t *testing.T) {
	m := make(map[UUID]interface{})
	var yes interface{}

	for i := 0; i < 10000; i++ {
		v := New()

		if _, ok := m[v]; ok {
			t.Errorf("uuid: collision on %d: %s", i, v)
		}

		m[v] = yes
	}
}
