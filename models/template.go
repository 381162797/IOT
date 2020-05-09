package models

import (
	"mserver/def"
	"strconv"
)

func AddTemplate(req *Template) (code int) {
	code = def.SUCCESS
	if len(req.Username) < 6 {
		return def.INVALID_PARAMS
	}
	old := &Template{}
	user := &User{}
	db.First(user, User{Username: req.Username})
	if user.ID == 0 {
		return def.USER_NOT_EXIST
	}
	db.First(old, Template{Username: req.Username, TemplateName: req.TemplateName})
	if old.ID > 0 {
		*req = *old
		return def.TEMPLATE_EXIST
	}
	db.Create(req)
	return
}

// GetTemplate 接收用户名，返回所有数据模板
func GetTemplate(username string) (temps []Template, code int) {
	temps = make([]Template, 0)
	if len(username) < 6 {
		return temps, def.INVALID_PARAMS
	}
	db.Preload("DataPoints").Find(&temps, Template{Username: username})
	if len(temps) == 0 {
		return temps, def.TEMPLATE_NOT_EXIST
	}
	return temps, def.SUCCESS
}

// UpdateTemplate 修改数据模板
func UpdateTemplate(req *Template) (code int) {
	old := &Template{}
	db.First(old, Template{Username: req.Username, TemplateName: req.TemplateName})
	if old.ID == 0 {
		return def.TEMPLATE_NOT_EXIST
	}
	db.Model(old).Updates(req)
	req.ID = old.ID
	db.Delete(DataPoint{}, DataPoint{TemplateID: req.ID})
	for _, v := range req.DataPoints {
		//存入ID
		v.TemplateID = req.ID
		db.Create(v)
	}
	return def.SUCCESS
}

//DeleteTemplate 接收templateID删除模板以及对应数据点
func DeleteTemplate(username string, id string) (code int) {
	uID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return def.INVALID_PARAMS
	}
	temp := &Template{
		ID:       uID,
		Username: username,
	}
	old := &Template{}
	db.First(old, temp)
	if old.ID == 0 {
		return def.TEMPLATE_NOT_EXIST
	}
	db.Delete(temp)
	db.Delete(DataPoint{}, DataPoint{TemplateID: uID})
	return def.SUCCESS
}
