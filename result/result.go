package result

import (
	"ginblog-be/enum/code"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Result struct {
	Code  code.MyCode `json:"code"`
	Msg   string      `json:"message"`
	Data  interface{} `json:"data,omitempty"` // 可以是任何类型的数据
	Total interface{} `json:"total,omitempty"`
}

func ResOk(c *gin.Context, code code.MyCode, data interface{}, total interface{}) {
	Json := &Result{
		Code:  code,
		Msg:   code.GetCodeMsg(),
		Data:  data,
		Total: total,
	}
	c.JSON(http.StatusOK, Json)
}

func ResErr(c *gin.Context, code code.MyCode) {
	Json := &Result{
		Code:  code,
		Msg:   code.GetCodeMsg(),
		Data:  nil,
		Total: nil,
	}
	c.JSON(http.StatusBadRequest, Json)
}

func ResErrWithMsg(c *gin.Context, code code.MyCode, msg string) {
	Json := &Result{
		Code:  code,
		Msg:   msg,
		Data:  nil,
		Total: nil,
	}
	c.JSON(http.StatusBadRequest, Json)
}
