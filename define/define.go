package define

import (
	"github.com/dgrijalva/jwt-go"
	"os"
)

type UserClaim struct {
	Id       int
	Identity string
	Name     string
	jwt.StandardClaims
}

var JwtKey = "net-disk-key"
var EmailPassword = "QMHFCTQOAJDGGBMG"
var CodeLength = 6

//验证码过器时间
var CodeExpire = 300

//腾讯云COS对象存储配置
var TencentSecretId = os.Getenv("AKIDf7IDtIRM1RebHkvpRloZ3FFyPUhja9QP")
var TencentSecretKey = os.Getenv("j9GdjY5RVIn0vd3BxRfdMnxpPhhPqyoN")
var CosBucketDomain = "https://net-disk-1257185745.cos.ap-beijing.myqcloud.com"

//分页默认参数
var PageSize = 20

//go的格式化时间
var DateTime = "2006-01-02 15:04:05"

//token有效期
var TokenExpire = 3600
var RefreshTokenExpire = 7200
