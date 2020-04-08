/**
 * @Author: youxingxiang
 * @Description:
 * @File:  permissionModel
 * @Version: 1.0.0
 * @Date: 2020-04-03 15:47
 */
package models

import (
	"fmt"
	"net/http"
	"strings"
)

type Permission struct {
	Base
	Name       string `json:"name"`
	Slug       string `json:"slug"`
	HttpMethod string `json:"http_method"`
	HttpPath   string `json:"http_path"`
}

func (Permission) TableName() string {
	return "admin_permissions"
}

func (this *Permission) ShouldPassThrough(request *http.Request) bool {
	// /v1/admin/user/53
	//path := request.URL.Path
	// PUT
	//method := request.Method
	// 权限允许的方法
	haveMethod := this.getHttpMethod()
	// 权限包含的路由
	havePath := strings.Split(this.getHttpPath(), "\n")
	for _, v := range havePath {
		if this.matchRequest(request, haveMethod, v) == true {
			fmt.Println(request.URL.Path, v, haveMethod)
			return true
		}
	}
	return false
}

func (this *Permission) matchRequest(request *http.Request, haveMethod []string, havePathStr string) bool {
	var path string
	if havePathStr == "/" {
		path = "/"
	} else {
		path = "/v1/admin" + strings.TrimRight(havePathStr, "/")
	}
	if path != strings.TrimRight(request.URL.Path, "/") {
		if strings.Contains(path, "*") {
			s := strings.TrimRight(request.URL.Path, "/")
			if strings.HasPrefix(s, strings.TrimRight(path, "/*")) == false {
				return false
			}
		} else {
			return false
		}
	}
	if len(haveMethod) == 0 {
		return true
	}
	for _, v := range haveMethod {
		if strings.ToUpper(v) == request.Method {
			return true
		}
	}
	return false
}

func (this *Permission) getHttpPath() string {
	return strings.Replace(this.HttpPath, "\r\n", "\n", -1)
}

func (this *Permission) getHttpMethod() []string {
	if this.HttpMethod == "" {
		return nil
	}

	return strings.Split(this.HttpMethod, ",")
}
