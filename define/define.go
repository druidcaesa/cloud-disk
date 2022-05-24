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

// Endpoint minio对象存储相关数据
var Endpoint = "39.105.57.46:9000"
var AccessKeyID = "cloud-disk"
var SecretAccessKey = "cloud-disk"
var BucketName = "cloud-disk"
var BucketLocation = "beijing"

// BucketPolicy 设置存储桶权限
var BucketPolicy = "{\"Version\":\"2012-10-17\",\"Statement\":[{\"Effect\":\"Allow\",\"Principal\":{\"AWS\":[\"*\"]},\"Action\":[\"s3:GetBucketLocation\",\"s3:ListBucket\",\"s3:ListBucketMultipartUploads\"],\"Resource\":[\"arn:aws:s3:::%s\"]},{\"Effect\":\"Allow\",\"Principal\":{\"AWS\":[\"*\"]},\"Action\":[\"s3:AbortMultipartUpload\",\"s3:DeleteObject\",\"s3:GetObject\",\"s3:ListMultipartUploadParts\",\"s3:PutObject\"],\"Resource\":[\"arn:aws:s3:::%s/*\"]}]}"
