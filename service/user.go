package service

import (
	"errors"
	"log"
	"sso/bean"
	"sso/dao"
	"sso/uuid"
	"time"
)

var (
	ErrorServiceParams   = errors.New("参数错误")
	ErrorServiceExist    = errors.New("资源已存在")
	ErrorServiceDBError  = errors.New("数据库异常")
	ErrorServiceNotFound = errors.New("数据未找到")
)

func GetUserBean(name string) (*bean.User, error) {

	user, err := dao.GetUserByName(name)

	if err != nil {
		if err == dao.ErrorDaoNotFound {
			return nil, ErrorServiceNotFound //说明没有查询到，用户不存在
		}
		return nil, ErrorServiceDBError //查询发生错误
	}
	return user, nil //查询
}

func GetUserBeanById(id string) (*bean.User, error) {

	user, err := dao.GetUserById(id)

	if err != nil {
		if err == dao.ErrorDaoNotFound {
			return nil, ErrorServiceNotFound //说明没有查询到，用户不存在
		}
		return nil, ErrorServiceDBError //查询发生错误
	}
	return user, nil //查询
}

func GetTokenBean(token string) (*bean.Token, error) {

	tokenBean, err := dao.GetTokenByToken(token)

	if err != nil {
		if err == dao.ErrorDaoNotFound {
			return nil, ErrorServiceNotFound
		}
		return nil, ErrorServiceDBError
	}

	return tokenBean, nil
}

func UserLogin(user *bean.User) (*bean.Token, *bean.User, error) {

	if len(user.Name) == 0 || len(user.Password) == 0 {
		return nil, nil, ErrorServiceParams
	}

	userBean, err := dao.GetUserByName(user.Name)

	if err != nil {
		return nil, nil, err
	}

	if userBean == nil {
		return nil, nil, ErrorServiceNotFound
	}

	token := new(bean.Token)
	token.UserId = userBean.ID
	token.Token = uuid.Rand().Hex()
	token.CreateTime = time.Now()

	dao.InsertToken(token)

	if err != nil {
		return nil, nil, ErrorServiceDBError
	}

	return token, userBean, nil
}

func UserLogout(token string) error {

	if len(token) == 0 {
		return ErrorServiceParams
	}

	err := dao.RemoveTokenByToken(token)

	if err != nil {
		if err == dao.ErrorDaoNotFound {
			return ErrorServiceNotFound
		}
		return ErrorServiceDBError
	}

	return nil
}

func UserRegister(user *bean.User) (*bean.User, error) {

	if len(user.Name) == 0 || len(user.Password) == 0 {
		return nil, ErrorServiceParams
	}
	result, err := dao.GetUserByName(user.Name)
	if err != nil && err != dao.ErrorDaoNotFound {
		return nil, ErrorServiceNotFound
	}
	if result != nil {
		log.Println("用户已经存在,username = ", user.Name)
		return nil, ErrorServiceExist
	}

	user.CreateTime = time.Now()
	user.UpdateTime = time.Now()

	err = dao.InsertUser(user)

	if err != nil {
		return nil, ErrorServiceDBError
	}

	return user, nil
}
