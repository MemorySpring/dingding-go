package dingding_go

type MessageContent struct {
	Content string `json:"content"`
}

type MessageMarkdown struct {
	Title string `json:"title"`
	Text  string `json:"text"`
}

type MessageLinked struct {
	Text       string `json:"text"`
	Title      string `json:"title"`
	PicUrl     string `json:"picUrl"`
	MessageUrl string `json:"messageUrl"`
}

type MessageActionCard struct {
	Text  string `json:"text"`
	Title string `json:"title"`
	// default "0"
	HideAvatar string `json:"hideAvatar"`
	// default "1"
	// 按钮排列方式 横向or竖向
	BtnOrientation string           `json:"btnOrientation"`
	Buttons        []*MessageButton `json:"btns"`
}

type MessageButton struct {
	Title     string `json:"title"`
	ActionURL string `json:"actionURL"`
}

type MessageAt struct {
	AtMobiles []string `json:"atmobiles"`
	IsAtAll   bool     `json:"isAtAll"`
}

type MessageType string

const (
	Text       MessageType = "text"
	MarkDown   MessageType = "markdown"
	Linked     MessageType = "link"
	ActionCard MessageType = "actionCard"
)

type TextMessage struct {
	MsgType MessageType    `json:"msgtype"`
	At      MessageAt      `json:"at"`
	Text    MessageContent `json:"text"`
}

type MarkdownMessage struct {
	MsgType      MessageType     `json:"msgtype"`
	At           MessageAt       `json:"at"`
	MarkdownText MessageMarkdown `json:"markdown"`
}

type LinkMessage struct {
	MsgType MessageType   `json:"msgtype"`
	Link    MessageLinked `json:"markdown"`
}

type ActionMessage struct {
	MsgType    MessageType       `json:"msgtype"`
	ActionCard MessageActionCard `json:"actionCard"`
}
