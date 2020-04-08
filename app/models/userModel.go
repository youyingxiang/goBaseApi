package models

import (
	"rbac/databases"
	"rbac/util"
)

type User struct {
	Base
	Username string `json:"username"`
	Name     string `json:"name"`
	Password string `json:"password"`
	//Id2      int    `json:"id2"`
	Roles       []*Role       `gorm:"many2many:admin_role_users"`
	Permissions []*Permission `gorm:"many2many:admin_user_permissions"`
	//Logs []Log `gorm:"foreignkey:Log.Ip;association_foreignkey:Id2"`
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

// 获取用户所有的权限
func (this *User) GetAllPermissions() (permissions []*Permission) {
	permissions = this.Permissions
	for _, role := range this.Roles {
		permissions = append(permissions, role.Permissions...)
	}
	return
}

// 获取当前用户的角色
func (this *User) GetRoles() (roles []*Role, err error) {
	//res := Mysql.DB.Model(this).Related(&roles, "Roles").RecordNotFound()
	//// 没有数据
	//if res {
	//	return nil, util.DataNotFoundError
	//}
	//return roles, nil
	return
}

func GetUserById(Id int) (*User, error) {
	var user User
	result := Mysql.DB.Preload("Roles.Permissions").Preload("Permissions").First(&user, Id)
	if result.Error != nil {
		return nil, util.DataNotFoundError
	}
	return &user, nil
}

func GetUsersByWhere(wheres map[interface{}]interface{}) (users []*User, err error) {
	db := Mysql.DB
	if wheres != nil {
		for i, v := range wheres {
			db = Mysql.DB.Where(i.(string)+"=?", v)
		}
	}
	result := db.Preload("Roles").Preload("Permissions").Find(&users)
	if result.Error != nil {
		err = result.Error
		return
	}
	return
}
