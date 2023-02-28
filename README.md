# tiktok
Bytedance winter camp project: a backend of simple tiktok
运行
❯ git clone https://github.com/ozline/Todolist-Kratos
❯ cd Todolist-Kratos
❯ go mod tidy
❯ kratos run
方法概览
todolist.v1.Todolist
Add ： 添加一条新的备忘录
ShowAll：查看所有 已完成/未完成/所有事项，支持分页
ShowKey：按关键词查询事项
Delete：删除 一条/所有已经完成/所有待办/所有事项。
Update：将 一条/所有 待办事项设置为已完成（或将已完成设置为待办）。
user.v1.Users
Register：用户注册
Login：用户登录
配置结构 - ./configs/config.yaml
server: # 服务器
  http:
    addr: 0.0.0.0:8000
    timeout: 1s
  grpc:
    addr: 0.0.0.0:9000
    timeout: 1s
data: # 数据
  database: # 数据库
    username:
    password:
    address:
    port: 3306
    dbname:
  redis: # Redis 暂时未用到
    addr: 127.0.0.1:6379
    read_timeout: 0.2s
    write_timeout: 0.2s
auth: # 鉴权
  secret:  # 秘钥
数据库结构
users表
列名	类型
id	用户id（主键）
username	用户名
password	密码
updated_at	信息更新时间
created_at	账号创建时间
deleted_at	软删除时间
phone	手机号
email	邮箱
todolist表
列名	类型
id	待办事项id（主键）
userid	对应主人
status	当前状态 0=未完成 1=已完成
title	标题
message	内容
end_at	Deadline时间
create_at	创建条目时间
update_at	更新时间（状态改变的时间）
统一返回格式
正确返回
{
    "code": "200", //代码
    "msg": "ok", //消息
    "data": {} //获取的数据，或其他
}
错误返回
{
    "code": "400", //代码
    "msg": "Missing Params Data", //错误消息内容
}
User服务
Register 注册用户
请求：

{
    "username" : 用户名,
    "password" : 密码,
    "phone" : 手机号,
    "email": 邮箱
}
Response样例：

{
    "code": "200",
    "msg": "ok",
    "data": {
        "userid": "46", //用户ID
        "username": "ozlinetestfordabian", //用户名
        "created_at": "1653204667053" //创建时间
    }
}
Login 用户登录
请求：

{
    "username": 用户名,
    "password": 密码
}
Response样例：

{
    "code": "200",
    "msg": "ok",
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6IjM4IiwidXNlcm5hbWUiOiJvemxpbmV0ZXN0Iiwic3RhdHVzIjowLCJleHAiOjE2NTMyMDgzMDMsIm5iZiI6MTY1MzIwMTEwM30.-tYj7-k8DrGiBH_iJ9Pfc4k4y_BgAlFMx1oW7o4kU9M", //JWT TOKEN 
    "data": { //用户数据
        "userid": "38",
        "username": "ozlinetest",
        "created_at": "1653048059640"
    }
}
Todolist服务
Metadata
对于这项服务内的所有请求，均需要携带一个Key为,value为当前用户对应jwttoken的metadatax-md-global-token

eg：

KEY：x-md-global-token

VALUE：eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6IjM4IiwidXNlcm5hbWUiOiJvemxpbmV0ZXN0Iiwic3RhdHVzIjowLCJleHAiOjE2NTMyMDM2NDksIm5iZiI6MTY1MzE5NjQ0OX0.pull2jVh7G-LNxw4wl8b7OTA2uJ4RUHkg7MjTHnv_fU

Add 添加待办事项
请求：

{
    "endtime" :"1653197175", //截止时间，13位Timestamp
    "title" : "test", //标题
    "message" : "ozline_啊(Test for Divide Page)" //内容
}
返回：

{
    "code": "200",
    "msg": "ok"
    //data
}
Delete 删除待办事项
请求：

{
    "id": -1
   // id == -3 = 全部（当前用户）
	 // id == -2 = 未完成
	 // id == -1 = 已完成
}
返回：

{
    "code": "200",
    "msg": "ok",
    "count": "2" //删除的条目数
}
Update 更新待办事项
请求：

{
    "id": 1, //需要更新状态的todolistID
    "status": "1" //新的状态
  
  // id == -1 = 全部已完成设为待办
	// id == -2 = 全部待办设为已完成
}
返回：

{
    "code": "200",
    "msg": "ok"
}
ShowAll 显示特定状态的全部待办事项
请求：

{
    "status": "0", //需要获取的状态
    "page": "1", //页码
    "pagesize": "10" //每一页显示的数目
  // status <0 = 获取全部
}
返回：

{
    "code": "200",
    "msg": "ok",
    "data": {
        "list": [
            {
                "title": "test", //标题
                "message": "deadline", //内容
                "end_at": "1653197175", //截止时间
                "create_at": "1653201647164", //创建时间
                "update_at": "1653203962161" //更新时间
            },
          	....(隐藏剩余9条)
        ],
        "total": "12" //搜索到的总数，与pagesize不同
    }
}
ShowKey 关键词查找待办事项
请求：

{
    "key": "ozline", //关键词
    "page": "1", //页码
    "pagesize": "10" //每一页显示的数目
}
返回：

{
    "code": "200",
    "msg": "ok",
    "data": {
        "list": [
            {
                "title": "test", //标题
                "message": "ozline_deadline", //内容
                "end_at": "1653197175", //截止时间
                "create_at": "1653201679897", //创建时间
              	"update_at": "1653203962161"
            }
          	...(隐藏剩余9条)
        ],
        "total": "10" //搜索到的总数，与pagesize不同
    }
}
开发过程中遇到的困难
官方提供的教程狠难懂
中间件的使用（引入jwt）
protoc的使用
metadata的使用
面向结果编程导致项目结构理解歪曲
尚未解决的问题
Todolist.Update方法中count的显示(2022.05.22 Fixed)
错误处理规范化(准备使用Recovery解决)
结束
