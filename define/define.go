package define

import (
	"github.com/dgrijalva/jwt-go"
)

type UserClaim struct {
	Id       int
	Identity string
	Name     string
	jwt.StandardClaims
}

var JwtKey = "cloud-disk-key"
var EmailPassword = "QMHFCTQOAJDGGBMG"
var CodeLength = 6

// CodeExpire 验证码过器时间
var CodeExpire = 300

// PageSize 分页默认参数
var PageSize = 20

// DateTime go的格式化时间
var DateTime = "2006-01-02 15:04:05"

// TokenExpire token有效期
var TokenExpire = 3600
var RefreshTokenExpire = 7200

//gofastDfs的地址
var GoFastUrl = "http://39.105.57.46:18080/group1"
