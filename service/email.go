package service

import (
	"fmt"
	//"log"
	"math/rand"
	"net/smtp"
	"sso/bean"
	"sso/dao"
	ssoerror "sso/error"
	"strings"
	"time"
)

const (
	HOST        = "smtp.163.com"
	SERVER_ADDR = "smtp.163.com:25"
	USER        = "shenqi329@163.com"
	PASSWORD    = "acnjmvn329"
)

type (
	Email struct {
		to      string "to"
		subject string "subject"
		msg     string "msg"
	}
)

func UserRegisetrEMailVerifyCode(user *bean.User) error {

	if err := CheckUserName(user.UserName); err != nil {
		return err
	}
	if err := CheckEmail(user.Email); err != nil {
		return err
	}

	has, err := dao.GetUser(&bean.User{UserName: user.UserName})

	if err != nil {
		return ssoerror.ErrorInternalServerError
	}
	if has {
		return ssoerror.ErrorRegisterUserExist
	}

	has, err = dao.GetUser(&bean.User{Email: user.Email})
	if err != nil {
		return ssoerror.ErrorInternalServerError
	}
	if has {
		return ssoerror.ErrorRegisterEmailInUse
	}

	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	emailCode := fmt.Sprintf("%06d", r.Intn(999999))
	expiredTime := time.Now().Add(10 * time.Minute)
	emailBean := bean.Email{
		UserName:    user.UserName,
		Email:       user.Email,
		Code:        emailCode,
		ExpiredTime: &expiredTime,
	}
	count, err := dao.InsertEmail(&emailBean)
	if err != nil {
		return ssoerror.ErrorInternalServerError
	}
	if count <= 0 {
		return ssoerror.ErrorInternalServerError
	}

	emailToSend := newEmail(user.Email, "验证码", fmt.Sprintf("注册easynote的验证码为:%s,10分钟内有效", emailCode))

	go sendEmail(emailToSend)

	return nil
}

func newEmail(to, subject, msg string) *Email {
	return &Email{to: to, subject: subject, msg: msg}
}

func sendEmail(email *Email) {
	auth := smtp.PlainAuth("", USER, PASSWORD, HOST)
	sendTo := strings.Split(email.to, ";")
	done := make(chan error, 1024)

	go func() {
		defer close(done)

		for _, v := range sendTo {
			str := strings.Replace("From: "+USER+"~To: "+v+"~Subject: "+email.subject+"~~", "~", "\r\n", -1) + email.msg
			err := smtp.SendMail(
				SERVER_ADDR,
				auth,
				USER,
				[]string{v},
				[]byte(str),
			)
			done <- err
		}
	}()

	for i := 0; i < len(sendTo); i++ {
		<-done
	}
	return
}
