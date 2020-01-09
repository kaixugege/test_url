package syserror

//定义错误类型

type Error interface {
	Code() int
	//go规定只要默认实现了 Error() string 方法的，就可以判定成error类型，
	// 因为Error的接口同事也实现了系统的error接口
	Error() string
	ReasonError() error
}

func New(msg string, reason error) Error {
	return UnKnowError{msg, reason}
}
