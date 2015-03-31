package main

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha1"
	"crypto/x509"
	// "encoding/hex"
	"encoding/pem"
	"fmt"
	"net/http"

	"io/ioutil"
	// "os"
	"bytes"
	"encoding/base64"
	"github.com/qiniu/iconv"
	"log"
	"net/url"
	// "strings"
	// "os"
	"time"
)

var (
	publicKey = `-----BEGIN PUBLIC KEY-----
MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQDDI6d306Q8fIfCOaTXyiUeJHkr
IvYISRcc73s3vF1ZT7XN8RNPwJxo8pWaJMmvyTn9N4HQ632qJBVHf8sxHi/fEsra
prwCtzvzQETrNRwVxLO5jVmRGi60j8Ue1efIlzPXV9je9mkjzOmdssymZkh2QhUr
CmZYI/FCEa3/cNMW0QIDAQAB
-----END PUBLIC KEY-----`

	privateKey = `-----BEGIN RSA PRIVATE KEY-----
MIICXgIBAAKBgQDdJAQqGm0tHaMs0cgHl29N3gFv9aSsCcKFcK+edI4OQFl0iLt6
U4In/st9XXJMQjN2Ltun6JsD3cHEx1iNmE26H2Z+C/AU6usaqnLQwmQnAhvik7XE
/wkHAhcNRq55qCm6Xt48yrmE6hkO5NH2y6DQIIdiaYC5XhKNqWb7tezLJQIDAQAB
AoGBALmTYN9IP/hNV8Lj5N3iCiipNkGTPXaV1iSPFQF/RDrXa3psyA92htIzcuao
haNTJsZ1uiVlALk03kfZFgn1FrubIRvLtJTFVUF+bz+fp8KlZklcDB3/nlys4rfB
FHvbwQhqYVSuGKOGZfOKvjaTRh+wXlMcyLr9jldZHbPmRRUBAkEA/oIg1z7ba9dz
u8FGAe/SvmT9Ax0kJIqdfFqh67HCrm5FFXlyhV50N6fdDDzAOTcSwLh4rHDKluLE
CYRxk2MiQQJBAN5v0n0wZDkjZmsW7rzAht24Aqavh5ybR6cmpyIffuXmDt6wMpU+
m4fp6GwybqQw7ZdNflCG6kQoPOrKKr+3Z+UCQQDXgM5YFFRtg1jvIZ+q4ix7pT2M
Fm/VNT5W3tN+pN1pH9wFa/mprqoPumb1BrfpepW5dDpSIYuZqdg/CtO07ltBAkBw
loEgRKI2Gaj5g34LpBefmkgdPrORnTdDb9kg+HguvafBJ8YyrKHkxYyTV2ORUAKy
ltLcx61EGmnbHcFNkPPRAkEAmY4er0TC6IYSOTfLHdiaIbo3cuFytYOA5cv9jCt6
gTtQgAiJ55P/fLnQ22f6hCT7KLAcFKEAofpKsPQ+fnmSzQ==
-----END RSA PRIVATE KEY-----`

	privKey = "MIICdwIBADANBgkqhkiG9w0BAQEFAASCAmEwggJdAgEAAoGBAKK0PXoLKnBkgtOl0kvyc9X2tUUdh/lRZr9RE1frjr2ZtAulZ+Moz9VJZFew1UZIzeK0478obY/DjHmD3GMfqJoTguVqJ2MEg+mJ8hJKWelvKLgfFBNliAw+/9O6Jah9Q3mRzCD8pABDEHY7BM54W7aLcuGpIIOa/qShO8dbXn+FAgMBAAECgYA8+nQ380taiDEIBZPFZv7G6AmT97doV3u8pDQttVjv8lUqMDm5RyhtdW4n91xXVR3ko4rfr9UwFkflmufUNp9HU9bHIVQS+HWLsPv9GypdTSNNp+nDn4JExUtAakJxZmGhCu/WjHIUzCoBCn6viernVC2L37NL1N4zrR73lSCk2QJBAPb/UOmtSx+PnA/mimqnFMMP3SX6cQmnynz9+63JlLjXD8rowRD2Z03U41Qfy+RED3yANZXCrE1V6vghYVmASYsCQQCoomZpeNxAKuUJZp+VaWi4WQeMW1KCK3aljaKLMZ57yb5Bsu+P3odyBk1AvYIPvdajAJiiikRdIDmi58dqfN0vAkEAjFX8LwjbCg+aaB5gvsA3t6ynxhBJcWb4UZQtD0zdRzhKLMuaBn05rKssjnuSaRuSgPaHe5OkOjx6yIiOuz98iQJAXIDpSMYhm5lsFiITPDScWzOLLnUR55HL/biaB1zqoODj2so7G2JoTiYiznamF9h9GuFC2TablbINq80U2NcxxQJBAMhw06Ha/U7qTjtAmr2qAuWSWvHU4ANu2h0RxYlKTpmWgO0f47jCOQhdC3T/RK7f38c7q8uPyi35eZ7S1e/PznY="
	// requestUrl = "https://openapi.alipay.com/gateway.do?app_id=2014121600020041&biz_content=%s&charset=GBK&method=%s&sign_type=RSA&timestamp=%s&version=1.0&sign=%s"

	requestBody = "app_id=2014121600020041&biz_content=%s&charset=GBK&method=%s&sign_type=RSA&timestamp=%s&version=1.0"

	methodMSG = "alipay.mobile.public.message.custom.send"

	methodTEM = "alipay.mobile.public.message.single.send"
	//图文消息
	ImgText = "{'articles':[{'actionName':'立即查看','desc':'这是图文内容','imageUrl':'http://pic.alipayobjects.com/e/201311/1PaQ27Go6H_src.jpg','title':'这是标题','url':'https://www.alipay.com/'}],'msgType':'image-text','toUserId':'kquy1L6pBsGPHDYCYeBhLJqKo7ILToCxhO4skMe5bRK-S4-HyunMvYTqiRYEUD+U01'}"

	//模板消息
	templateText = "{" +
		"   'toUserId':'kquy1L6pBsGPHDYCYeBhLJqKo7ILToCxhO4skMe5bRK-S4-HyunMvYTqiRYEUD+U01'," +
		"   'template':{" +
		"       'templateId':'4e400db8fe204734bf5938ee6a1d916d'," +
		"       'context':{" +
		"           'headColor':'#85be53'," +
		"           'url':'http://121.41.85.237:8080/microinter/static/menulist/foodmenu.html?foodListId=100'," +
		"           'actionName':'查看详情'," +
		"           'first':{" +
		"               'color':'#000000'," +
		"               'value':'亲，您的订单信息如下：'" +
		"           }," +
		"           'keyword1':{" +
		"               'color':'#000000'," +
		"               'value':'go商店'" +
		"           }," +
		"           'keyword2':{" +
		"               'color':'#000000'," +
		"               'value':'2015-03-24 15:00:00'" +
		"           }," +
		"           'keyword3':{" +
		"               'color':'#000000'," +
		"               'value':'65788889'" +
		"           }," +
		"           'keyword4':{" +
		"               'color':'#000000'," +
		"               'value':'餐饮'" +
		"           }," +
		"           'remark':{" +
		"               'color':'#ff0000'," +
		"               'value':'谢谢您的消费'" +
		"           }" +
		"       }" +
		"   }" +
		"}"
)

