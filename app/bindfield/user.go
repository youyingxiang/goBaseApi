package bindfield

import (
	"rbac/app/models"
	"rbac/util"
)

type UserAdd struct {
	Username   string `form:"username" binding:"required,min=6,max=20"`
	Name       string `form:"name" binding:"required"`
	Password   string `form:"password" binding:"required,min=6,max=20"`
	RePassword string `form:"repassword" binding:"required,min=6,max=20"`
}

func (this *UserAdd) UserNameUnique() (err error) {
	where := make(map[interface{}]interface{}, 1)
	where["username"] = this.Username
	users, err := models.GetUsersByWhere(where)
	if err != nil {
		return
	}
	if len(users) != 0 {
		err = util.UserNameExistsError
		return
	}
	return
}

func (this *UserAdd) CheckPasswordInput() (err error) {
	if this.Password != this.RePassword {
		err = util.InputPasswordCheckError
	}
	return
}
func (this *UserAdd) Check() (err error) {
	err = this.CheckPasswordInput()
	if err != nil {
		return
	}
	err = this.UserNameUnique()
	if err != nil {
		return
	}
	return
}

// 用户保存
type UserSave struct {
	UserAdd
}

func (this *UserSave) Check() (err error) {
	err = this.CheckPasswordInput()
	if err != nil {
		return
	}
	return
}





type UserLogin struct {
	Username string `form:"username" binding:"required,min=6,max=20"`
	Password string `form:"password" binding:"required,min=6,max=20"`
}
