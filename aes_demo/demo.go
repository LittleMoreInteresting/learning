package aes_demo

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
)

type Aes struct {
	key []byte
	iv  []byte
}

type Option func(aes *Aes)

func WithKey(key string) Option {
	return func(aes *Aes) {
		aes.key = []byte(key)
		aes.iv = reverseKey(key)
	}
}

func reverseKey(s string) []byte {
	str := []byte(s)
	l := 0
	r := len(s)
	for l < r {
		str[l], str[r-1] = str[r-1], str[l]
		l++
		r--
	}
	return str
}
func NewAes(opt ...Option) *Aes {
	aesObj := &Aes{}
	for _, o := range opt {
		o(aesObj)
	}
	return aesObj
}

func (ac *Aes) Encrypt(in string) string {
	//创建加密实例
	block, err := aes.NewCipher(ac.key)
	if err != nil {
		return in
	}
	blockSize := block.BlockSize()
	//填充
	encryptBytes := PKCS5Padding([]byte(in), blockSize)
	//初始化加密数据接收切片
	crypted := make([]byte, len(encryptBytes))
	//使用cbc加密模式
	blockMode := cipher.NewCBCEncrypter(block, ac.iv)
	//执行加密
	blockMode.CryptBlocks(crypted, encryptBytes)
	return base64.StdEncoding.EncodeToString(crypted)
}

func (ac *Aes) Decrypt(in string) string {
	dataByte, err := base64.StdEncoding.DecodeString(in)
	if err != nil {
		panic(err)
	}
	block, err := aes.NewCipher(ac.key)
	if err != nil {
		panic(err)
	}
	blockMode := cipher.NewCBCDecrypter(block, ac.iv)
	orig := make([]byte, len(dataByte))
	// 解密
	blockMode.CryptBlocks(orig, dataByte)
	//去填充
	orig = PKCS5UnPadding(orig)
	return string(orig)
}

func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}
func PKCS5UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}
