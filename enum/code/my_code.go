package code

type MyCode int64

const (
	CodeSuccess           MyCode = 200
	CodeBadRequest        MyCode = 400
	CodeServerFail        MyCode = 500
	CodeInvalidParams     MyCode = 1001
	CodeUserExist         MyCode = 1002
	CodeUserNotExist      MyCode = 1003
	CodeInvalidPassword   MyCode = 1004
	CodeServerBusy        MyCode = 1005
	CodeInvalidToken      MyCode = 1006
	CodeInvalidAuthFormat MyCode = 1007
	CodeNotLogin          MyCode = 1008
	ErrVoteRepeated       MyCode = 1009
	ErrorVoteTimeExpire   MyCode = 1010
)

var codeMsgMap = map[MyCode]string{
	CodeSuccess:           "操作成功",
	CodeBadRequest:        "请求错误",
	CodeServerFail:        "服务器错误",
	CodeInvalidParams:     "请求参数错误",
	CodeUserExist:         "用户已存在",
	CodeUserNotExist:      "用户不存在",
	CodeInvalidPassword:   "密码错误",
	CodeServerBusy:        "服务器正忙，请稍后再试",
	CodeInvalidToken:      "无效的Token",
	CodeInvalidAuthFormat: "无效的认证格式",
	CodeNotLogin:          "用户未登录",
	ErrVoteRepeated:       "投票重复",
	ErrorVoteTimeExpire:   "投票时间已过",
}

func (code MyCode) GetCodeMsg() string {
	msg, ok := codeMsgMap[code]
	if ok {
		return msg
	}
	return codeMsgMap[CodeServerFail]
}
