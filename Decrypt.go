package main

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"errors"
)

func Decrypt(key, textb64 []byte) ([]byte, error) {
	//text := make([]byte, base64.StdEncoding.DecodedLen(len(textb64)))
	//_, err := base64.StdEncoding.Decode(text, []byte(textb64))
	//if err != nil {
	//	fmt.Println("decode error:", err)
	//}
	text, err := base64.StdEncoding.DecodeString(string(textb64))
	if err != nil {
		return nil, err
	}
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	if len(text) < aes.BlockSize {
		return nil, errors.New("ciphertext too short")
	}
	iv := text[:aes.BlockSize]
	text = text[aes.BlockSize:]
	cfb := cipher.NewCFBDecrypter(block, iv)
	cfb.XORKeyStream(text, text)
	data, err := base64.StdEncoding.DecodeString(string(text))
	if err != nil {
		return nil, err
	}
	return data, nil
}
