package qesygo

import (
	"crypto"
	"crypto/aes"
	"crypto/cipher"
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

// 微信小程序解密 （获取电话号码等）
func AesCBCDncrypt(encryptData, key, iv []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}
	blockSize := block.BlockSize()
	if len(encryptData) < blockSize {
		panic("ciphertext too short")
	}
	if len(encryptData)%blockSize != 0 {
		panic("ciphertext is not a multiple of the block size")
	}
	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(encryptData, encryptData)
	// 解填充
	encryptData = PKCS7UnPadding(encryptData)
	return encryptData, nil
}

func PKCS7UnPadding(plantText []byte) []byte {
	length := len(plantText)
	unpadding := int(plantText[length-1])
	return plantText[:(length - unpadding)]
}
