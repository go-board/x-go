package xcrypt

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"errors"
)

const AesCipherKeySize = 32 // aes key size, 256 bit

// AesEncryptRaw encrypt original `data` with `token` and initializer vector `iv`
// iv should have at least `aes.BlockSize` length
func AesEncryptRaw(data []byte, key []byte, iv []byte) ([]byte, error) {
	if len(key) < AesCipherKeySize {
		return nil, errors.New("err: key too short, length less than 32")
	}
	key = key[:AesCipherKeySize]
	if len(iv) < aes.BlockSize {
		return nil, errors.New("err: cipher iv too short, length less than 16")
	}
	iv = iv[:aes.BlockSize]
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockMode := cipher.NewCBCEncrypter(block, iv)
	encodedData := pkcs5Padding(data, block.BlockSize())
	encryptedData := make([]byte, len(encodedData))
	blockMode.CryptBlocks(encryptedData, encodedData)
	dst := make([]byte, hex.EncodedLen(len(encryptedData)))
	hex.Encode(dst, encryptedData)
	return dst, nil
}

// AesEncryptRaw encrypt original `data` with `token` and initializer vector `iv`
// iv should have at least `aes.BlockSize` length
func AesDecryptRaw(data []byte, key []byte, iv []byte) ([]byte, error) {
	if len(key) < AesCipherKeySize {
		return nil, errors.New("err: key too short, length less than 32")
	}
	key = key[:AesCipherKeySize]
	if len(iv) < aes.BlockSize {
		return nil, errors.New("err: cipher iv too short, length less than 16")
	}
	iv = iv[:aes.BlockSize]
	decodedData := make([]byte, hex.DecodedLen(len(data)))
	if _, err := hex.Decode(decodedData, data); err != nil {
		return nil, err
	}
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return nil, err
	}
	blockMode := cipher.NewCBCDecrypter(block, iv)
	origData := make([]byte, len(decodedData))
	blockMode.CryptBlocks(origData, decodedData)
	origData = pkcs5UnPadding(origData)
	return origData, nil
}

func pkcs5Padding(cipherText []byte, blockSize int) []byte {
	padding := blockSize - len(cipherText)%blockSize
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(cipherText, padText...)
}

func pkcs5UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}
