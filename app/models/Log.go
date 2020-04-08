/**
 * @Author: youxingxiang
 * @Description:
 * @File:  Log
 * @Version: 1.0.0
 * @Date: 2020-04-07 14:20
 */
package models
type Log struct {
	Base
	UserId int `json:"user_id"`
	Path string `json:"path"`
	Method string `json:"method"`
	Ip string `json:"ip"`

}
func (Log) TableName() string {
	return "admin_operation_log"
}
