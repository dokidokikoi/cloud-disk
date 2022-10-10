package models

import (
	"cloud-disk/core/helper"
	"time"
)

type UserBasic struct {
	Id        int
	Identity  string
	Name      string
	Password  string
	Email     string
	CreatedAt time.Time `xorm:"created"`
	UpdatedAt time.Time `xorm:"updated"`
	DeletedAt time.Time `xorm:"deleted"`
}

func (table UserBasic) TableName() string {
	return "user_basic"
}

func FindUserByNameAndPwd(name, pwd string) *UserBasic {
	user := new(UserBasic)
	has, err := engine.Where("name = ? AND password = ?", name, helper.Md5(pwd)).Get(user)
	if err != nil || !has {
		return nil
	}

	return user
}

func FindUserByIdentity(identity string) *UserBasic {
	user := new(UserBasic)
	has, err := engine.Where("identity = ?", identity).Get(user)
	if err != nil || !has {
		return nil
	}

	return user
}

func CheckEmailExist(email string) (int64, error) {
	return engine.Where("email = ?", email).Count(new(UserBasic))
}

func CheckUserExist(name string) (int64, error) {
	return engine.Where("name = ?", name).Count(new(UserBasic))
}

func (u *UserBasic) Save() error {
	_, err := engine.Insert(u)
	return err
}
