package QesyGo

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
)

// 加密
func RsaEncrypt(origData []byte, Privatekey []byte) ([]byte, error) {
	block, _ := pem.Decode(Privatekey)
	if block == nil {
		return nil, Err("get block err")
	}
	private, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	h := crypto.Hash.New(crypto.SHA1)
	h.Write(origData)
	hashed := h.Sum(nil)
	return rsa.SignPKCS1v15(rand.Reader, private, crypto.SHA1, hashed)
}
