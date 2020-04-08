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
	"rbac/app/models"
	"rbac/util"
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
		utilGin := util.Gin{c}
		if shouldPassThrough(c.Request.URL.Path) {
			c.Next()
		} else {
			// 获取当前用户的所有权限
			value, _ := c.Get("loginUser")
			// 类型断言
			loginUser := value.(*models.User)
			// 用户的所有权限
			permissions := loginUser.GetAllPermissions()
			//permissionShouldPassThrough(permissions,c.Request.URL.Path,c.Request.Method)
			havePermission := false
			for _,permission := range permissions {
				if permission.ShouldPassThrough(c.Request) == true {
					havePermission = true
					break
				}
			}
			if havePermission == true {
				c.Next()
			} else {
				utilGin.ParamsError(util.UnAuthError.Error())
				utilGin.Ctx.Abort()
				return
			}
		}

		//c.Next();

	}
}

func shouldPassThrough(path string) (ok bool) {
	for _, v := range Excepts {
		if path != "/" {
			path = strings.TrimRight(path, "/")
		}
		if strings.EqualFold(v, path) {
			ok = true
			return
		}
	}
	return
}
