package util

import "github.com/pkg/errors"

var (
	Success                 = errors.New("请求成功")
	UnAuthError             = errors.New("没有权限！")
	ParamsError             = errors.New("输入参数错误!")
	ServerError             = errors.New("服务器内部错误！")
	NotFoundError           = errors.New("未找到查询结果!")
	PasswordCheckError      = errors.New("密码比对错误")
	DataNotFoundError       = errors.New("数据未查询到!")
	GenTokenFailError       = errors.New("鉴权失败")
	ParseTokenError         = errors.New("token解析失败！")
	UserNameExistsError     = errors.New("用户账号已存在！")
	InputPasswordCheckError = errors.New("两次密码输入不一致")
)
