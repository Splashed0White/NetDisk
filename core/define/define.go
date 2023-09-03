package define

import (
	"github.com/golang-jwt/jwt/v4"
)

type UserClaim struct {
	Id       int
	Identity string
	Name     string
	jwt.RegisteredClaims
}

// Token过期时间
var TokenExpire int = 3600

var JwtKey = "Netdisk-key"
var MailPassword = ""

// 验证码长度
var CodeLength = 6
var CodeExpire = 100

// COS环境变量
var CosID = ""
var COsKey = ""
var Url = "" //存储桶路径

// 分页的默认参数
var PageSize int = 20
