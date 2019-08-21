# 基于gin和gorilla创建websocket的聊天项目

  个人学习gin和gorilla的一个匿名聊天练手项目，本项目采用了go module,实现以下的功能：

```json
index页面的操作
1.登录服务
{"nickname":"czarhao","sno":"1","spw":"1"}
2.创建房间
{"operating":0,"name":"1","max":10}
3.显示房间
{"operating":1,"name":"","max":0}
4.加入房间
{"operating":2,"name":"1","max":0}

chat页面的操作
1.登录服务
{"nickname":"czarhao","sno":"1","spw":"1"}
2.发送消息
{"operating":1,"nickname":"czarhao","content":"hello"}
3.退出房间
{"operating":2,"nickname":"czarhao","content":""}
4.谁在房间
{"operating":3,"nickname":"czarhao","content":""}
4.发送自己的真现实信息
{"operating":4,"nickname":"czarhao","content":""}
```

前后端使用json交互，部署服务直接：

```shell
docker build -t chat .
docker run -p 8080:8080 -d --name="chat" chat
```



