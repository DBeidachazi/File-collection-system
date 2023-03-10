package res

// ErrorCode 预定义错误码
type ErrorCode int

const (
	SettingError ErrorCode = 1001 // 系统错误

)

var (
	ErrorMap = map[ErrorCode]string{
		SettingError: "系统错误",
	}
)
