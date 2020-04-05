钉钉机器人 SDK for Golang

## 安装

```
go get github.com/cyanBone/dingtalk_robot
```

```
import "github.com/cyanBone/dingtalk_robot"
```

## 如何调用

```go
package main

import (
	"github.com/cyanBone/dingtalk_robot"
	"github.com/cyanBone/dingtalk_robot/message"
)

func main() {
	//webhook 机器人地址
	webhook := "https://oapi.dingtalk.com/robot/send?access_token=xxxx"
	//机器人密钥，也可以设置为关键字
	secert := "xxxx"

	//初始化客户端
	client, err := dingtalk_robot.New(webhook, secert)
	if err != nil {
		panic(err)
	}

	//普通文本信息
	textMessage := message.NewTextMessage()
	textMessage.Content="我就是我, 是不一样的烟火111"
    
	//@对应的人
	textMessage.AtMobiles([]string{"1111111111"})
	//@所有人
	textMessage.AtAll(true)

	//发送信息
	err = client.Send(textMessage)
	if err != nil {
		panic(err)
	}
}
```


## 发送类型

### text类型

```
textMessage := message.NewTextMessage()
textMessage.Content="我就是我, 是不一样的烟火"
```

### link类型

```
linkMessage := message.NewLinkMessage()
linkMessage.Text="这个即将发布的新版本，创始人xx称它为“红树林”。而在此之前，每当面临重大升级，产品经理们都会取一个应景的代号，这一次，为什么是“红树林”？"
linkMessage.Title="时代的火车向前开"
linkMessage.MessageURL="https://www.dingtalk.com/s?__biz=xxxx"
```

### markdown类型

```
markdownMessage := message.NewMarkdownMessage()
markdownMessage.Title="杭州天气"
markdownMessage.Text="#### 杭州天气 @156xxxx8827\n" +
    "> 9度，西北风1级，空气良89，相对温度73%\n\n" +
    "> ![screenshot](https://gw.alicdn.com/tfs/TB1ut3xxbsrBKNjSZFpXXcXhFXa-846-786.png)\n"  +
    "> ###### 10点20分发布 [天气](http://www.thinkpage.cn/) \n"
```

### 整体跳转ActionCard类型

```
cardMessage := message.NewActionCardMessage()
cardMessage.Title="乔布斯 20 年前想打造一间苹果咖啡厅，而它正是 Apple Store 的前身"
cardMessage.Text=`![screenshot](@lADOpwk3K80C0M0FoA)
### 乔布斯 20 年前想打造的苹果咖啡厅
Apple Store 的设计正从原来满满的科技感走向生活化，而其生活化的走向其实可以追溯到 20 年前苹果一个建立咖啡馆的计划`
cardMessage.HideAvatar="0"
cardMessage.BtnOrientation="0"
cardMessage.SingleTitle="阅读全文"
cardMessage.SingleURL="https://www.dingtalk.com/"
```

### 独立跳转ActionCard类型

```
cardMessage := message.NewActionCardMessage()
cardMessage.Title = "乔布斯 20 年前想打造一间苹果咖啡厅，而它正是 Apple Store 的前身"
cardMessage.Text = `![screenshot](@lADOpwk3K80C0M0FoA)
 ### 乔布斯 20 年前想打造的苹果咖啡厅
 Apple Store 的设计正从原来满满的科技感走向生活化，而其生活化的走向其实可以追溯到 20 年前苹果一个建立咖啡馆的计划`
cardMessage.HideAvatar = "0"
cardMessage.BtnOrientation = "0"
cardMessage.Btns = []message.Btn{
    {
        Title:     "内容不错",
        ActionURL: "https://www.dingtalk.com/",
    },
    {
        Title:     "不感兴趣",
        ActionURL: "https://www.dingtalk.com/",
    },
}
```

## 开源不易，请大家多多支持

<font color='#0088ff'>微信：</font>
<img width="200" height="200" src="https://video.0-w.cc/assets/images/wx.jpg"/>
<font color='#0088ff'>支付宝：</font>
<img width="200" height="200" src="https://video.0-w.cc/assets/images/zfb.jpg"/>

## License

MIT