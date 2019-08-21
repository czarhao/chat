package controllers

// 相关的学生信息type
type student struct {
	// 匿名的名字
	nickname string
	// 学号
	sno string
	// 所在的房间
	room *room
	// 性别
	sex string
	// 接收消息chan
	getChan chan sendToStudentChat
	// 是否在房间中
	inRoom bool
}

// 相关的房间信息type
type room struct {
	// 话题的名字
	roomName string
	// 房间创建者的名字
	createName string
	// 最大的学生数
	max int
	// 房间中的学生nickname
	students []*student
	// 接受到的消息
	getChan chan sendToStudentChat
}
