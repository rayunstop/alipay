package sign

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha1"
	"crypto/x509"
	"encoding/base64"
	"errors"
	"fmt"
	"log"
)

func genPubKey(key string) (pubKey *rsa.PublicKey) {

	// 解base64
	encodedKey, err := base64.StdEncoding.DecodeString(key)
	if err != nil {
		log.Fatal(err)
	}

	pkix, err := x509.ParsePKIXPublicKey(encodedKey)
	if err != nil {
		log.Fatal("unable to parse pxix key")
	}
	ok := false

	if pubKey, ok = pkix.(*rsa.PublicKey); !ok {
		log.Fatal("aliPubKey can not be parsed to rsa.PublicKey")
	}
	return
}

// Verfiy 验签函数
func Verfiy(body, sign, aliPubKey string) error {
	//解base64
	decoded, err := base64.StdEncoding.DecodeString(sign)

	if err != nil {
		log.Fatal(err)
	}
	//hashed
	h := sha1.New()
	h.Write([]byte(body))

	//to rsa.publickey
	pubKey := genPubKey(aliPubKey)
	//rsa验签
	return rsa.VerifyPKCS1v15(pubKey, crypto.SHA1, h.Sum(nil), decoded)
}

// Sign 签名
func RsaSign(content, cusPrivKey string) (string, error) {

	//to rsa.privateKey
	privKey := genPrivKeyFromPKSC8(cusPrivKey)
	// TODO content 必须转编码
	hashed := sha1.Sum([]byte(content))
	signed, err := rsa.SignPKCS1v15(rand.Reader, privKey, crypto.SHA1, hashed[:])
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(signed), nil
}

func genPrivKeyFromPKSC8(pkcs8Key string) (privkey *rsa.PrivateKey) {
	// 解base64
	encodedKey, err := base64.StdEncoding.DecodeString(pkcs8Key)
	if err != nil {
		log.Fatal(err)
	}
	// 使用pkcs8格式
	pkcs8, err := x509.ParsePKCS8PrivateKey(encodedKey)

	var ok bool
	if privkey, ok = pkcs8.(*rsa.PrivateKey); !ok {
		log.Fatal(ok)
	}
	return
}

// EncryptAndSignResponse 统一对响应消息签名
// 返回示例：
// <?xml version="1.0" encoding="GBK"?>
// <alipay>
// <response>密文/明文</response>
// <encryption_type>RSA</encryption_type>
// <sign>sign</sign>
// <sign_type>RSA</sign_type>
// </alipay>
func EncryptAndSignResponse(content, cusPrivKey string, isEncrypt, isSign bool) (string, error) {

	builder := `<?xml version=1.0 encoding=GBK?>
				<alipay>
					<response>%s</response>
					<encryption_type>RSA</encryption_type>
					<sign>%s</sign>
					<sign_type>RSA</sign_type>
				</alipay>`

	switch {
	case isEncrypt == true:
		// TODO
		fallthrough
	case isSign == true:
		// sign
		sign, err := RsaSign(content, cusPrivKey)
		if err != nil {
			return "", err
		}
		builder = fmt.Sprintf(builder, content, sign)
	default:
		// 不加密 不签名
		return "", errors.New("params wrong")
	}
	return builder, nil
}
