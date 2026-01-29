package common

var ErrInvalidId = Error{ErrCode: "A0001", ErrMsg: "无效的ID"}
var ErrRecordNotFound = Error{ErrCode: "A0002", ErrMsg: "未查到记录"}
var ErrUnknown = Error{ErrCode: "XXXXX", ErrMsg: "未知错误"}

type Error struct {
	ErrCode string
	ErrMsg  string
}

func (e Error) Error() string {
	return e.ErrMsg
}
