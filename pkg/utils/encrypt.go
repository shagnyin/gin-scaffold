package utils

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"encoding/hex"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

func Md5(s string, salt ...string) string {
	h := md5.New()
	h.Write([]byte(s))
	if len(salt) > 0 {
		h.Write([]byte(salt[0]))
	}
	cipher := h.Sum(nil)
	return hex.EncodeToString(cipher)
}

// GeneratePassword  生成用户密码
func GeneratePassword(password string) (string, error) {
	pass := []byte(password)
	fromPassword, err := bcrypt.GenerateFromPassword(pass, bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(fromPassword), nil
}

// ComparePassword .
func ComparePassword(oldPassword, password string) error {
	oldPass := []byte(oldPassword)
	pass := []byte(password)
	return bcrypt.CompareHashAndPassword(oldPass, pass)
}

func AesEncrypt(data []byte, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	encryptBytes := pkcs7Padding(data, blockSize)
	dst := make([]byte, len(encryptBytes))
	blockMode := cipher.NewCBCEncrypter(block, key[:blockSize])
	blockMode.CryptBlocks(dst, encryptBytes)
	return dst, nil
}

func AesDecrypt(data []byte, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	blockMode := cipher.NewCBCDecrypter(block, key[:blockSize])
	dst := make([]byte, len(data))
	blockMode.CryptBlocks(dst, data)
	dst, err = pkcs7UnPadding(dst)
	if err != nil {
		return nil, err
	}
	return dst, nil
}

func pkcs7Padding(data []byte, blockSize int) []byte {
	padding := blockSize - len(data)%blockSize
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(data, padText...)
}

func pkcs7UnPadding(data []byte) ([]byte, error) {
	length := len(data)
	if length == 0 {
		return nil, errors.New("encrypt error")
	}
	unPadding := int(data[length-1])
	return data[:(length - unPadding)], nil
}
