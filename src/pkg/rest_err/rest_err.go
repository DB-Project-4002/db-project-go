package rest_err

type restErr struct {
	statusCode int
	err        string
}

func NewRestErr(statusCode int, err string) RestErr {
	return &restErr{statusCode: statusCode, err: err}
}

func (re *restErr) Error() string {
	return re.err
}

func (re *restErr) StatusCode() int {
	return re.statusCode
}
