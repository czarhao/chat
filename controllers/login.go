package controllers

import (
	"chat/models"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

func login(ctx *gin.Context) (bool, *websocket.Conn, *student) {
	// 进行websocket连接
	loginJson := getLoginJson{}
	ws, err := upGrader.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		return false, ws, nil
	}
	// 进行登录判断
	mt, loginMessage, err := ws.ReadMessage()
	err = json.Unmarshal(loginMessage, &loginJson)
	if err != nil {
		sendInfo(ws, mt, false, err.Error())
		return false, ws, nil
	}
	// 学生登录注册 进行判断
	inDb, studentDb := models.DbContr.StudentExit(loginJson.Sno)
	// 学生未注册
	if !inDb {
		sendInfo(ws, mt, false, "同学似乎还没有注册我们的服务呀，请先去我们的公众号绑定信息")
		return false, ws, nil
	}
	// 学生的密码错误
	if studentDb.Password != loginJson.Spw {
		sendInfo(ws, mt, false, "同学输入的密码错啦！")
		return false, ws, nil
	}
	// 名字出错误
	opStudent, jud := students.Load(loginJson.Nickname)
	// 已经在map中
	if jud == true {
		// 已经在map中
		opStd := opStudent.(*student)
		if !opStd.inRoom {
			sendInfo(ws, mt, true, "")
			return true, ws, opStd
		}
		sendInfo(ws, mt, false, "同学已经在服务中了")
		return false, ws, nil
	}
	// 用户姓名是否为空
	if loginJson.Nickname == "" {
		sendInfo(ws, mt, false, "名字不可为空的")
		return false, ws, nil
	}
	_, returnType := saveToDb(loginJson, studentDb.Sex)
	// 发送成功的json
	cont := sendInfo(ws, mt, true, "")
	if !cont {
		return false, ws, nil
	}
	// 确保以上这一切都正确后返回true
	return true, ws, returnType
}

func sendInfo(ws *websocket.Conn, mt int, jud bool, info string) bool {
	tmpJson := sendResJson{
		Success: jud,
		Info:    info,
	}
	sendJson, _ := json.Marshal(tmpJson)
	err := ws.WriteMessage(mt, sendJson)
	if err != nil {
		return false
	}
	return true
}

// 保存到map中
func saveToDb(logInfo getLoginJson, sex string) (bool, *student) {
	saveType := &student{
		nickname: logInfo.Nickname,
		sno:      logInfo.Sno,
		room:     nil,
		sex:      sex,
		getChan:  make(chan sendToStudentChat, 10),
		inRoom:   false,
	}
	students.Store(logInfo.Nickname, saveType)
	return true, saveType
}
