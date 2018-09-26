# WorkAPI 企业微信 golang 接口

为了学习golang和企业微信接口，重新造轮子。

json 包使用 [json-iterator](https://github.com/json-iterator/go)

cache 包使用 [go-cache](https://github.com/patrickmn/go-cache)

# 功能

- 实现了应用令牌的获取。（注意: 同一 `CorpID` 下，会有不同应用的 `Secrect`，返回的访问令牌 `Access Token` 也是不同的。调用不同的 `Agent`，要使用不同的 `Access Token`)

- 实现了通讯录中部门、用户、标签信息的获取。

# 使用

```shell
go get -u github.com/sail1972/workapi
```

```golang
package main

import github.com/sail1972/workapi

func main() {
    ...
    wx := workapi.NewWorkAPI(corpid, secrect, appid)
    go wx.GetAccessToken()
    ...

    token, err := wx.GetTokenString()
    if err != nil {
        fmt.Println("Get access token error.")
    }
}
```

# 感谢

本项目开发创意参考了以下项目，特此感谢！

[sdvdxl 的 falcon-message](https://github.com/sdvdxl/falcon-message)

[sbzhu 的 weworkapi_python](https://github.com/sbzhu/weworkapi_python)

[hai046 的 workweixin-go](https://github.com/hai046/workweixin-go)

[silenceper 的 wechat](https://github.com/silenceper/wechat)

[aiportal 的 wechat-proxy](https://github.com/aiportal/wechat-proxy)
