package define

import (
	"github.com/golang-jwt/jwt/v4"
)

type UserClaim struct {
	Id       int
	Identity string
	Name     string
	jwt.StandardClaims
}

var JwtKey = "cloud-disk-key"
var MailPassword = "UPZSBDEQINCFXWXI"
var CodeLength = 6
var CodeExpire = 60

var OssBucket = "harukaze-cloud-disk.oss-cn-shenzhen.aliyuncs.com"
var BucketName = "oss-cn-shenzhen.aliyuncs.com"
var AccessId = ""
var Secret = ""

var PageSize = 10
var Datetime = "2006-01-02 15:04:05"
var TokenExpire = 3600
var RefreshTokenExpire = 7200
