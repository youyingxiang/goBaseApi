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
	users, err = models.GetUsersByWhere(nil)
	return
}

func (this *User) Update(id int, bind *bindfield.UserSave) (err error) {
	userModel := models.User{}

	found := Mysql.DB.Not("id", id).Where("username=?", bind.Username).First(&userModel).RecordNotFound()
	// 用户不存在为true 存在为false
	if !found {
		err = util.UserNameExistsError
		return
	}
	user, err := models.GetUserById(id)
	if err != nil {
		return
	}
	err = user.SetPassword(bind.Password)
	if err != nil {
		return
	}
	result := Mysql.DB.Model(&user).Update(models.User{
		Name:bind.Name,
		Username:bind.Username,
		Password:user.Password,
	})
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
	users, err := models.GetUsersByWhere(where)
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

	tokenString, err := util.GenToken(users[0].ID)
	if err != nil {
		return nil, err
	}
	tokenMap := make(map[string]string, 1)
	tokenMap["token"] = tokenString
	return tokenMap, nil
}

func shouldPassThrough()  {

}




