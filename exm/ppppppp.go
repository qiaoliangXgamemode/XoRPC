package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"
	"fmt"
	"encoding/hex"
)

func CreatAES256(key []byte) (Keys cipher.Block, e error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	iv := make([]byte, aes.BlockSize)
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil, err
	}
	return block, nil
}

func RandAes256(size int) (str string) {
	b := make([]byte, size)
	rand.Read(b)
	rand_str := hex.EncodeToString(b)
	return rand_str
}

func main() {
	k, _ := CreatAES256([]byte("FGHJKDFVBGK1DFGhfs3443242234adbg"))
	a := RandAes256(32)
	fmt.Println(a)
	fmt.Println(k)
}