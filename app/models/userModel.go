package models

import (
	"fmt"
	"rbac/databases"
	"rbac/util"
)

type User struct {
	Base
	Username string `json:"username"`
	Name     string `json:"name"`
	Password string `json:"password"`
	Roles    []Role `gorm:"many2many:admin_role_users"`
}

func (User) TableName() string {
	return "admin_users"
}

func (this *User) Store() (id int, err error) {
	result := Mysql.DB.Create(&this)
	if result.Error != nil {
		err = result.Error
		return
	}
	id = this.ID
	return
}

func (this *User) GetUsersByWhere(wheres map[interface{}]interface{}) (users []*User, err error) {
	db := Mysql.DB
	if wheres != nil {
		for i, v := range wheres {
			db = Mysql.DB.Where(i.(string)+"=?", v)
		}
	}
	result := db.Find(&users)
	if result.Error != nil {
		err = result.Error
		return
	}
	return
}

func (this *User) GetUserById(Id int) (*User, error) {
	var user User
	result := Mysql.DB.First(&user, Id)
	if result.Error != nil {
		return nil, util.DataNotFoundError
	}
	return &user, nil
}

func (this *User) SetPassword(password string) (err error) {
	this.Password, err = util.GeneratePassword(password)
	if err != nil {
		return
	}
	return

}

func (this *User) CheckPassword(password string) (isOk bool, err error) {
	isOk, err = util.ValidatePassword(password, this.Password)
	return
}

func (this *User) DeleteById(id int) (err error) {
	user := User{}
	result := Mysql.DB.Where("id =?", id).Delete(&user)
	if result != nil {
		err = result.Error
		return
	}
	return
}

func (this *User) Role() (err error) {
	var role []Role
	related := Mysql.DB.Model(this).Related(&role)
	if related.Error != nil {
		err = related.Error
		return
	}
	fmt.Println(this)
	fmt.Println(role)
	return
}
