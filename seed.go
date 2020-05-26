package seed

import (
	"crypto/cipher"
	"errors"
	"fmt"
)

type KeySizeError int

func (k KeySizeError) Error() string {
	return fmt.Sprintf("Invalid key size, key size must be 16: %d", int(k))
}

type seed128ECB struct {
	encrypter cipher.BlockMode
	decrypter cipher.BlockMode
}

func (x *seed128ECB) encryptBlock(src []byte) []byte {
	dst := make([]byte, 16)
	x.encrypter.CryptBlocks(dst, src)
	return dst

}

func (x *seed128ECB) decryptBlock(src []byte) []byte {
	dst := make([]byte, 16)
	x.decrypter.CryptBlocks(dst, src)
	return dst
}

var (
	ecbEncDecrypter seed128ECB
)

func InitECB(key string) error {
	keyByte := []byte(key)
	l := len(keyByte)

	switch l {
	case 16:
		createECBEncDecrypter(keyByte)
	case 32:
		return errors.New("Unsupported key size 32")
	default:
		return KeySizeError(l)
	}
	return nil
}

func createECBEncDecrypter(key []byte) {
	block := newSeed128Cipher(key)
	ecbEncDecrypter.encrypter = newECBEncrypter(block)
	ecbEncDecrypter.decrypter = NewECBDecrypter(block)
}

func ECBEncryptAll(plainText []byte) ([]byte, error) {
	if ecbEncDecrypter.encrypter == nil {
		return nil, errors.New("Not initialized, initialize with InitECB()")
	}
	var encrypt []byte
	var padding int
	if l := len(plainText) % 16; l != 0 {
		padding = 16 - l
	}
	for i := 0; i < padding; i++ {
		plainText = append(plainText, 0)
	}

	for i := 0; i < len(plainText); i += 16 {
		seg := ecbEncDecrypter.encryptBlock(plainText[i : i+16])
		encrypt = append(encrypt, seg...)
	}
	return encrypt, nil
}

func ECBDecryptAll(encrypted []byte) ([]byte, error) {
	if ecbEncDecrypter.decrypter == nil {
		return nil, errors.New("Not initialized, initialize with InitECB()")
	}
	if len(encrypted)%16 != 0 {
		return nil, errors.New("Invalid encrypted message")
	}
	var decrypt []byte

	for i := 0; i < len(encrypted); i += 16 {
		seg := ecbEncDecrypter.decryptBlock(encrypted[i : i+16])
		decrypt = append(decrypt, seg...)
	}
	return decrypt, nil
}
