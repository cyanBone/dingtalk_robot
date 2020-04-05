package message

//markdown
type markdownMessage struct {
	Title string `json:"title"` //首屏会话透出的展示内容
	Text  string `json:"text"`  //markdown格式的消息
	Ats   *at    `json:"-"`
}

func NewMarkdownMessage() *markdownMessage {
	return &markdownMessage{}
}

func (m2 *markdownMessage) Build() (m message) {
	m = NewMessage("markdown")
	m.Markdown = m2
	return
}

func (m2 *markdownMessage) AtAll(isAtAll bool) {
	m2.Ats = NewAt(nil, isAtAll)
}

func (m2 *markdownMessage) AtMobiles(atMobiles []string) {
	m2.Ats = NewAt(atMobiles, false)
}
