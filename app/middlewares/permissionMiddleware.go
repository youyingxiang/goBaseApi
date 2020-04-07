/**
 * @Author: youxingxiang
 * @Description:
 * @File:  permission
 * @Version: 1.0.0
 * @Date: 2020-04-03 14:54
 */
package middlewares

import (
	"github.com/gin-gonic/gin"
	"strings"
)

var Excepts []string

func init() {
	Excepts = append(Excepts,
		"/v1/admin/user",
	)
}

func PerMission() func(c *gin.Context) {
	return func(c *gin.Context) {
		if shouldPassThrough(c.Request.URL.Path) {
			c.Next()
		}
		// 获取当前用户的所有权限

		c.Next();

	}
}

func shouldPassThrough(path string) (ok bool) {
	for _, v := range Excepts {
		if path != "/" {
			path = strings.TrimRight(path,"/")
		}
		if strings.EqualFold(v,path) {
			ok = true
			return
		}
	}
	return
}
