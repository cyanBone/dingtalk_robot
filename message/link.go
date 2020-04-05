package message

//link类型
type linkMessage struct {
	Text       string `json:"text"`       //消息内容。如果太长只会部分展示
	Title      string `json:"title"`      //消息标题
	MessageURL string `json:"messageUrl"` //点击消息跳转的URL
	PicUrl     string `json:"picUrl"`     //图片URL
}

func NewLinkMessage() *linkMessage {
	return &linkMessage{}
}

func (l *linkMessage) Build() (m message) {
	m = NewMessage("link")
	m.Link = l
	return
}
