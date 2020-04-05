package message

//text类型
type textMessage struct {
	Content string `json:"content"` //消息内容
	Ats     *at    `json:"-"`
}

func NewTextMessage() *textMessage {
	return &textMessage{}
}

func (t *textMessage) Build() (m message) {
	m = NewMessage("text")
	m.Text = t
	m.Ats = t.Ats
	return
}

func (t *textMessage) AtAll(isAtAll bool) {
	t.Ats = NewAt(nil, isAtAll)
}

func (t *textMessage) AtMobiles(atMobiles []string) {
	t.Ats = NewAt(atMobiles, false)
}
