package tools

type RequestParams struct {
	MsgType string            `json:"msgtype"`
	Text    RequestParamsText `json:"text"`
}

type RequestParamsText struct {
	Content       string   `json:"content"`
	MentionedList []string `json:"mentioned_list"`
}
