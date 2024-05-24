package enc

import (
	"bytes"
	"crypto/cipher"
	"encoding/hex"
	"crypto/aes"
	"fmt"
)

var key []byte = []byte("")

func padding(data []byte, blockSize int) []byte {
	padding := blockSize - (len(data) % blockSize)
	padText := bytes.Repeat([]byte{byte(padding)}, padding)

	return append(data, padText...)
}

func unpadding(data []byte) ([]byte, error){
	length := len(data)
	unpaddingData := int(data[length-1])
	if unpaddingData > length {
		return nil,fmt.Errorf("Invalid padding")
	}

	return data[:length-unpaddingData], nil
}

func Encrypt(plaintext string) string {
	c, err := aes.NewCipher(key)
	if err != nil {
		fmt.Println("Cipher error.")
	}

	paddedText := padding([]byte(plaintext), c.BlockSize())
	out := make([]byte, len(paddedText))
	enc := cipher.NewCBCEncrypter(c, key[:c.BlockSize()])
	enc.CryptBlocks(out, paddedText)

	return hex.EncodeToString(out)
}

func Decrpyt(plaintext string) string {
	ciphertext, _ := hex.DecodeString(plaintext)
	block, err := aes.NewCipher(key)
	if err != nil {
		fmt.Println("Cipher error.")
	}

	decrypter := cipher.NewCBCDecrypter(block, key[:block.BlockSize()])
	decrypted := make([]byte, len(ciphertext))
	decrypter.CryptBlocks(decrypted, ciphertext)
	unpaddedPlaintext, err := unpadding(decrypted)

	return string(unpaddedPlaintext)
}

