package encrypt

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"
)


func AES256Encrypt(plainText []byte, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	// 创建一个随机初始化向量
	iv := make([]byte, aes.BlockSize)
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil, err
	}

	// 使用AES-CFB模式进行加密
	stream := cipher.NewCFBEncrypter(block, iv)
	ciphertext := make([]byte, len(plainText))
	stream.XORKeyStream(ciphertext, plainText)

	// 将初始化向量和密文合并
	result := append(iv, ciphertext...)

	return result, nil
}

func AES256Decrypt(ciphertext []byte, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	// 从密文中提取初始化向量
	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]

	// 使用AES-CFB模式进行解密
	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(ciphertext, ciphertext)

	return ciphertext, nil
}