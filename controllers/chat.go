package controllers

import (
	"chat/system"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

func Chat(ctx *gin.Context) {
	// 先判断是否在全局的map中
	jud, ws, opStudent, mt := judInMap(ctx)
	if ws == nil {
		return
	}
	if jud == false {
		if err := ws.Close(); err != nil {
			system.Save.OtherError("close ws error ", err)
		}
		return
	}
	chatRoom := opStudent.room
	// 设置用户在房间
	opStudent.inRoom = true
	// 告诉房间里别人们，有人进入
	loginChat(opStudent.nickname, opStudent.sex, chatRoom)
	// 开启用户层面的发送协程
	go listenStudent(opStudent, ws, mt)
	// 接受用户操作
	for {
		// 读取操作
		mt, message, err := ws.ReadMessage()
		if err != nil {
			break
		}
		operateJson := getChatOperateJson{}
		err = json.Unmarshal(message, &operateJson)
		if err != nil {
			break
		}
		if dealChatOperate(operateJson, chatRoom, opStudent, ws, mt) {
			break
		}
	}
	close(opStudent.getChan)
	if err := ws.Close(); err != nil {
		system.Save.OtherError("close ws error ", err)
	}
}

// 判断用户是否已经在map中了
func judInMap(ctx *gin.Context) (bool, *websocket.Conn, *student, int) {
	loginJson := getLoginJson{}
	ws, err := upGrader.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		system.Save.OtherError("websocket连接失败", err)
		return false, ws, nil, 0
	}
	// 进行登录判断
	mt, loginMessage, err := ws.ReadMessage()
	err = json.Unmarshal(loginMessage, &loginJson)
	if err != nil {
		system.Save.OtherError("json解析出错", err)
		sendInfo(ws, mt, false, err.Error())
		return false, ws, nil, 0
	}
	join, jud := students.Load(loginJson.Nickname)
	opStudent, ok := join.(*student)
	if jud == false || ok == false || opStudent.sno != loginJson.Sno {
		sendInfo(ws, mt, false, "内部错误，请与我们联系")
		return false, ws, nil, 0
	}
	roomName := ctx.Param("name")
	if opStudent.room.roomName == roomName {
		sendInfo(ws, mt, true, "")
		return true, ws, opStudent, mt
	}
	sendInfo(ws, mt, false, "似乎走错房间啦。。。")
	return false, ws, nil, 0
}

// 处理一系列的请求
func dealChatOperate(operate getChatOperateJson, room *room, student *student, ws *websocket.Conn, mt int) bool {
	switch operate.Operating {
	case CHAT:
		return sendChat(operate.Info, student.nickname, room, ws, mt)
	case QUIT:
		return quiteChat(student.nickname, room, ws, mt)
	case WHO:
		return whoInChat(student, room)
	case ME:
		return showMeInChat(student, room, ws, mt)
	}
	sendInfo(ws, mt, false, "出错了，服务器无法解析")
	return false
}

// 每一个用户开一个协程侦听
func listenStudent(opStudent *student, ws *websocket.Conn, mt int) {
	for info := range opStudent.getChan {
		sendJson, _ := json.Marshal(info)
		err := ws.WriteMessage(mt, sendJson)
		if err != nil {
			system.Save.OtherError("发送chat json 时出了错误 ", err)
		}
	}
}
