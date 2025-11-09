package apperrors

type AppError struct {
	ErrCode
	Message string
  Err error
}

func (appErr *AppError) Error() string {
  return appErr.Err.Error()
}

func (appErr *AppError) Unwrap() error {
  return appErr.Err
}
