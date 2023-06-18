package security

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"os"
)

// GenerateRSAPrivateKey 生成RSA私钥文件
func GenerateRSAPrivateKey(privateKeyPath string, bits int) error {
	privateKey, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		return err
	}
	file, err := os.Create(privateKeyPath)
	if err != nil {
		return err
	}
	defer file.Close()
	err = pem.Encode(file, &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(privateKey),
	})
	if err != nil {
		return err
	}
	return nil
}

// GenerateRSAPublicKey 从私钥生成公钥文件
func GenerateRSAPublicKey(privateKeyPath, publicKeyPath string) error {
	privateKeyFile, err := os.ReadFile(privateKeyPath)
	if err != nil {
		return err
	}
	block, _ := pem.Decode(privateKeyFile)
	if block == nil {
		return errors.New("parse privateKeyFile failed")
	}
	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return err
	}
	file, err := os.Create(publicKeyPath)
	if err != nil {
		return err
	}
	defer file.Close()
	err = pem.Encode(file, &pem.Block{
		Type:  "RSA PUBLIC KEY",
		Bytes: x509.MarshalPKCS1PublicKey(&privateKey.PublicKey),
	})
	if err != nil {
		return err
	}
	return nil
}

// RSAEncrypt RSA加密
func RSAEncrypt(publicKeyPath string, ciphertext []byte) ([]byte, error) {
	publicKeyFile, err := os.ReadFile(publicKeyPath)
	if err != nil {
		return nil, err
	}
	block, _ := pem.Decode(publicKeyFile)
	if block == nil {
		return nil, errors.New("parse publicKeyFile failed")
	}
	key, err := x509.ParsePKCS1PublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	encryptedData, err := rsa.EncryptPKCS1v15(rand.Reader, key, ciphertext)
	if err != nil {
		return nil, err
	}
	return encryptedData, nil
}

// RSADecrypt RSA解密
func RSADecrypt(privateKeyPath string, encryptedData []byte) ([]byte, error) {
	privateKeyFile, err := os.ReadFile(privateKeyPath)
	if err != nil {
		return nil, err
	}
	block, _ := pem.Decode(privateKeyFile)
	if block == nil {
		return nil, errors.New("parse privateKeyFile failed")
	}
	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	decryptedData, err := rsa.DecryptPKCS1v15(rand.Reader, privateKey, encryptedData)
	if err != nil {
		return nil, err
	}
	return decryptedData, nil
}

// RSAStringEncrypt RSA加密
func RSAStringEncrypt(publicKeyStr string, message string) (string, error) {
	block, _ := pem.Decode([]byte(publicKeyStr))
	if block == nil {
		return "", errors.New("parse publicKeyStr failed")
	}
	key, err := x509.ParsePKCS1PublicKey(block.Bytes)
	if err != nil {
		return "", err
	}
	encryptedData, err := rsa.EncryptPKCS1v15(rand.Reader, key, []byte(message))
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(encryptedData), nil
}

// RSAStringDecrypt RSA解密
func RSAStringDecrypt(privateKeyStr string, encryptedDataStr string) (string, error) {
	block, _ := pem.Decode([]byte(privateKeyStr))
	if block == nil {
		return "", errors.New("parse privateKeyStr failed")
	}
	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return "", err
	}
	cipherBytes, err := base64.StdEncoding.DecodeString(encryptedDataStr)
	decryptedData, err := rsa.DecryptPKCS1v15(rand.Reader, privateKey, cipherBytes)
	if err != nil {
		return "", err
	}
	return string(decryptedData), nil
}
