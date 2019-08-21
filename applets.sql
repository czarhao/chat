create database applet;
use applet;
CREATE TABLE `student` (
                           `id` int(11) NOT NULL AUTO_INCREMENT comment '学生id',
                           `name` varchar(20) NOT NULL comment '学生姓名',
                           `no` char(8) NOT NULL comment '学生学号',
                           `password` varchar(64) NOT NULL comment '教务网密码',
                           `college` varchar(20) NOT NULL comment '学院',
                           `class` varchar(20) not null comment '班级',
                           `prof` varchar(20) not null comment '专业',
                           `sex` varchar(4) not null comment '性别',
                           `birth` varchar(20) not null comment '生日',
                           `grade` varchar(10) not null comment '年级',
                           PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4;

create table student_course(
                               `id` int(11) NOT NULL AUTO_INCREMENT comment '关联id',
                               `sid` int(11) NOT NULL comment '学生id',
                               `cid` int(11) NOT NULL comment '课程id',
                               PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4;

CREATE TABLE `course` (
                          `id` int(11) NOT NULL AUTO_INCREMENT comment '课程id',
                          `time` int(11) NOT NULL comment '上课时间',
                          `type`  int(11) NOT NULL comment '课程类型',
                          `week_start`  int(11) NOT NULL comment '第几周开始',
                          `week_end`  int(11) NOT NULL comment '第几周结束',
                          `day_week`  int(11) NOT NULL comment '周几上课',
                          `day_start`  int(11) NOT NULL comment '第几节开始',
                          `day_end`  int(11) NOT NULL comment '第几节结束',
                          `course_name` varchar(96) NOT NULL comment '课程名字',
                          `teacher_name` varchar(24) NOT NULL comment '教师名字',
                          `cour_where` varchar(24) NOT NULL comment '上课地点',
                          `jud` int not null comment '单双周',
                          PRIMARY KEY (id)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4;

CREATE TABLE `grade` (
                         `id` int(11) NOT NULL AUTO_INCREMENT comment '成绩id',
                         `year` varchar(24) NOT NULL comment '学年',
                         `semester` int NOT NULL comment '学期',
                         `course_name` varchar(96) NOT NULL comment '课程名字',
                         `credit` double not null comment '学分',
                         `point` double not null comment '绩点',
                         `usually` varchar(8) not null comment '平时成绩',
                         `mid` varchar(8) not null comment '期中成绩',
                         `final` varchar(8) not null comment '期末成绩',
                         `experiment` varchar(8) not null comment '实验成绩',
                         `grade` varchar(8) not null comment '总成绩',
                         `sid` int  not null comment '学生学号',
                         PRIMARY KEY (id)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4;

CREATE TABLE `article` (
                           `id` int(11) NOT NULL AUTO_INCREMENT comment '文章id',
                           `date` varchar(24) NOT NULL comment '创建时间',
                           `title` varchar(100) NOT NULL comment '标题',
                           `head` varchar(100) NOT NULL comment '作者头像',
                           `img` varchar(100) NOT NULL comment '标题图片',
                           `content` text not null comment '内容简介',
                           `author` varchar(100) not null comment '作者',
                           `passage` text not null comment '文章内容',
                           PRIMARY KEY (id)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4;
