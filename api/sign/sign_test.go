package sign

import (
	"crypto"
	"crypto/md5"
	"crypto/rand"
	"crypto/rsa"
	// "fmt"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"testing"
)

var cusPrivKey = `-----BEGIN RSA PRIVATE KEY-----
MIICXwIBAAKBgQDN3lip7UdsCY0lxRO88Fc1DLJzQ6UTsckWyaDAzV6z1zRtSSqa
fKb8Gr0a3OyJqRnqBKkAa+k1NzmIOy2RNAoGMHoR8D5z77nTDhgmBNIORgBjIExg
8AsLIrpz5kbIUnnd5AHFqfDC0GzIrKfg/JRCiEkMltqssxF67Pn0dlK8jwIDAQAB
AoGBAIbqvuSuUm5lXBFytNrQD/b+WTzdCiR8ETNT9Hwmm3f1A7DTkI4qPy1dQK2I
N6SIJCmP4Eocbnuns5aqYdSVbDMwCDsMeKXFvv0tP0A4oxtmKfVqircJF8Fo4E3r
qepeKZoCYfoK2XTVaRVPf5/TPr9g5SzDLTMvDBgnkzL50ouxAkEA+7Fb5UhmuViB
zrFEsKLWqmOXaIUm2bK+38bC1jk1Iuh/uvSrTCnBmNAYuAPrXEBBVIVlQ8ZMydT6
om0tKBDhZQJBANFkPF/jEnqlcICwk/aZEvSkISzmUafq35+H5hAdFXHApRPYFqdt
YsKiJIzExQghoqwN3OUJ3sh4Yl8zXnSxYOMCQQCr6iz2o0lm+AlSAMsGS3OM1pGo
lqo+sFSnzL9wS4r72QrFzDDkyPCvTBRWPHcaf9kfvi246U2x3oODRkc9wqnxAkEA
tpXcG5TK2U2H03+mFkMWh2LTVu6jR3QEeXysPWLO/zkH+UzVPDujAasXifRPBy8l
RTh1ZPb9X/uxc+g8Ni8yFwJBAMa+x/MwmxPpbdmHcI3elpxUlztmRa4Fft8SO5t8
nmrmie/TmjfJpGXqvP2sSv9ZF6wIqIN17guGxLhgPUcAdtU=
-----END RSA PRIVATE KEY-----`

// var cusPrivKey = "MIICeQIBADANBgkqhkiG9w0BAQEFAASCAmMwggJfAgEAAoGBAM3eWKntR2wJjSXFE7zwVzUMsnNDpROxyRbJoMDNXrPXNG1JKpp8pvwavRrc7ImpGeoEqQBr6TU3OYg7LZE0CgYwehHwPnPvudMOGCYE0g5GAGMgTGDwCwsiunPmRshSed3kAcWp8MLQbMisp+D8lEKISQyW2qyzEXrs+fR2UryPAgMBAAECgYEAhuq+5K5SbmVcEXK02tAP9v5ZPN0KJHwRM1P0fCabd/UDsNOQjio/LV1ArYg3pIgkKY/gShxue6ezlqph1JVsMzAIOwx4pcW+/S0/QDijG2Yp9WqKtwkXwWjgTeup6l4pmgJh+grZdNVpFU9/n9M+v2DlLMMtMy8MGCeTMvnSi7ECQQD7sVvlSGa5WIHOsUSwotaqY5dohSbZsr7fxsLWOTUi6H+69KtMKcGY0Bi4A+tcQEFUhWVDxkzJ1PqibS0oEOFlAkEA0WQ8X+MSeqVwgLCT9pkS9KQhLOZRp+rfn4fmEB0VccClE9gWp21iwqIkjMTFCCGirA3c5QneyHhiXzNedLFg4wJBAKvqLPajSWb4CVIAywZLc4zWkaiWqj6wVKfMv3BLivvZCsXMMOTI8K9MFFY8dxp/2R++LbjpTbHeg4NGRz3CqfECQQC2ldwblMrZTYfTf6YWQxaHYtNW7qNHdAR5fKw9Ys7/OQf5TNU8O6MBqxeJ9E8HLyVFOHVk9v1f+7Fz6Dw2LzIXAkEAxr7H8zCbE+lt2Ydwjd6WnFSXO2ZFrgV+3xI7m3yeauaJ79OaN8mkZeq8/axK/1kXrAiog3XuC4bEuGA9RwB21Q=="

func TestSign(t *testing.T) {

	content := "Version=2.0.0&MerchantId=502050000113&MerchOrderId=1432541906918&Amount=1&OrderDesc=2&TradeTime=20150525161826&ExpTime=&NotifyUrl=http://www.xxxxx.com/Notify.do&ExtData=测试&MiscData=13922897656|0||张三|440121197511140912|62220040001154868428||466666||2|&NotifyFlag=0"

	// privKey := genPrivKeyFromPKSC8(cusPrivKey)
	encrypted := md5.Sum([]byte(content))

	p, _ := pem.Decode([]byte(cusPrivKey))
	privKey, _ := x509.ParsePKCS1PrivateKey(p.Bytes)
	// var privKey *rsa.PrivateKey
	// privKey = pk.(*rsa.PrivateKey)

	signed, _ := rsa.SignPKCS1v15(rand.Reader, privKey, crypto.MD5, encrypted[:])

	base64 := base64.StdEncoding.EncodeToString(signed)
	t.Log(base64)
}
