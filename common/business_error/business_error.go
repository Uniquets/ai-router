package business_error

type BusinessError struct {
	// msg
	msg string
}

// BusinessError 应实现 error 接口
func (e *BusinessError) Error() string {
	// 返回错误描述
	return e.msg
}
func New(msg string) error {
	return &BusinessError{msg: msg}
}

var (
	NoAuthError        = New("未登录")
	NoOperateAuthError = New("无操作权限")
	NoManageAuthError  = New("无管理权限")
	StatusError        = New("状态错误")
)
