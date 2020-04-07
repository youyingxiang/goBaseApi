/**
 * @Author: youxingxiang
 * @Description:
 * @File:  userRoleModel
 * @Version: 1.0.0
 * @Date: 2020-04-03 16:55
 */
package models
type UserRole struct {
	Base
	RoleId int `json:"role_id"`
	UserId int `json:"user_id"`

}

func (UserRole) TableName() string {
	return "admin_role_users"
}
