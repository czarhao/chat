package models

import "chat/system"

// 学生模型
type Student struct {
	Id       int    `xorm:"not null pk autoincr comment('学生id') INT(11)"`
	Name     string `xorm:"not null comment('学生姓名') VARCHAR(20)"`
	No       string `xorm:"not null comment('学生学号') CHAR(8)"`
	Password string `xorm:"not null comment('教务网密码') VARCHAR(64)"`
	College  string `xorm:"not null comment('学院') VARCHAR(20)"`
	Class    string `xorm:"not null comment('班级') VARCHAR(20)"`
	Prof     string `xorm:"not null comment('专业') VARCHAR(20)"`
	Sex      string `xorm:"not null comment('性别') VARCHAR(4)"`
	Birth    string `xorm:"not null comment('生日') VARCHAR(20)"`
	Grade    string `xorm:"not null comment('年级') VARCHAR(10)"`
}

// 判断学号是否存在，返回学生信息
func (db *DbController) StudentExit(sno string) (bool, *Student) {
	studentInfo := &Student{No: sno}
	if has, err := db.eng.Get(studentInfo); err != nil {
		system.Save.DbError("StudentExit()", err)
		return false, nil
	} else {
		return has, studentInfo
	}
}
