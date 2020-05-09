package models

import (
	"mserver/def"
	"time"
)

func CheckAuth(username, password string) bool {
	user := &User{}
	db.Select("id").Where("username = ? AND password = ?", username, password).First(user)
	if user.ID > 0 {
		return true
	}
	return false
}
func QueryUser(username, password string) (user *User) {
	user = &User{}
	if len(username)+len(password) < 12 {
		return
	}
	db.Where("username = ? AND password = ?", username, password).First(user)
	user.Password = ""
	t := time.Now().Format("2006-01-02 15:04:05")
	db.Model(&User{ID: user.ID}).UpdateColumn("last_login", t)
	return user
}
func InsertUser(username, password string) (user *User, code int) {
	user = &User{}
	//校验
	if len(username) < 6 || len(password) < 6 {
		return user, def.INVALID_PARAMS
	}
	db.First(user, &User{Username: username})
	if user.ID > 0 {
		user.Password = ""
		return user, def.USER_EXIST
	}
	user.Username = username
	user.Password = password
	db.Create(user)
	return user, def.SUCCESS
}

func DeleteUser(userID uint64, password string) (code int) {
	if len(password) < 6 {
		return def.INVALID_PARAMS
	}
	user := &User{
		ID: userID,
	}
	db.Where(user).First(user)
	if user.Password == password {
		db.Delete(user)
		return def.SUCCESS
	}
	return def.ERROR_AUTH
}

func UpdateUser(userID uint64, old, new string) (code int) {
	user := &User{}
	db.First(user, userID)
	if len(old)+len(new) < 12 {
		return def.INVALID_PARAMS
	}
	if old == new {
		return def.SUCCESS
	}
	if user.Password == old {
		db.Model(user).Update("password", new)
		return def.SUCCESS
	}
	return def.ERROR_AUTH
}
