package security

import (
	"fmt"
	"testing"
)

func TestAES(t *testing.T) {
	plainText, key := "abc", "abcdefgehjhijkmlkjjwwoew"
	cipherText, err := AESEncrypt(plainText, key)
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(cipherText)

	decodeText, err := AESDecrypt(cipherText, key)
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(decodeText)
}
