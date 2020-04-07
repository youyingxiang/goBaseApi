/**
 * @Author: youxingxiang
 * @Description:
 * @File:  menuModel
 * @Version: 1.0.0
 * @Date: 2020-04-03 15:52
 */
package models

type Menu struct {
	Base
	ParentId int `json:"parent_id"`
	Order int `json:"order"`
	Title string `json:"title"`
	Icon string `json:"icon"`
	Uri string `json:"uri"`
	Permission string `json:"permission"`
}
func (Menu) TableName() string {
	return "admin_menu"
}
