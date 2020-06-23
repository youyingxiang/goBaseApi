package util

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Gin struct {
	Ctx *gin.Context
}
type Response interface {
}
type SuccessResponse struct {
	Data interface{} `json:"data"`
}

type ErrResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (this *Gin) response(resp Response) {
	this.Ctx.JSON(http.StatusOK, resp)
	return
}

func (this *Gin) Success(data interface{}) {
	this.response(SuccessResponse{
		Data: data,
	})
}

// 301错误
func (this *Gin) UnAuthError(msg string) {
	if msg == "" {
		msg = UnAuthError.Error()
	}
	this.response(ErrResponse{
		http.StatusUnauthorized,
		msg,
	})
}

// 400错误
func (this *Gin) ParamsError(msg string) {
	if msg == "" {
		msg = UnAuthError.Error()
	}
	this.response(ErrResponse{
		http.StatusBadRequest,
		msg,
	})
}

// 500错误
func (this *Gin) ServerError(msg string) {
	if msg == "" {
		msg = ServerError.Error()
	}
	this.response(ErrResponse{
		http.StatusInternalServerError,
		msg,
	})
}

// 404错误
func (this *Gin) notFoundError(msg string) {
	if msg == "" {
		msg = NotFoundError.Error()
	}
	this.response(ErrResponse{
		http.StatusNotFound,
		msg,
	})
}

func (this *Gin) ShouldBind(bind bind) (err error) {
	err = this.Ctx.ShouldBind(bind)
	if err != nil {
		return
	}
	err = bind.Check()
	if err != nil {
		return
	}
	return
}
