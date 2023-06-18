package security

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"
)

func TestRSA(t *testing.T) {
	path, _ := os.Getwd()
	privateKeyPath := filepath.Join(path, "private.pem")
	publicKeyPath := filepath.Join(path, "public.pem")

	err := GenerateRSAPrivateKey(privateKeyPath, 2048)
	if err != nil {
		panic(err)
	}
	err = GenerateRSAPublicKey(privateKeyPath, publicKeyPath)
	if err != nil {
		panic(err)
	}
	data := []byte("Hello, Golang!")
	fmt.Println("Original Data:", string(data))

	encryptedData, err := RSAEncrypt(publicKeyPath, data)
	if err != nil {
		panic(err)
	}
	fmt.Println("Encrypted Data:", encryptedData)

	decryptedData, err := RSADecrypt(privateKeyPath, encryptedData)
	if err != nil {
		panic(err)
	}
	fmt.Println("Decrypted Data:", string(decryptedData))
}

func TestRSAString(t *testing.T) {
	publicKeyStr := `-----BEGIN RSA PUBLIC KEY-----
MIIBCgKCAQEAnpOYPt0YePUtUzCLTx5ywmqYKsZfJmtN4as5keTWFvD88cE5tT3C
R7X4itbfGEzi1ITp3+KlMiH6GkMMZtrafodwmPKcdvB9Fc0rWk2NnrmGMBQrcN+J
gpVbL6oQjgZn7vv3PyHxldg41UPPQsGJTyMCFvNNz4QTHglPmUctJou+uQTBp5YW
pVKtgugDkGx0GhUsRNN3DRSFVtGMSqNO/U0DwbLnJEvJ3jzVtQkqDrMnqyixIbls
H0jSbS6/judWWroXjzzKzj6IdE3NwH7PbM4vF5mEHW8IHrpHUzTMbohl85rfa/Dr
jVOXV2Kcv8iHbXefNtOITDvpGYkRRSfyzQIDAQAB
-----END RSA PUBLIC KEY-----`
	privateKeyStr := `-----BEGIN RSA PRIVATE KEY-----
MIIEpAIBAAKCAQEAnpOYPt0YePUtUzCLTx5ywmqYKsZfJmtN4as5keTWFvD88cE5
tT3CR7X4itbfGEzi1ITp3+KlMiH6GkMMZtrafodwmPKcdvB9Fc0rWk2NnrmGMBQr
cN+JgpVbL6oQjgZn7vv3PyHxldg41UPPQsGJTyMCFvNNz4QTHglPmUctJou+uQTB
p5YWpVKtgugDkGx0GhUsRNN3DRSFVtGMSqNO/U0DwbLnJEvJ3jzVtQkqDrMnqyix
IblsH0jSbS6/judWWroXjzzKzj6IdE3NwH7PbM4vF5mEHW8IHrpHUzTMbohl85rf
a/DrjVOXV2Kcv8iHbXefNtOITDvpGYkRRSfyzQIDAQABAoIBAHsySW5VUKTwPYVK
yn/uNNPsAkoBEX0Ekl4HK1O4B8eM88ZSCqZO07fonK4onuv/F55poFTqfNLE6Mws
WN8zmdoBGS4KFhqlXVhprAudArBUi/agRNuNHwTPSidupSvhuC9Hm1913H1AnW66
HrRLLYTIut/fqAzXHA8Sxr9bN62gXI6E70DArDPZpgbFqnxJmn0Jl30QKg8C/evn
mu6dqUpvPSzOcQmtlCrfD59/8479lrvl3tAhZizM864hLmY5lWMI+33QuK8I46Xp
9IJD1DvaJsaNZjnZVKosBK7XsaU+zK8I76xJh1Akn1HoyoVSt0MyLUrSqTZhgCGA
CHAf5dUCgYEAws2tivVGrHN471dw3nqdkhA7xklPdYfNJp5cBIJQiD3MV9bZMNny
7CIya0uXjZdqpO48FC7/ea9LYAgqSbUeDVdoWY8S13o6/TtBiPTw1zR6Hps8DYkY
yTQMEfWnfHpg0cWC8X8vWvU7nCZEkJ2u4sDvDCafOMHAj1nhAx4+cx8CgYEA0GSB
98CyFtz+c+tcMEM8TZvquOAo/Rrr5/17DUyzxtpo5WrPQ8MmyHvnC9uWFsWkySR3
tXMUa0oHME0hUX8PtEG//kZAwUQG6L6RcWMi+0ifRRcQuieDa26XYxxbcoLqxAMb
KKkg8A1Xdlb2ZP9OKFt+xcAdHxdOolUq+rfRKJMCgYEAwflgNyApMYsB/wGI6GAc
CteTTcyuDJAfHbdOVUcVvbQbh9PuDRDZLXyXy/NHg7wkevqmZEqwJZEBcyxuP1rU
A+DnjVZEIsAZgRLCHQgZ0ZD1kQhuceP0BFWJN1DvZ6nMQtVzn9lPZDkRFFmeqIn7
HNuUrPrATRnRDm0m+53f/W0CgYEAvky6HsvaFWTT/HlWY0BS5jBWpWMWKyQVf3GL
mDaOCS5UEgR6p0+jr/rtn3dz1PHBrGjf5FPltqAQdnxIy8ozRhGwyPvQkGyVvp6f
5KJ6RGwp/Ya1oLkKmuWP21L+81A4IK0RdQ0VZgFY+FkrgkleTx5WYzEvpr+68CTE
LdYEa38CgYBRh8PCX3k1PQz2cnM7IflNTfbezfgZ+Kpkn0dQpLvRTpNrf6zy3Nml
APeOd75Bceo1kLVflqb8aOYbV49tuDF+46elimjSBi7GBBt98QpJlOIIsJW/JdVz
tQ3oxmJLx8u4Gg/h6YpmrmMQCKcAN+4CAOomrCgHXN86wAXHMBl4UA==
-----END RSA PRIVATE KEY-----`
	message := "Hello, Golang!"
	fmt.Println("Original Message:", message)

	encryptedDataStr, err := RSAStringEncrypt(publicKeyStr, message)
	if err != nil {
		panic(err)
	}
	fmt.Println("Encrypted Message:", encryptedDataStr)

	decryptedMessage, err := RSAStringDecrypt(privateKeyStr, encryptedDataStr)
	if err != nil {
		panic(err)
	}
	fmt.Println("Decrypted Message:", decryptedMessage)
}
