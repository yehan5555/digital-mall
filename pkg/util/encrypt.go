package util

import (
	"bytes"
	"crypto/aes"
	"encoding/base64"
	"errors"
)

//AES 对称加密

//全局加密器

var Encrypt *Encryption

type Encryption struct {
	key string
}

func init() {
	Encrypt = NewEncryption()
}

func NewEncryption() *Encryption {
	return &Encryption{}

}

func PadPwd(srcByte []byte, blockSize int) []byte {
	padNum := blockSize - len(srcByte)%blockSize
	ret := bytes.Repeat([]byte{byte(padNum)}, padNum)
	srcByte = append(srcByte, ret...)
	return srcByte
}

func (k *Encryption) AesEncoding(src string) string {
	srcByte := []byte(src)
	block, err := aes.NewCipher([]byte(k.key))
	if err != nil {
		panic(err)
	}
	// 密码填充
	NewSrcByte := PadPwd(srcByte, block.BlockSize()) // 字节长度不足，进行填充
	dst := make([]byte, len(NewSrcByte))
	block.Encrypt(dst, NewSrcByte)
	//base64编码
	pwd := base64.StdEncoding.EncodeToString(dst)
	return pwd

}

func UnPadPwd(dst []byte) ([]byte, error) {
	if len(dst) < 1 {
		return dst, errors.New("长度有误")
	}
	unpadNum := int(dst[len(dst)-1])
	strErr := "error"
	op := []byte(strErr)
	if len(dst) < unpadNum {
		return op, nil
	}
	str := dst[:len(dst)-unpadNum]
	return str, nil
}

func (k *Encryption) AesDecoding(pwd string) string {
	pwdByte := []byte(pwd)
	pwdByte, err := base64.StdEncoding.DecodeString(string(pwdByte))
	if err != nil {
		panic(err)
	}
	block, errBlock := aes.NewCipher([]byte(k.key))
	if errBlock != nil {
		panic(errBlock)
	}
	dst := make([]byte, len(pwdByte))
	block.Decrypt(dst, pwdByte)
	dst, err = UnPadPwd(dst)
	if err != nil {
		panic(err)
	}
	return string(dst)

}

func (k *Encryption) SetKey(key string) {
	k.key = key
}
