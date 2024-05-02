package mssages

import (
	"encoding/hex"
	"crypto/rand"
	"time"
	"XoRPC/XoRPC/src/encrypt"
)


func RandAes256(size int) (str string) {
	b := make([]byte, size)
	rand.Read(b)
	rand_str := hex.EncodeToString(b)
	return rand_str
}