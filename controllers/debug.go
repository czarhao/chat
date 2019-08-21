package controllers

import "fmt"

func showAll() {
	fmt.Println("createName		getChan 	max 	roomName 	students")
	rooms.Range(func(key, value interface{}) bool {
		tmpRoom := value.(*room)
		fmt.Println(tmpRoom.createName, tmpRoom.getChan, tmpRoom.max, tmpRoom.roomName, tmpRoom.students)
		return true
	})
	fmt.Println("getChan	room	inRoom	nickname	sex		sno")
	students.Range(func(key, value interface{}) bool {
		tmpStudent := value.(*student)
		fmt.Println(tmpStudent.getChan, tmpStudent.room, tmpStudent.inRoom, tmpStudent.nickname, tmpStudent.sex, tmpStudent.sno)
		return true
	})
}
