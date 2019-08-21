package controllers

import (
	"chat/system"
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
)

func createRoom(roomName string, max int, opStudent *student, ws *websocket.Conn, mt int) bool {
	_, jud := rooms.Load(roomName)
	if jud == true {
		sendInfo(ws, mt, false, "已经有了相同的话题啦")
		return false
	}
	if max <= 1 {
		sendInfo(ws, mt, false, "设置的房间人数太少啦")
		return false
	}
	newRoom := &room{
		roomName:   roomName,
		createName: opStudent.nickname,
		max:        max,
		students:   []*student{opStudent},
		getChan:    make(chan sendToStudentChat, 10),
	}
	opStudent.room = newRoom
	go listenRoom(newRoom)
	rooms.Store(roomName, newRoom)
	sendInfo(ws, mt, true, "")
	return true
}

func listenRoom(listenRoom *room) {
	for info := range listenRoom.getChan {
		fmt.Println("发送消息给各个用户", info)
		for _, opStudent := range listenRoom.students {
			// 确定用户已经进入了房间
			if opStudent.inRoom {
				opStudent.getChan <- info
			}
		}
	}
	rooms.Delete(listenRoom.roomName)
	fmt.Println("room end")
}

func showRoom(ws *websocket.Conn, mt int) bool {
	var roomList []sendRoomInfoJson
	rooms.Range(func(key, value interface{}) bool {
		tmpRoom := value.(*room)
		roomList = append(roomList, sendRoomInfoJson{
			CreateName: tmpRoom.createName,
			Info:       tmpRoom.roomName,
			Now:        len(tmpRoom.students),
			Max:        tmpRoom.max,
		})
		return true
	})
	tmpJson := sendShowRoomJson{
		Success: true,
		Info:    "",
		Rooms:   roomList,
	}
	sendJson, _ := json.Marshal(tmpJson)
	err := ws.WriteMessage(mt, sendJson)
	if err != nil {
		system.Save.OtherError("发送show Room出了错误", err)
	}
	return false
}

func joinRoom(roomName string, opStudent *student, ws *websocket.Conn, mt int) bool {
	join, jud := rooms.Load(roomName)
	if !jud {
		sendInfo(ws, mt, false, "当前话题不存在")
		return false
	}
	tmpRoom, ok := join.(*room)
	if !ok {
		sendInfo(ws, mt, false, "内部错误，请联系我们")
		return false
	}
	if len(tmpRoom.students) >= tmpRoom.max {
		sendInfo(ws, mt, false, "房间人满啦，换一个吧")
		return false
	}
	opStudent.room = tmpRoom
	tmpRoom.students = append(tmpRoom.students, opStudent)
	sendInfo(ws, mt, true, "")
	return true
}
