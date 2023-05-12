package errors

type errors struct {
	Code int
	Err  error
}

func (e *errors) Error() string {
	msg := ""
	if e.Err != nil {
		msg = e.Err.Error()
	}
	return msg
}
func IsError(err error) bool {
	v, ok := err.(*errors)
	return ok && v != nil
}
func FromError(err error) (int, error) {
	if err == nil {
		return 200, nil
	}
	v, ok := err.(*errors)
	if !ok || v == nil {
		return 500, err
	}
	return v.Code, err
}
func BadRequest(err error) *errors {
	return &errors{Code: 400, Err: err}
}
func UnAuthorized(err error) *errors {
	return &errors{Code: 401, Err: err}
}
func Forbidden(err error) *errors {
	return &errors{Code: 403, Err: err}
}
func NotFound(err error) *errors {
	return &errors{Code: 404, Err: err}
}
func NotAcceptable(err error) *errors {
	return &errors{Code: 406, Err: err}
}
func Conflict(err error) *errors {
	return &errors{Code: 409, Err: err}
}
func Internal(err error) *errors {
	return &errors{Code: 500, Err: err}
}
