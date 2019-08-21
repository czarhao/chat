package controllers

import (
	"chat/models"
	"encoding/json"
	"github.com/gorilla/websocket"
)

func loginChat(nickname string, sex string, inRoom *room) {
	sendChat := sendToStudentChat{
		Types:    LOGIN,
		Nickname: nickname,
		Content:  sex,
	}
	inRoom.getChan <- sendChat
}

func sendChat(chatInfo string, nickname string, inRoom *room, ws *websocket.Conn, mt int) bool {
	sendChat := sendToStudentChat{
		Types:    CHAT,
		Nickname: nickname,
		Content:  chatInfo,
	}
	inRoom.getChan <- sendChat
	sendInfo(ws, mt, true, "")
	return false
}

func quiteChat(nickname string, inRoom *room, ws *websocket.Conn, mt int) bool {
	sendChat := sendToStudentChat{
		Types:    QUIT,
		Nickname: nickname,
		Content:  "",
	}
	// 退出房间
	for k, v := range inRoom.students {
		if v.nickname == nickname {
			v.room = nil
			v.inRoom = false
			inRoom.students = append(inRoom.students[:k], inRoom.students[k+1:]...)
			break
		}
	}
	inRoom.getChan <- sendChat
	// 房间没有人关闭
	if len(inRoom.students) == 0 {
		close(inRoom.getChan)
	}
	sendInfo(ws, mt, true, "")
	return true
}

func whoInChat(studentInfo *student, roomInfo *room) bool {
	var userList []whoInChatRoom
	for _, studentInfo := range roomInfo.students {
		userList = append(userList, whoInChatRoom{
			Nickname: studentInfo.nickname,
			Sex:      studentInfo.sex,
		})
	}
	sendJson, _ := json.Marshal(userList)
	sendChat := sendToStudentChat{
		Types:    WHO,
		Nickname: "",
		Content:  string(sendJson),
	}
	studentInfo.getChan <- sendChat
	return false
}

func showMeInChat(studentInfo *student, inRoom *room, ws *websocket.Conn, mt int) bool {
	_, studentInDb := models.DbContr.StudentExit(studentInfo.sno)
	sendJson := studentInfoJson{
		Name:    studentInDb.Name,
		Sno:     studentInDb.No,
		College: studentInDb.College,
		Class:   studentInDb.Class,
		Birth:   studentInDb.Birth,
		Grade:   studentInDb.Grade,
	}
	content, _ := json.Marshal(sendJson)
	inRoom.getChan <- sendToStudentChat{
		Types:    ME,
		Nickname: studentInfo.nickname,
		Content:  string(content),
	}
	sendInfo(ws, mt, true, "")
	return false
}
