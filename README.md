# 阿里服务窗sdk
##Alipay Client
```
go get github.com/z-ray/alipay
go get github.com/z-ray/alipay/api/xxx
```
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

###API
* github.com/z-ray/alipay/api/sign 签名类
* github.com/z-ray/alipay/api/request 事件请求类
* github.com/z-ray/alipay/api/alipass 卡券平台类
* github.com/z-ray/alipay/api/response 事件返回包装类，对应request
* github.com/z-ray/alipay/api/utils 一些通用方法
* github.com/z-ray/alipay/api/conver 解析map，映射到结构体


