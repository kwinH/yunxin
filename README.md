# yunxin

yunxin 是用 GO 语言实现的网易云信的服务端 API 封装，目前实现了常用的大部分功能，如有其他的需要或者功能失效，可以提
issue 告知, 欢迎提 PR 来完善代码。

![](https://img.shields.io/badge/language-golang-blue.svg)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

## 使用方法

#### 安装:

`go get -u github.com/kwinh/yunxin`

#### 导入:

`import "github.com/kwinh/yunxin"`

#### 使用:

##### 获取 token：

```
client := yunxin.CreateImClient("AppKey", "AppSecret", nil)
user := yunxin.ImUser{ID: "3", Name: "test3", Gender: 1}
tk, err := client.CreateImUser(user)
```

##### 发送文本消息

```
msg := yunxin.TextMessage{Message: "hello world"}
err := client.SendTextMessage("1", "3", msg, nil)
```

##### 发送图片

```
msg := yunxin.ImageMessage{URL: "https://golang.org/doc/gopher/frontpage.png", Md5: "可以填任意md5", Extension: "png"}
err := client.SendBatchImageMessage("1", []string{"3"}, msg, nil)
```

##### 发送语音

```
msg := yunxin.VoiceMessage{URL: "audio url", Md5: "可以填任意md5", Duration: 10, Extension: "aac"}
err := client.SendBatchVoiceMessage("1", []string{"3"}, msg, nil)
```

##### 发送视频

```
msg := yunxin.VideoMessage{URL: "video file url", Md5: "可以填任意md5", Extension: "mp4"}
err := client.SendBatchVideoMessage("1", []string{"3"}, msg, nil)
```

## 已实现功能

- [ ] 通信服务
    - [x] 获取 IM 通信 token
    - [x] 更新并获取新 token
    - [x] 用户名片
    - [x] 发送文本消息
    - [x] 发送图片
    - [x] 发送视频
    - [x] 批量发送文本消息
    - [x] 批量发送点对点自定义系统通知
    - [x] 查询单聊历史消息
    - [x] 消息回调
    - [x] 消息抄送

## License

yunxin 使用[MIT](https://opensource.org/licenses/MIT)开源协议
