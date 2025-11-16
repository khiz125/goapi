package apperrors

type AppError struct {
	ErrCode
	Message string
	Err     error `json:"-"`
}

func (appErr *AppError) Error() string {
	return appErr.Err.Error()
}

func (code ErrCode) Wrap(err error, message string) error {
	return &AppError{ErrCode: code, Message: message, Err: err}
}

func (appErr *AppError) Unwrap() error {
	return appErr.Err
}
