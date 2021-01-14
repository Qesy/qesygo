package qesygo

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

func RsaVeri(origData []byte, PublicKey []byte, Sign string) error {
	block, _ := pem.Decode(PublicKey)
	if block == nil {
		return Err("get block err")
	}
	public, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return err
	}
	h := crypto.Hash.New(crypto.SHA1)
	h.Write(origData)
	hashed := h.Sum(nil)
	return rsa.VerifyPKCS1v15(public.(*rsa.PublicKey), crypto.SHA1, hashed, []byte(Sign))
}
