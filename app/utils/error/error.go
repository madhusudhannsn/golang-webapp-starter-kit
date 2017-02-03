package error

import "github.com/madhusudhannsn/go-web-app/app/utils/logger"

// AppError : It is our custom error handler type
type AppError struct {
	msg  string
	code int
	err  error
}

func (e *AppError) Error() string {
	return e.msg
}

// GetInstance : This will create a new instance of AppError
func GetInstance(msg string, code int, err error) *AppError {
	logger.Error.Printf("Error msg: %s code: %d err: %v", msg, code, err)
	return &AppError{msg, code, err}
}
