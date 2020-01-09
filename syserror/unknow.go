package syserror

type UnKnowError struct {
	msg    string
	reason error
}

func (u UnKnowError) Code() int {
	return 1000
}

func (u UnKnowError) Error() string {
	if len(u.msg) == 0{
		return "未知错误"
	}
	return u.msg
}

func (u UnKnowError) ReasonError() error {
	return u.reason
}
