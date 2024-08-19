## Gin+Gorm
可以看tag init
- gin小案例
- router和路由文件封装
- Json返回并统一响应格式
- 控制器controller
- 获取请求参数并和结构体绑定
- 通过defer recover 捕获异常
- 自定义logger中间件收集日志
- 用gorm实现数据库操作
- crud实现与封装
## 朋友圈求投票
实现活动规则，用户注册登录，参赛选手列表，参赛选手详情，比赛信息，比赛成绩，排行榜，投票。
- 用户注册登录以及session的使用
- mysql投票功能
- mysql排行榜功能
- Redis配置
- Redis Sorted Sets优化排行榜功能
- 部署上线
## 遇到的小问题
#### 时间格式
>  sql: Scan error on column index 3, name "add_time": unsupported Scan, storing driver.Value type []uint8 into type *time.Time
- 解决
```go
	AddTime    time.Time `json:"add_time"`
	UpdateTime time.Time `json:"update_time"`

    Mysqldb = "root:root@tcp(127.0.0.1:3306)/ranking?charset=utf8&parseTime=True&loc=Local"
```
#### 加密密码
EncryMd5 函数通过创建 MD5 哈希对象，将字符串转换为字节数组并写入到该对象中，计算 MD5 哈希值，并将结果编码成十六进制字符串返回，实现了对输入字符串进行 MD5 加密的功能。
```go
func EncryMd5(s string) string {
	// 初始化一个MD5哈希对象
	ctx := md5.New()
	// 将字符串写入ctx 对象的内部缓冲区
	ctx.Write([]byte(s))
	// 计算MD5哈希值, 并将结果转换为十六进制编码的字符串
	return hex.EncodeToString(ctx.Sum(nil))
}
```
#### session

```go
	store, _ := sessions_redis.NewStore(10, "tcp", config.RedisAddress, "", []byte("secret"))
	r.Use(sessions.Sessions("mysession", store))
```
#### 删除数据后id继续递增

```sql
DELETE FROM your_table;
ALTER TABLE your_table AUTO_INCREMENT = 1; 
```

#### 实现数据库的某一列更新加一
```go
dao.Db.Model(&player).Where("id =?", id).UpdateColumn("score", gorm.Expr("score +?", 1))
```

#### score不改变
因为结构体字段名首字母小写了......