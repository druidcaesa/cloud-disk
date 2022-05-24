package utils

import (
	"cloud-disk/define"
	"context"
	"crypto/md5"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	uuid "github.com/satori/go.uuid"
	"log"
	"net/http"
	"path"
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

func UploadFileToMinio(r *http.Request) (string, error) {
	c := context.Background()
	client, err := minio.New(define.Endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(define.AccessKeyID, define.SecretAccessKey, ""),
		Secure: false,
	})
	if err != nil {
		log.Fatalln("创建Minio链接出现异常", err)
		return "", err
	}
	//校验存储桶是否存在
	exists, err := client.BucketExists(c, define.BucketName)
	if err != nil {
		log.Fatalln("查询存储桶状态异常", err)
		return "", err
	}
	//存储桶不存在进行创建
	if !exists {
		err = client.MakeBucket(c, define.BucketName, minio.MakeBucketOptions{Region: define.BucketLocation, ObjectLocking: false})
		if err != nil {
			log.Fatalln("创建存储桶异常", err)
			return "", err
		}
		//设置存储桶为公读写模式
		err := client.SetBucketPolicy(c, define.BucketName, fmt.Sprintf(define.BucketPolicy, define.BucketName, define.BucketName))
		if err != nil {
			log.Fatalln("修改存储桶权限异常", err)
			return "", err
		}
	}
	//进行文件上传
	formFile, header, err := r.FormFile("file")
	fileName := GetUUID() + path.Ext(header.Filename)
	_, err = client.PutObject(c, define.BucketName, fileName, formFile, header.Size, minio.PutObjectOptions{})
	if err != nil {
		return "", err
	}
	return define.Endpoint + "/" + define.BucketName + "/" + fileName, nil
}
