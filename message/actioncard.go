package message

//独立跳转ActionCard类型
type actionCardMessage struct {
	Title          string `json:"title"`          //首屏会话透出的展示内容
	Text           string `json:"text"`           //markdown格式的消息
	HideAvatar     string `json:"hideAvatar"`     //0-正常发消息者头像，1-隐藏发消息者头像
	BtnOrientation string `json:"btnOrientation"` //0-按钮竖直排列，1-按钮横向排列
	SingleTitle    string `json:"singleTitle"`    //单个按钮的方案。(设置此项和singleURL后btns无效)
	SingleURL      string `json:"singleURL"`      //点击singleTitle按钮触发的URL
	Btns           []Btn  `json:"btns"`           //按钮的信息：title-按钮方案，actionURL-点击按钮触发的URL
}

type Btn struct {
	Title     string `json:"title"`
	ActionURL string `json:"actionURL"`
}

func NewActionCardMessage() *actionCardMessage {
	return &actionCardMessage{}
}

func (a *actionCardMessage) Build() (m message) {
	m = NewMessage("actionCard")
	m.ActionCard = a
	return
}
