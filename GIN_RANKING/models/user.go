package models

import "gin_ranking/dao"

type User struct {
	Id       int
	Username string
	//Password string
}

func (User) TableName() string {
	return "user"
}

func GetUserTest(id int) (User, error) {
	var user User
	err := dao.Db.Where("id=?", id).First(&user).Error
	return user, err
}

func AddUser(username string) (int, error) {
	user := User{Username: username}
	result := dao.Db.Create(&user)
	return user.Id, result.Error
}

func UpdateUser(id int, username string) {
	dao.Db.Model(&User{}).Where("id = ?", id).Update("username", username)
}

func DeleteUser(id int) error {
	result := dao.Db.Delete(&User{}, id)
	return result.Error
}
