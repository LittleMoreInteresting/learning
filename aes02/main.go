package main

import (
	"bytes"
	"crypto/aes"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"log"
)

func main() {
	/*
	 * src 要加密的字符串
	 * key 用来加密的密钥 密钥长度可以是128bit、192bit、256bit中的任意一个
	 * 16位key对应128bit
	 */

	src := `13164238899`

	ae := &AesECB{
		key:       []byte("y27bulYuw6cmm@ln"),
		blockSize: aes.BlockSize,
	}

	// 普通文本的加密与解密
	// ECB加密
	encrypted1, _ := ae.Encrypt([]byte(src))
	log.Println(
		hex.EncodeToString(encrypted1),                //输出16进制串
		base64.StdEncoding.EncodeToString(encrypted1), //输出16进制串
	)

	// ECB解密
	decrypted1, _ := ae.Decrypt(encrypted1)
	log.Println(
		string(decrypted1),
	)

	/*	src = "aa10"
		// 16进制的加密与解密
		hb, _ := hex.DecodeString(src)
		encrypted2, _ := ae.Encrypt(hb)
		log.Println(
			hex.EncodeToString(encrypted2),
		)

		// ECB解密
		decrypted2, _ := ae.Decrypt(encrypted2)
		log.Println(
			hex.EncodeToString(decrypted2),
		)*/
}

type AesECB struct {
	key       []byte
	blockSize int
}

func (a *AesECB) Encrypt(src []byte, isPad ...bool) ([]byte, error) {
	block, err := aes.NewCipher(a.key)
	if err != nil {
		return nil, err
	}

	if len(src) == 0 {
		return nil, errors.New("content is empty")
	}

	if len(isPad) > 0 && isPad[0] == false {
		src = a.noPadding(src)
	} else {
		src = a.padding(src)
	}

	buf := make([]byte, a.blockSize)
	encrypted := make([]byte, 0)
	for i := 0; i < len(src); i += a.blockSize {
		block.Encrypt(buf, src[i:i+a.blockSize])
		encrypted = append(encrypted, buf...)
	}
	return encrypted, nil
}

func (a *AesECB) Decrypt(src []byte, isPad ...bool) ([]byte, error) {
	block, err := aes.NewCipher(a.key)
	if err != nil {
		return nil, err
	}
	if len(src) == 0 {
		return nil, errors.New("content is empty")
	}
	buf := make([]byte, a.blockSize)
	decrypted := make([]byte, 0)
	for i := 0; i < len(src); i += a.blockSize {
		block.Decrypt(buf, src[i:i+a.blockSize])
		decrypted = append(decrypted, buf...)
	}

	if len(isPad) > 0 && isPad[0] == false {
		decrypted = a.unNoPadding(decrypted)
	} else {
		decrypted = a.unPadding(decrypted)
	}

	return decrypted, nil
}

//nopadding模式
func (a *AesECB) noPadding(src []byte) []byte {
	count := a.blockSize - len(src)%a.blockSize
	if len(src)%a.blockSize == 0 {
		return src
	} else {
		return append(src, bytes.Repeat([]byte{byte(0)}, count)...)
	}
}

//nopadding模式
func (a *AesECB) unNoPadding(src []byte) []byte {
	for i := len(src) - 1; ; i-- {
		if src[i] != 0 {
			return src[:i+1]
		}
	}
	return nil
}

//padding模式
func (a *AesECB) padding(src []byte) []byte {
	count := a.blockSize - len(src)%a.blockSize
	padding := bytes.Repeat([]byte{byte(0)}, count)
	padding[count-1] = byte(count)
	return append(src, padding...)
}

//padding模式
func (a *AesECB) unPadding(src []byte) []byte {
	l := len(src)
	p := int(src[l-1])
	return src[:l-p]
}
