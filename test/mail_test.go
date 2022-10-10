package test

import (
	"crypto/tls"
	"net/smtp"
	"testing"

	"cloud-disk/core/define"

	"github.com/jordan-wright/email"
)

func TestSendMail(t *testing.T) {
	e := email.NewEmail()
	e.From = "haru <harukaze_doki@163.com>"
	e.To = []string{"2397351356@qq.com"}
	e.Subject = "验证码发送测试"
	e.HTML = []byte("你的验证码是：<h1>123456</h1>")
	err := e.SendWithTLS("smtp.163.com:465", smtp.PlainAuth("", "harukaze_doki@163.com", define.MailPassword, "smtp.163.com"),
		&tls.Config{InsecureSkipVerify: true, ServerName: "smtp.163.com"})
	if err != nil {
		t.Fatal(err)
	}
}
