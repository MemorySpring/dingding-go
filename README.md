# dingding-go
钉钉通知GO语言SDK

方便使用GO语言进行钉钉机器人信息通知

支持如下[消息类型](https://ding-doc.dingtalk.com/doc#/serverapi2/qf2nxq)
- text类型
- link类型
- markdown类型
- 独立跳转ActionCard类型

使用示例
```
var dingUrl = ""
var text = &TextMessage{
	MsgType: Text,
	At: MessageAt{
		AtMobiles: []string{},
		IsAtAll:   false,
	},
	Text: MessageContent{
		Content: "test",
	},
}

func TestSendMessage(t *testing.T) {
	err := SendMessage(text, dingUrl)
	if err != nil {
		t.Error(err)
	}
}
```

## TODO 添加构造函数
