package controllers

import (
	"chat/system"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

// 登录首页，以及整个非聊天页面的控制
func Index(ctx *gin.Context) {
	cont, ws, student := login(ctx)
	showAll()
	if student == nil {
		if ws == nil {
			return
		}
		if err := ws.Close(); err != nil {
			system.Save.OtherError("close ws error ", err)
		}
		return
	}
	if !cont {
		if err := ws.Close(); err != nil {
			system.Save.OtherError("close ws error ", err)
		}
		return
	}
	for {
		// 读取操作
		mt, loginMessage, err := ws.ReadMessage()
		if err != nil {
			break
		}
		operateJson := getIndexOperateJson{}
		err = json.Unmarshal(loginMessage, &operateJson)
		if err != nil {
			break
		}
		if dealIndexOperate(operateJson, student, ws, mt) {
			break
		}
		showAll()
	}
	if err := ws.Close(); err != nil {
		system.Save.OtherError("close ws error ", err)
	}
}

// 处理操作，类似路由
func dealIndexOperate(operate getIndexOperateJson, student *student, ws *websocket.Conn, mt int) bool {
	switch operate.Operating {
	case CREATE:
		return createRoom(operate.Name, operate.Max, student, ws, mt)
	case SHOW:
		return showRoom(ws, mt)
	case JOIN:
		return joinRoom(operate.Name, student, ws, mt)
	}
	return false
}
