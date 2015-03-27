package sign

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha1"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"fmt"
	"log"
)

const (
	aliPubKey = []byte(`-----BEGIN PUBLIC KEY-----
MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQDDI6d306Q8fIfCOaTXyiUeJHkr
IvYISRcc73s3vF1ZT7XN8RNPwJxo8pWaJMmvyTn9N4HQ632qJBVHf8sxHi/fEsra
prwCtzvzQETrNRwVxLO5jVmRGi60j8Ue1efIlzPXV9je9mkjzOmdssymZkh2QhUr
CmZYI/FCEa3/cNMW0QIDAQAB
-----END PUBLIC KEY-----`)
)

var pubKey *rsa.PublicKey

func init() {

	block, _ := pem.Decode(aliPubKey)
	if block == nil {
		log.Fatal("parse PUBLIC KEY PEM error")
	}
	if block.Type != "PUBLIC KEY" {
		log.Fatal("wrong key type" + block.Type)
	}
	key, err := x509.ParsePKIXPublicKey(block.Bytes)

	ok := false
	if pubKey, ok = key.(*rsa.PublicKey); !ok {
		log.Fatal("aliPubKey can not be parsed to rsa.PublicKey")
	}
}

// Verfiy 验签函数
func Verfiy(body, sign string) error {
	//解base64
	decoded, err := base64.StdEncoding.DecodeString(sign)

	if err != nil {
		log.Fatal(err)
	}
	//hashed
	h := sha1.New()
	h.Write([]byte(body))
	//rsa验签
	return rsa.VerifyPKCS1v15(pubKey, crypto.SHA1, h.Sum(nil), decoded)
}
