package dao

import (
	"sso/bean"
	"sso/mysql"
)

func InsertVerify(verify *bean.Verify) (int64, error) {

	engine := mysql.GetXormEngine()

	sql := "insert into t_verify (`type`,`verify_id`,`code`,`expired_time`) values(?,?,?,?) on duplicate key update code=?,expired_time=?"
	strTime := verify.ExpiredTime.Format("2006-01-02 15:04:05")
	_, err := engine.Exec(sql, verify.Type, verify.VerifyId, verify.Code, strTime, verify.Code, strTime)

	return 1, err
}

func GetVerify(verify *bean.Verify) (bool, error) {
	engine := mysql.GetXormEngine()

	has, err := engine.Get(verify)

	return has, err
}
