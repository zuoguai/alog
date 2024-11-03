package pkg

import (
	"fmt"
)

const IsAdmin = "is_admin"
const UserID = "user_id"
const ReqID = "req_id"
const (
	TokenExpireDuration = 60 * 60 * 24 * 3
	JWTDefaultSecret    = "wenchengwudeqianqiuwanzai"
	JWTDefaultIssuer    = "作怪"
)

// 时间标准化
const (
	SysTimeFormat      = "2006-01-02 15:04:05"
	SysTimeFormatShort = "2006-01-02"
)

const (
	SortDesc = "desc"
	SortAsc  = "asc"
)

const (
	ArticleStatusNormal  = 1
	ArticleStatusDraft   = 2
	ArticleStatusDeleted = 3
	DefaultPageSize      = 10
)

type ErrCode int // 错误码

const (
	Success         ErrCode = 0
	ErrLogin        ErrCode = 0400
	ErrBind         ErrCode = 0401
	ErrInput        ErrCode = 0402
	ErrJwtParse     ErrCode = 0403
	ErrDataNotFound ErrCode = 0404
	ErrDataExist    ErrCode = 0405
)

var errMsgDic = map[ErrCode]string{
	Success:         "ok",
	ErrLogin:        "登录错误 请检查用户名和密码",
	ErrBind:         "参数绑定错误",
	ErrInput:        "参数错误",
	ErrJwtParse:     "jwt解析错误",
	ErrDataNotFound: "数据不存在",
	ErrDataExist:    "数据已存在",
}

// GetErrMsg 获取错误描述
func GetErrMsg(code ErrCode) string {
	if msg, ok := errMsgDic[code]; ok {
		return msg
	}
	return fmt.Sprintf("unknown error code %d", code)
}
