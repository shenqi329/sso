package service

import (
	"fmt"
	"math/rand"
	"time"
)

func newVerifyCode() string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	verifyCode := fmt.Sprintf("%06d", r.Intn(999999))

	return verifyCode
}
