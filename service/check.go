package service

import (
	"log"
	"regexp"
	ssoerror "sso/error"
)

func CheckEmail(email string) error {
	if len(email) <= 0 {
		return ssoerror.ErrorIllegalParams
	}

	pattern := `\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*`
	regexp.MustCompile(pattern)
	match, err := regexp.MatchString(pattern, email)
	if err != nil {
		log.Println(err.Error())
		return ssoerror.ErrorInternalServerError
	}
	if !match {
		return ssoerror.ErrorRegisterEmailFormat
	}

	return nil
}

/*
1.用户名为4-20个字符（英文字母、数字）组成

2.密码为6-20个字符（英文字母、数字）组成

注：用户名均不能单独以数字命名，且字母与数字组合的用户名不能以数字开头，不便之处，请用户谅解！
*/
func CheckUserName(userName string) error {
	if len(userName) <= 0 {
		return ssoerror.ErrorIllegalParams
	}
	pattern := `^[A-Za-z]\w{3,19}$`
	regexp.MustCompile(pattern)
	match, err := regexp.MatchString(pattern, userName)
	if err != nil {
		log.Println(err.Error())
		return ssoerror.ErrorInternalServerError
	}
	if !match {
		return ssoerror.ErrorUserNameFormatError
	}
	return nil
}

func CheckPassword(password string) error {
	if len(password) <= 0 {
		return ssoerror.ErrorIllegalParams
	}

	//ascii字符集
	pattern := `^[:ascii:]{8,20}$`
	regexp.MustCompile(pattern)
	match, err := regexp.MatchString(pattern, password)
	if err != nil {
		log.Println(err.Error())
		return ssoerror.ErrorInternalServerError
	}
	if !match {
		log.Println("密码不匹配")
		return ssoerror.ErrorPasswordFormatError
	}

	//不能是全数字
	pattern = `^\d+$`
	regexp.MustCompile(pattern)
	match, err = regexp.MatchString(pattern, password)
	if err != nil {
		log.Println(err.Error())
		return ssoerror.ErrorInternalServerError
	}
	if !match {
		log.Println("密码不匹配")
		return ssoerror.ErrorPasswordFormatError
	}

	return nil
}
