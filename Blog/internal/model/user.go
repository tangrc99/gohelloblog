package model

import "github.com/tangrc99/gohelloblog/global"

type User struct { // gorm 默认这里数据结构的名称需要和表名称对应，例如表名称为 tables，那么这里要为 table
	UserName string `gorm:"primary_key" json:"user_name"`
	Password string `json:"password"`
	IsAdmin  bool   `json:"is_admin"`
}

func (usr *User) IsUsrPwdMatch() bool {

	user := User{}
	global.MySQL.Select("password").Where("user_name = ?", usr.UserName).Take(&user)
	return user.Password == usr.Password
}

func (usr *User) CreateNewUser() error {

	return global.MySQL.Create(usr).Error
}

func (usr *User) IsAdministrator() bool {
	global.MySQL.Select("is_admin").Where("user_name = ?", usr.UserName).Take(usr)

	return usr.IsAdmin
}
