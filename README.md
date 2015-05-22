# 阿里服务窗sdk
##demo
基本sdk开发的一个示例，包括基本的事件，待补充
* 定制自己的client
```go
c := DefaultAlipayClient{
		AppId:       YourAppId,
		ServerURL:   "https://openapi.alipay.com/gateway.do",
		PrivKey:     YourPrivateKey,
		Format:      "json",
		ConTimeOut:  2000,
		ReadTimeOut: 2000,
		SignType:    "RSA",
	}
```
* 选择事件类型
```go
r := &request.AlipayMobilePublicMessageCustomSendRequest{
		BizContent: ImgText,
	}
```
* 执行
```go
resp, err := c.Execute(r)
```
###sdk
参考alipay-java-sdk的实现，开发了go版的alipay-sdk
