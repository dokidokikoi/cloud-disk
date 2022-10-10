package helper

import (
	"crypto/md5"
	"crypto/tls"
	"errors"
	"fmt"
	"math/rand"
	"net/http"
	"net/smtp"
	"path"
	"time"

	"cloud-disk/core/define"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/golang-jwt/jwt/v4"
	"github.com/jordan-wright/email"
	uuid "github.com/satori/go.uuid"
)

func Md5(s string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(s)))
}

func GenerateToke(id int, identity, name string, second int64) (string, error) {
	uc := define.UserClaim{
		Id:       id,
		Identity: identity,
		Name:     name,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Second * time.Duration(second)).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, uc)
	tokenString, err := token.SignedString([]byte(define.JwtKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func MailSendCode(mail, code string) (err error) {
	e := email.NewEmail()
	e.From = "haru <harukaze_doki@163.com>"
	e.To = []string{mail}
	e.Subject = "验证码发送测试"
	e.HTML = []byte("你的验证码是：<h1>" + code + "</h1>")
	err = e.SendWithTLS("smtp.163.com:465", smtp.PlainAuth("", "harukaze_doki@163.com", define.MailPassword, "smtp.163.com"),
		&tls.Config{InsecureSkipVerify: true, ServerName: "smtp.163.com"})

	return
}

func RandCode() string {
	nums := "1234567890"
	code := ""
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < define.CodeLength; i++ {
		code += string(nums[rand.Intn(len(nums))])
	}

	return code
}

func GetUUid() string {
	return uuid.NewV4().String()
}

func OssUpload(r *http.Request) (string, error) {
	client, err := oss.New(define.BucketName, define.AccessId, define.Secret)
	if err != nil {
		return "", err
	}

	file, fileHeader, err := r.FormFile("file")
	key := "cloud-disk/" + GetUUid() + path.Ext(fileHeader.Filename)

	// Get bucket
	bucket, err := client.Bucket("harukaze-cloud-disk")
	if err != nil {
		return "", err
	}

	err = bucket.PutObject(key, file)
	if err != nil {
		panic(err)
	}

	return define.OssBucket + "/" + key, nil
}

func AnalyzeToken(token string) (*define.UserClaim, error) {
	uc := new(define.UserClaim)
	claims, err := jwt.ParseWithClaims(token, uc, func(t *jwt.Token) (interface{}, error) {
		return []byte(define.JwtKey), nil
	})
	if err != nil {
		return nil, err
	}
	if !claims.Valid {
		return uc, errors.New("token is invalid")
	}
	return uc, err
}
