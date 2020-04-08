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
	Users []*User `gorm:"many2many:admin_role_users"`
	Permissions []*Permission `gorm:"many2many:admin_role_permissions"`
}

func (Role) TableName() string {
	return "admin_roles"
}
