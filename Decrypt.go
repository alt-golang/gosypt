package main

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"errors"
)

func DecryptString(key, textb64 string) (string, error) {
	result, err := Decrypt([]byte(key), []byte(textb64))
	return string(result), err
}

func Decrypt(key, textb64 []byte) ([]byte, error) {
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
