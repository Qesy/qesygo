package qesygo

import (
	"crypto"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/hex"
	"encoding/pem"
	"log"

	"github.com/forgoer/openssl"
)

// 16进制解码
func HexDecode(s string) []byte {
	dst := make([]byte, hex.DecodedLen(len(s))) //申请一个切片, 指明大小. 必须使用hex.DecodedLen
	n, err := hex.Decode(dst, []byte(s))        //进制转换, src->dst
	if err != nil {
		log.Fatal(err)
		return nil
	}
	return dst[:n] //返回0:n的数据.
}

// 字符串转为16进制
func HexEncode(s string) []byte {
	dst := make([]byte, hex.EncodedLen(len(s))) //申请一个切片, 指明大小. 必须使用hex.EncodedLen
	n := hex.Encode(dst, []byte(s))             //字节流转化成16进制
	return dst[:n]
}

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

// 解密
func RsaDecrypt(ciphertext []byte) ([]byte, error) {
	block, _ := pem.Decode(privateKey)
	if block == nil {
		return nil, errors.New("private key error!")
	}
	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	return rsa.DecryptPKCS1v15(rand.Reader, priv, ciphertext)
}

// 验证
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

// 微信小程序开放数据解密
func AesDecrypt(crypted, key, iv []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockMode := cipher.NewCBCDecrypter(block, iv)
	origData := make([]byte, len(crypted))
	blockMode.CryptBlocks(origData, crypted)
	//获取的数据尾端有'/x0e'占位符,去除它
	for i, ch := range origData {
		if ch == '\x0e' {
			origData[i] = ' '
		}
	}
	//{"phoneNumber":"15082726017","purePhoneNumber":"15082726017","countryCode":"86","watermark":{"timestamp":1539657521,"appid":"wx4c6c3ed14736228c"}}//<nil>
	return origData, nil
}

func PKCS7UnPadding(plantText []byte) []byte {
	length := len(plantText)
	unpadding := int(plantText[length-1])
	return plantText[:(length - unpadding)]
}

func AesECBEncrypt(src, key []byte) []byte { //加密
	dst, _ := openssl.AesECBEncrypt(src, key, openssl.PKCS7_PADDING)
	return dst
}

func AesECBDecrypt(dst, key []byte) []byte { //解密
	dst, _ = openssl.AesECBDecrypt(dst, key, openssl.PKCS7_PADDING)
	return dst
}

func AesCBCEncrypt(src, key, iv []byte) []byte { //加密
	dst, _ := openssl.AesCBCEncrypt(src, key, iv, openssl.PKCS7_PADDING)
	return dst
}

func AesCBCDecrypt(dst, key, iv []byte) []byte { //解密
	dst, _ = openssl.AesCBCDecrypt(dst, key, iv, openssl.PKCS7_PADDING)
	return dst
}
