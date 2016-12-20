package dao

import (
	"sso/bean"
	"sso/mysql"
)

func InsertVerify(verify *bean.Verify) (int64, error) {

	engine := mysql.GetXormEngine()

	sql := "insert into t_verify (t_verify_verify_type,t_verify_verify_id,t_verify_code,t_verify_expired_time) values(?,?,?,?) on duplicate key update t_verify_code=?,t_verify_expired_time=?"

	_, err := engine.Exec(sql, verify.Type, verify.VerifyId, verify.VerifyCode, verify.ExpiredTime, verify.VerifyCode, verify.ExpiredTime)

	return 1, err
}

func GetVerify(verify *bean.Verify) (bool, error) {
	engine := mysql.GetXormEngine()

	has, err := engine.Get(verify)

	return has, err
}
