package errors

type appError struct {
	message string
	status  int
}

type IAppError interface {
	GetMessage() string
	GetStatus() int
}

func NewAppError(message string, status int) IAppError {
	return &appError{
		message: message,
		status:  status,
	}
}

func (e *appError) GetMessage() string {
	return e.message
}

func (e *appError) GetStatus() int {
	return e.status
}
