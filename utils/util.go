package utils

import (
	"cloud-disk/define"
	"crypto/md5"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	uuid "github.com/satori/go.uuid"
	"time"
)

// AnalyzeToken 解析token
func AnalyzeToken(token string) (*define.UserClaim, error) {
	uc := new(define.UserClaim)
	claims, err := jwt.ParseWithClaims(token, uc, func(token *jwt.Token) (interface{}, error) {
		return []byte(define.JwtKey), nil
	})
	if err != nil {
		return nil, err
	}
	if !claims.Valid {
		return uc, errors.New("登录已失效")
	}
	return uc, err
}

// GenerateToken 生成token
func GenerateToken(id int, identity string, userName string, second int) (error, string) {
	claim := define.UserClaim{
		Id:       id,
		Identity: identity,
		Name:     userName,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Second * time.Duration(second)).Unix(),
		}}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	signedString, err := token.SignedString([]byte(define.JwtKey))
	if err != nil {
		return err, ""
	}
	return nil, signedString
}

// Md5ToString 生成md5
func Md5ToString(s string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(s)))
}

// GetUUID 获取UUID
func GetUUID() string {
	return uuid.NewV4().String()
}
