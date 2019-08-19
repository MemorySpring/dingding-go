package dingding_go

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func SendMessage(message interface{}, url string) error {
	if message == nil {
		return fmt.Errorf("消息为空")
	}
	body, err := json.Marshal(message)
	if err != nil {
		return fmt.Errorf("序列化消息失败 err:%s message:%+v", err, message)
	}
	res, err := http.Post(url, "application/json", bytes.NewBuffer(body))
	if err != nil {
		fmt.Println("钉钉消息发送失败 ", err)
		return fmt.Errorf("钉钉消息发送失败 err:%s", err)
	}
	if res != nil && res.Body != nil {
		defer res.Body.Close()
		content, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return fmt.Errorf("读取钉钉响应失败 err:%s", err)
		}
		if res.StatusCode == http.StatusFound {
			return fmt.Errorf("机器人可能已被限流 请注意发送频率不要过高 会自行恢复")
		}
		if res.StatusCode != http.StatusOK {
			return fmt.Errorf("钉钉消息发送失败 err:%s", content)
		}
		return nil
	}
	return fmt.Errorf("钉钉响应为空")
}
