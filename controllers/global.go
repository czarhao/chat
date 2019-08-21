package controllers

import "sync"

var (
	// 用于存放用户的信息
	// 对应学号的用户信息
	students sync.Map

	// 用于存放房间的信息
	// 对应聊天话题的信息
	rooms sync.Map
)
