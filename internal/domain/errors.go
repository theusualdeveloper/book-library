package domain

type ErrorCode string

const (
	ErrCodeNotFound ErrorCode = "NOT_FOUND"
	ErrCodeInternal ErrorCode = "INTERNAL"
)

type DomainError struct {
	Code    ErrorCode
	Message string
	Err     error
}

func (de DomainError) Error() string {
	return de.Message
}

func (de DomainError) Unwrap() error {
	return de.Err
}
