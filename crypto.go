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

	"github.com/thinkoner/openssl"
)

//16进制解码
func HexDecode(s string) []byte {
	dst := make([]byte, hex.DecodedLen(len(s))) //申请一个切片, 指明大小. 必须使用hex.DecodedLen
	n, err := hex.Decode(dst, []byte(s))        //进制转换, src->dst
	if err != nil {
		log.Fatal(err)
		return nil
	}
	return dst[:n] //返回0:n的数据.
}

//字符串转为16进制
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
