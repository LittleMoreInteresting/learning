package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"fmt"
)

func main() {
	key := []byte("y27bulYuw6cmm@ln")
	plaintext := []byte("13081450221")

	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	paddedPlaintext := pkcs5Padding(plaintext, aes.BlockSize)
	ciphertext := make([]byte, len(paddedPlaintext))

	ecb := NewECBEncrypter(block)
	ecb.CryptBlocks(ciphertext, paddedPlaintext)

	fmt.Printf("%v\n", base64.StdEncoding.EncodeToString(ciphertext))
	fmt.Printf("%v\n", ciphertext)
	s, _ := base64.StdEncoding.DecodeString("nviBhS4T0yuNSP8cqtTjWw==")
	fmt.Printf("%v\n", s)
	ecb = NewECBDecrypter(block)
	ecb.CryptBlocks(ciphertext, s)
	padding := ciphertext[len(ciphertext)-1]
	fmt.Printf("%v\n", padding)
	fmt.Printf("%v\n", ciphertext[:len(ciphertext)-int(padding)])
	fmt.Printf("%v\n", string(ciphertext[:len(ciphertext)-int(padding)]))
}

func pkcs5Padding(data []byte, blockSize int) []byte {
	padding := blockSize - len(data)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(data, padtext...)
}

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

// NewECBEncrypter returns a BlockMode which encrypts in electronic code book
// mode, using the given Block.
func NewECBEncrypter(b cipher.Block) cipher.BlockMode {
	return (*ecbEncrypter)(newECB(b))
}
func (x *ecbEncrypter) BlockSize() int { return x.blockSize }
func (x *ecbEncrypter) CryptBlocks(dst, src []byte) {
	if len(src)%x.blockSize != 0 {
		panic("crypto/cipher: input not full blocks")
	}
	if len(dst) < len(src) {
		panic("crypto/cipher: output smaller than input")
	}
	for len(src) > 0 {
		x.b.Encrypt(dst, src[:x.blockSize])
		src = src[x.blockSize:]
		dst = dst[x.blockSize:]
	}
}

type ecbDecrypter ecb

// NewECBDecrypter returns a BlockMode which decrypts in electronic code book
// mode, using the given Block.
func NewECBDecrypter(b cipher.Block) cipher.BlockMode {
	return (*ecbDecrypter)(newECB(b))
}
func (x *ecbDecrypter) BlockSize() int { return x.blockSize }
func (x *ecbDecrypter) CryptBlocks(dst, src []byte) {
	if len(src)%x.blockSize != 0 {
		panic("crypto/cipher: input not full blocks")
	}
	if len(dst) < len(src) {
		panic("crypto/cipher: output smaller than input")
	}
	for len(src) > 0 {
		x.b.Decrypt(dst, src[:x.blockSize])
		src = src[x.blockSize:]
		dst = dst[x.blockSize:]
	}
}
