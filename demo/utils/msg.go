package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type RES struct {
	Code  int         `json:"code"`
	Msg   string      `json:"msg"`
	Data  interface{} `json:"data"`
	Total int64       `json:"total,omitempty"`
}

const (
	SUCCESS = 200
	ERROR   = 500

	// 用户模块的错误
	ERROR_USERNAME_USED    = 1001
	ERROR_PASSWORD_WRONG   = 1002
	ERROR_USER_NOT_EXIST   = 1003
	ERROR_TOKEN_EXIST      = 1004
	ERROR_TOKEN_RUNTIME    = 1005
	ERROR_TOKEN_WRONG      = 1006
	ERROR_TOKEN_TYPE_WRONG = 1007
	ERROR_USER_NO_RIGHT    = 1008
	// 文章模块的错误
	ERROR_ART_NOT_EXIST = 2001
	// 分类模块的错误
	ERROR_CATENAME_USED  = 3001
	ERROR_CATE_NOT_EXIST = 3002
)

var codeMsg = map[int]string{
	SUCCESS:                "OK",
	ERROR:                  "FAIL",
	ERROR_USERNAME_USED:    "用户名已存在！",
	ERROR_PASSWORD_WRONG:   "密码错误",
	ERROR_USER_NOT_EXIST:   "用户不存在",
	ERROR_TOKEN_EXIST:      "TOKEN不存在,请重新登陆",
	ERROR_TOKEN_RUNTIME:    "TOKEN已过期,请重新登陆",
	ERROR_TOKEN_WRONG:      "TOKEN不正确,请重新登陆",
	ERROR_TOKEN_TYPE_WRONG: "TOKEN格式错误,请重新登陆",
	ERROR_USER_NO_RIGHT:    "该用户无权限",

	ERROR_ART_NOT_EXIST: "文章不存在",

	ERROR_CATENAME_USED:  "该分类已存在",
	ERROR_CATE_NOT_EXIST: "该分类不存在",
}

func GetMsg(code int) string {
	return codeMsg[code]
}

func ResOk(c *gin.Context, data interface{}, total int64) {
	Json := &RES{
		Code:  SUCCESS,
		Msg:   GetMsg(SUCCESS),
		Data:  data,
		Total: total,
	}
	c.JSON(http.StatusOK, Json)
}

func ResErr(c *gin.Context, code int) {
	Json := &RES{
		Code:  code,
		Msg:   GetMsg(code),
		Data:  "",
		Total: 0,
	}
	c.JSON(http.StatusBadRequest, Json)
}
