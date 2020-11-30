package service

import (
	"gin-server-cli/global"
	"gin-server-cli/model"
)

//测试用
func GetCommunityById(id int64) (model.Community, error) {
	var commmunty model.Community
	err := global.DbEngine.Raw("SELECT * FROM `community` WHERE id = ?", id).Scan(&commmunty).Error
	return commmunty, err
}
