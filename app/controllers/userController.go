package userControllers

import (
	"github.com/gin-gonic/gin"
	"rbac/app/bindfield"
	"rbac/app/services"
	"rbac/util"
	"strconv"
)

// 保存用户
func Store(ctx *gin.Context) {
	var userBind bindfield.UserAdd
	var userService services.User
	utilGin := util.Gin{Ctx: ctx}

	err := utilGin.ShouldBind(&userBind)
	if err != nil {
		utilGin.ParamsError(err.Error())
		return
	}
	err = userService.Store(&userBind)
	if err != nil {
		utilGin.ServerError(err.Error())
		return
	}
	utilGin.Success(nil)

}

// 用户列表
func Index(ctx *gin.Context) {
	//params := ctx.Params
	utilGin := util.Gin{Ctx: ctx}
	userService := services.User{}
	users, err := userService.GetAll()
	if err != nil {
		utilGin.ParamsError(err.Error())
		return
	}
	utilGin.Success(users)
}

func Update(ctx *gin.Context) {
	var userBind bindfield.UserSave
	var userService services.User
	utilGin := util.Gin{Ctx: ctx}
	err := utilGin.ShouldBind(&userBind)
	if err != nil {
		utilGin.ParamsError(err.Error())
		return
	}
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		utilGin.ParamsError(err.Error())
		return
	}
	if id < 1 {
		utilGin.ParamsError(util.ParamsError.Error())
		return
	}
	err = userService.Update(id, &userBind)
	if err != nil {
		utilGin.ParamsError(err.Error())
		return
	}

	utilGin.Success(nil)

}

func Delete(ctx *gin.Context) {
	utilGin := util.Gin{Ctx: ctx}
	userService := services.User{}
	id, err := strconv.Atoi(utilGin.Ctx.Param("id"))
	if err != nil {
		utilGin.ParamsError(err.Error())
		return
	}
	if id < 1 {
		utilGin.ParamsError(util.ParamsError.Error())
		return
	}
	err = userService.Delete(id)
	if err != nil {
		utilGin.ParamsError(err.Error())
		return
	}
	utilGin.Success(nil)
}

// 用户登陆
func Login(ctx *gin.Context) {
	var bind bindfield.UserLogin
	var userService services.User
	utilGin := util.Gin{Ctx: ctx}
	err := utilGin.Ctx.ShouldBind(&bind)
	if err != nil {
		utilGin.ParamsError(err.Error())
		return
	}
	token, err := userService.Login(&bind)
	if err != nil {
		utilGin.UnAuthError(util.GenTokenFailError.Error())
		return
	}
	utilGin.Success(token)

}
