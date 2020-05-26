package seed

import "crypto/cipher"

type seed128Cipher struct {
	pdwRoundKey []uint32
}

func newSeed128Cipher(key []byte) cipher.Block {
	// c := seed128Cipher{make([]uint32, n)}
	c := new(seed128Cipher)
	c.pdwRoundKey = seedRoundKey(key)
	return c
}

func (c *seed128Cipher) BlockSize() int { return BlockSize }

func (c *seed128Cipher) Encrypt(dst, src []byte) {
	seedEncrypt(c.pdwRoundKey, dst, src)
}

func (c *seed128Cipher) Decrypt(dst, src []byte) {
	seedDecrypt(c.pdwRoundKey, dst, src)
}
