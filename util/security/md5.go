package security

import (
	"crypto/md5"
	"encoding/hex"
)

// md5值长度为128bit，对应16进制表示的长度为 32 digits

func MD5(str string) string {
	b := []byte(str)
	h := md5.New()
	h.Write(b)
	return hex.EncodeToString(h.Sum(nil))
}

func MD5_SALT(str string, salt string) string {
	b := []byte(str)
	s := []byte(salt)
	h := md5.New()
	h.Write(s) // 先写盐值
	h.Write(b)
	return hex.EncodeToString(h.Sum(nil))
}
