package seed

import "crypto/cipher"

type ecb struct {
	b         cipher.Block
	blockSize int
}

func newECB(b cipher.Block) *ecb {
	return &ecb{
		b:         b,
		blockSize: b.BlockSize(),
	}
}

type ecbEncrypter ecb

func newECBEncrypter(b cipher.Block) cipher.BlockMode {
	return (*ecbEncrypter)(newECB(b))
}

func (e *ecbEncrypter) BlockSize() int { return e.blockSize }

func (e *ecbEncrypter) CryptBlocks(dst, src []byte) {
	for len(src) > 0 {
		e.b.Encrypt(dst, src[:e.blockSize])
		src = src[e.blockSize:]
		dst = dst[e.blockSize:]
	}
}

type ecbDecrypter ecb

func NewECBDecrypter(b cipher.Block) cipher.BlockMode {
	return (*ecbDecrypter)(newECB(b))
}

func (d *ecbDecrypter) BlockSize() int { return d.blockSize }

func (d *ecbDecrypter) CryptBlocks(dst, src []byte) {
	for len(src) > 0 {
		d.b.Decrypt(dst, src[:d.blockSize])
		src = src[d.blockSize:]
		dst = dst[d.blockSize:]
	}
}