func main() {

	signature()

	// addr := ":8080"
	// log.Printf("QuickPay is running on %s", addr)

	// http.Handle("/", http.FileServer(http.Dir("static")))

	// http.HandleFunc("/microinter/gateway.do", func(w http.ResponseWriter, r *http.Request) {

	// })

	// log.Fatal(http.ListenAndServe(addr, nil))

}

func signature() {

	cd, err := iconv.Open("gbk", "utf-8") // convert utf-8 to gbk
	if err != nil {
		fmt.Println("iconv.Open failed!")
		return
	}
	defer cd.Close()

	// priblock, _ := pem.Decode([]byte(privateKey))
	// if priblock == nil {
	// 	panic("failed to parse certificate PEM")
	// }

	// cert, err := x509.ParseCertificate(pblock.Bytes)
	// key, err := x509.ParsePKCS1PrivateKey(priblock.Bytes)

	// fmt.Println(privateKey)

	encodedKey, err := base64.StdEncoding.DecodeString(privKey)
	if err != nil {
		log.Fatal(err)
	}
	private, err := x509.ParsePKCS8PrivateKey(encodedKey)
	var key *rsa.PrivateKey
	var ok bool
	if key, ok = private.(*rsa.PrivateKey); !ok {
		log.Fatal(ok)
	}

	// log.Println(private, err)
	publock, _ := pem.Decode([]byte(publicKey))
	if publock == nil {
		log.Fatal("Could not parse Certificate PEM")
	}
	// if PEMBlock.Type != "CERTIFICATE" {
	// 	log.Println(PEMBlock.Type)
	// 	log.Fatal("Found wrong key type" + PEMBlock.Type)
	// }
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	// pk, err := x509.ParsePKIXPublicKey(publock.Bytes)
	// fmt.Println(pk)

	gbk := cd.ConvString(fmt.Sprintf(requestBody, templateText, methodTEM, timestamp))
	hashed := sha1.Sum([]byte(gbk))
	sign, err := rsa.SignPKCS1v15(rand.Reader, key, crypto.SHA1, hashed[:])

	//base64
	signature := base64.StdEncoding.EncodeToString(sign)

	//
	u, err := url.Parse("https://openapi.alipay.com/gateway.do")
	q := u.Query()
	q.Set("app_id", "2014121600020041")
	q.Set("biz_content", cd.ConvString(templateText))
	q.Set("charset", "GBK")
	q.Set("method", methodTEM)
	q.Set("sign_type", "RSA")
	q.Set("timestamp", timestamp)
	q.Set("version", "1.0")
	q.Set("sign", signature)

	u.RawQuery = q.Encode()
	// ugbk := cd.ConvString(u.String())
	postMsg(u.String())
}

func postMsg(u string) {

	cd, err := iconv.Open("utf-8", "gbk") // convert utf-8 to gbk
	if err != nil {
		fmt.Println("iconv.Open failed!")
		return
	}
	defer cd.Close()

	log.Println(u)
	resp, err := http.Post(u, "POST", bytes.NewBufferString("test"))

	if err != nil {
		panic(err)
	}
	resMsg, err := ioutil.ReadAll(resp.Body)

	gbk := cd.ConvString(string(resMsg))
	fmt.Println(gbk)
	if err != nil {
		log.Println(gbk)
	}
	// fmt.Printf("%s", resMsg)
	// fmt.Println(sign)
}
