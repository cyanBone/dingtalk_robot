package message

//Builder
type Builder interface {
	Build() (m message)
}

//base message
type message struct {
	Msgtype    string             `json:"msgtype"`
	Text       *textMessage       `json:"text,omitempty"`
	Link       *linkMessage       `json:"link,omitempty"`
	Markdown   *markdownMessage   `json:"markdown,omitempty"`
	ActionCard *actionCardMessage `json:"actionCard,omitempty"`
	FeedCard   *feedCardMessage   `json:"feedCard,omitempty"`
	Ats        *at                `json:"at,omitempty"`
}

//at
type at struct {
	AtMobiles []string `json:"atMobiles,omitempty"` //被@人的手机号
	IsAtAll   bool     `json:"isAtAll,omitempty"`   //@所有人时：true，否则为：false
}

//new at
func NewAt(atMobiles []string, isAtAll bool) *at {
	return &at{AtMobiles: atMobiles, IsAtAll: isAtAll}
}

//new message
func NewMessage(msgtype string) message {
	return message{Msgtype: msgtype}
}

//at something
func (t message) At(isAtAll bool, atMobiles []string) {
	t.Ats.AtMobiles = atMobiles
	t.Ats.IsAtAll = isAtAll
}
