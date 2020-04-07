/**
 * @Author: youxingxiang
 * @Description:
 * @File:  permissionModel
 * @Version: 1.0.0
 * @Date: 2020-04-03 15:47
 */
package models

type PerMission struct {
	Base
	Name string `json:"name"`
	Slug string `json:"slug"`
	HttpMethod string `json:"http_method"`
	HttpPath string `json:"http_path"`
}
func (PerMission) TableName() string {
	return "admin_permissions"
}
