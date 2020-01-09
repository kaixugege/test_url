package syserror

type Error404 struct {
	UnKnowError
}

func (u Error404) Code() int {
	return 1002
}

func (u Error404) Error() string {
	return "非法访问" // 这里已经知道了具体的错误类型了
}

func (u Error404) ReasonError() error {
	return u.reason
}
