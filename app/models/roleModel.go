/**
 * @Author: youxingxiang
 * @Description:
 * @File:  roleModel
 * @Version: 1.0.0
 * @Date: 2020-04-03 15:48
 */
package models

type Role struct {
	Base
	Name string  `json:"name"`
	Slug string  `json:"slug"`
	User []*User `gorm:"many2many:admin_role_users"`
}

func (Role) TableName() string {
	return "admin_roles"
}
