package services

import (
	"rbac/app/bindfield"
	"rbac/app/models"
	Mysql "rbac/databases"
	"rbac/util"
)

type User struct {
}

func (this *User) Store(bind *bindfield.UserAdd) (err error) {

	userModel := models.User{}
	userModel.Name = bind.Name
	userModel.Username = bind.Username
	err = userModel.SetPassword(bind.Password)
	if err != nil {
		return
	}
	_, err = userModel.Store()
	return
}

func (this *User) GetAll() (users []*models.User, err error) {
	userModel := models.User{}
	users, err = userModel.GetUsersByWhere(nil)
	return
}

func (this *User) Update(id int, bind *bindfield.UserSave) (err error) {
	userModel := models.User{}

	found := Mysql.DB.Not("id", id).Where("username=?", bind.Username).First(&userModel).RecordNotFound()
	if !found {
		err = util.UserNameExistsError
		return
	}
	user, err := userModel.GetUserById(id)
	if err != nil {
		return
	}
	err = user.SetPassword(bind.Password)
	if err != nil {
		return
	}
	user.Username = bind.Username
	user.Name = bind.Name
	result := Mysql.DB.Save(&user)
	if result.Error != nil {
		return result.Error
	}
	return
}

func (this *User) Delete(id int) (err error) {
	userModel := models.User{}
	err = userModel.DeleteById(id)
	if err != nil {
		return
	}
	return
}

func (this *User) Login(bind *bindfield.UserLogin) (map[string]string, error) {
	where := make(map[interface{}]interface{}, 1)
	where["username"] = bind.Username
	userModel := models.User{}
	users, err := userModel.GetUsersByWhere(where)
	if err != nil {
		return nil, err
	}
	if len(users) == 0 {
		return nil,util.UserNameExistsError
	}
	_, err = users[0].CheckPassword(bind.Password)
	if err != nil {
		return nil, err
	}

	tokenString, err := util.GenToken(users[0].Username)
	if err != nil {
		return nil, err
	}
	tokenMap := make(map[string]string, 1)
	tokenMap["token"] = tokenString
	return tokenMap, nil
}
