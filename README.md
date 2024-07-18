## 启动

- 配置数据库

```js
// config/db.go
// 在数据库中，建立一个blog数据库
// root:root意为：账号：密码
const (
	Mysql = "root:root@tcp(127.0.0.1:3306)/blog?charset=utf8"
)
```

```js
// 启动
go run main.go
```

## token 字符串

- 生成一个随机字符串
- 定一个 util 用于处理函数

```go
// /util/function.go
var seededRand *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano())) // 随机数
var chars = "0123456789qwertyuiopasdfghjklzxcvbnmQWERTYUIOPASDFGHJKLZXCVBNM" // 字符串

// 根据length生成一个length长的随机字符串
func CreateNonceStr(length int) string {
	b := make([]byte, length) // 字符数组
	for i := range b {
		b[i] = chars[seededRand.Intn(len(chars))] // 随机选着一个存入字符数组
	}
	return string(b)
}

```

## 中间件

- 中间件，用于处理路由中间的事情，例如处理 token
- 链接：https://www.paperflying.com/2023/10/30/go%E8%AF%AD%E8%A8%80%E5%AD%A6%E4%B9%A0/%E8%AF%AD%E8%A8%80%E5%9F%BA%E7%A1%80/%E7%AC%AC17%E7%AB%A0%EF%BC%9AGo%20web%E6%A1%86%E6%9E%B6%E2%80%94%E2%80%94Gin%E4%B8%AD%E9%97%B4%E4%BB%B6%E5%92%8C%E8%B7%AF%E7%94%B1/#Go-web%E6%A1%86%E6%9E%B6%E2%80%94%E2%80%94Gin%E4%B8%AD%E9%97%B4%E4%BB%B6%E4%B8%8E%E8%B7%AF%E7%94%B1

```go
// router.go
// 接口中间处理请求
user.POST("/logout", ValidateToken, controllers.UserController{}.LogOut)
```

- 使用 next() 处理没问题，next 继续，否则抛出错误

```go
// router.go
func ValidateToken(ctx *gin.Context) {
	tokenStr := ctx.GetHeader("token") // 获取请求头，token
	if tokenStr == "" {
		controllers.ReturnError(ctx, err_code.ErrInvalidRequest, "user not login")
		return
	}

	token, err := modules.GetTokenInfo(tokenStr) // 获取数据库中请求头信息，用于判断是否过期
	if err != nil {
		controllers.ReturnError(ctx, err_code.ErrInvalidRequest, err.Error())
		return
	}

	if token.Expire < time.Now().Unix() {
		controllers.ReturnError(ctx, err_code.ErrInvalidToken, "token expired")
		return
	}

	if token.State == 0 {
		controllers.ReturnError(ctx, err_code.ErrInvalidToken, "token expired")
		return
	}
	ctx.Next()
}

```
