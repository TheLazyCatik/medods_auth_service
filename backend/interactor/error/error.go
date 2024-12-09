package prjerror

type NotFoundError struct{}

func (e *NotFoundError) Error() string {
	return "record not found"
}

type NotAuthorizedError struct{}

func (e *NotAuthorizedError) Error() string {
	return "not authorized"
}

type InvalidJWTError struct{}

func (e *InvalidJWTError) Error() string {
	return "invalid token format"
}

type AlreadyExistsError struct{}

func (e *AlreadyExistsError) Error() string {
	return "already exists error"
}

type InvalidArgumentError struct{}

func (e *InvalidArgumentError) Error() string {
	return "Invalid argument error"
}
