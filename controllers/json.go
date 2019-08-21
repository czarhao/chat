package controllers

// Operating为操作码
const (
	CREATE = iota
	SHOW
	JOIN
)

// 聊天操作码
const (
	LOGIN = iota
	CHAT
	QUIT
	WHO
	ME
)

// 接收和发送的json请求
// 下面表示是接收到的json
type getIndexOperateJson struct {
	Operating int    `json:"operating"`
	Name      string `json:"name"`
	Max       int    `json:"max"`
}

type getChatOperateJson struct {
	Operating int    `json:"operating"`
	Info      string `json:"info"`
}

type getLoginJson struct {
	Nickname string `json:"nickname"`
	Sno      string `json:"sno"`
	Spw      string `json:"spw"`
}

// 下面表示是发送的json
type sendResJson struct {
	Success bool   `json:"success"`
	Info    string `json:"info"`
}

// 显示房间信息
type sendShowRoomJson struct {
	Success bool               `json:"success"`
	Info    string             `json:"info"`
	Rooms   []sendRoomInfoJson `json:"rooms"`
}

type sendRoomInfoJson struct {
	CreateName string `json:"create_name"`
	Info       string `json:"info"`
	Now        int    `json:"now"`
	Max        int    `json:"max"`
}

// 收和发的聊天信息
type sendToStudentChat struct {
	Types    int    `json:"types"`
	Nickname string `json:"nickname"`
	Content  string `json:"content"`
}

type whoInChatRoom struct {
	Nickname string `json:"nickname"`
	Sex      string `json:"sex"`
}

// 发送个人真实的信息
type studentInfoJson struct {
	Name    string `json:"name"`
	Sno     string `json:"sno"`
	College string `json:"college"`
	Class   string `json:"class"`
	Birth   string `json:"birth"`
	Grade   string `json:"grade"`
}
