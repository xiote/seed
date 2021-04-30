package seed

import (
	"bytes"
	"testing"
)

func TestSeedEnDecryption(t *testing.T) {
	t.Parallel()

	cases := []struct {
		Key string
		PT  []byte
		CT  []byte
	}{
		{
			Key: "0123456789123458",
			PT:  []byte{0x83, 0xA2, 0xF8, 0xA2, 0x88, 0x64, 0x1F, 0xB9, 0xA4, 0xE9, 0xA5, 0xCC, 0x2F, 0x13, 0x1C, 0x7D},
			CT:  []byte{0xFA, 0xB0, 0x1D, 0x74, 0x7B, 0x38, 0xA9, 0xB8, 0xD5, 0x50, 0xD7, 0x14, 0xF5, 0xD0, 0x17, 0xBC},
		},
	}

	for i, c := range cases {
		if err := InitECB(c.Key); err != nil {
			t.Fatal(err)
		}

		enc, err := ECBEncryptAll(c.PT)
		if err != nil {
			t.Fatal(err)
		}

		if !bytes.Equal(enc, c.CT) {
			t.Errorf("[%d] failed encryption\n wanted: %x\n got:    %x\n", i, c.CT, enc)
		}

		dec, err := ECBDecryptAll(c.CT)
		if err != nil {
			t.Fatal(err)
		}

		if !bytes.Equal(enc, c.CT) {
			t.Errorf("[%d] failed decryption\n wanted: %x\n got:    %x\n", i, c.PT, dec)
		}
	}
}
