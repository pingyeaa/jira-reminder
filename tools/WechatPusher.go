package tools

import (
	"bytes"
	"encoding/json"
	"jira-reminder/repository"
	"log"
	"net/http"
)

type WechatPusher struct {
}

func (WechatPusher) Push(msg string, mentionedList []string) {
	params := RequestParams{
		MsgType: "text",
		Text: RequestParamsText{
			Content:       msg,
			MentionedList: mentionedList,
		},
	}
	var jsonByte, err = json.Marshal(params)
	if err != nil {
		log.Fatalln("JSON转换失败", err)
	}

	var buffer = bytes.NewBuffer(jsonByte)
	_, err = http.Post(repository.Cfg.Section("wechat").Key("url").String(), "application/json", buffer)
	if err != nil {
		log.Fatalln("企业微信发送失败", err)
	}
	log.Println("企业推送成功", mentionedList, msg)
}
