package yunxin

type ResponseBase struct {
	Code int    `json:"code"`
	Desc string `json:"desc,omitempty"`
}

type Response[T any] struct {
	ResponseBase
	Data *T `json:"data"`
}

// TokenInfo 云通信Token
type TokenInfo struct {
	Token string `json:"token"`
	Accid string `json:"accid"`
	Name  string `json:"name"`
}

type Uinfo struct {
	Accid     string `json:"accid,omitempty"`
	Name      string `json:"name,omitempty"`
	Icon      string `json:"icon,omitempty"`
	Sign      string `json:"sign,omitempty"`
	Email     string `json:"email,omitempty"`
	Birth     string `json:"birth,omitempty"`
	Mobile    string `json:"mobile,omitempty"`
	Ex        string `json:"ex,omitempty"`
	Gender    int    `json:"gender,omitempty"`
	Valid     bool   `json:"valid,omitempty"`
	Mute      bool   `json:"mute,omitempty"`
	MuteP2P   bool   `json:"muteP2P,omitempty"`
	MuteQChat bool   `json:"muteQChat,omitempty"`
	MuteTeam  bool   `json:"muteTeam,omitempty"`
	MuteRoom  bool   `json:"muteRoom,omitempty"`
}

// MessageHistory .
type MessageHistory struct {
	From string      `json:"from"`
	ID   int64       `json:"msgid"`
	Time int64       `json:"sendtime"`
	Type int         `json:"type"`
	Body interface{} `json:"body"`
}

type Message struct {
	MsgId    int64 `json:"msgid"`
	TimeTag  int64 `json:"timetag"`
	Antispam bool  `json:"antispam"`
}

// Broadcast 广播推送结果
type Broadcast struct {
	BroadcastID int64    `json:"broadcastId"`
	ExpireTime  int64    `json:"expireTime"`
	Body        string   `json:"body"`
	CreateTime  int64    `json:"createTime"`
	IsOffline   bool     `json:"isOffline"`
	TargetOs    []string `json:"targetOs"`
}
