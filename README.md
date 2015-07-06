# 阿里服务窗sdk

####SDK 说明
* 参考支付宝服务窗java版的实现，开发了golang版的alipay-sdk，包含卡券平台和服务窗平台。
* API类名称与阿里官方提供的sdk包一样，方便迁移，用过java版的直接上手。

####项目进度
* 目前还在开发中，但基础的api调用已通过，接下来会完善细节以及增加更多接口类。
* 配套的demo也在开发中，请移步 https://github.com/z-ray/alipaydemo

####安装调用
* install
```
go get github.com/z-ray/alipay
```
* import
```
github.com/z-ray/alipay
github.com/z-ray/alipay/api/xxx
```

####Alipay Client
基于sdk开发的一个示例，发送图文消息
* 定制自己的client
```go
c := alipay.DefaultAlipayClient{
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


####API
* github.com/z-ray/alipay/api/sign 签名类
* github.com/z-ray/alipay/api/request 事件请求类
* github.com/z-ray/alipay/api/alipass 卡券平台类
* github.com/z-ray/alipay/api/response 事件返回包装类，对应request
* github.com/z-ray/alipay/api/utils 一些通用方法

####TODO
* 常量化事件消息json串，每种请求类型都有固定的消息体
* 补充错误日志
* 补充更多的请求、返回类
* 优化对请求结果的处理


