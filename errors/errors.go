package errors

// IError defines the interface that we expect our internal errors
// to implement
type IError interface {
	GetCode() string
	GetMessage() string
	GetField() string
}

// ProtoConstructable allows for IErrors implementations
// that be converted to Error proto messages.
type ProtoConstructable interface {
	ToProto() *Error
}

// AppError is a custom error, implementing IError and ProtoConstructable
type AppError struct {
	Code    string
	Message string
	Field   string
}

func (e AppError) GetCode() string {
	return e.Code
}

func (e AppError) GetMessage() string {
	return e.Message
}

func (e AppError) GetField() string {
	return e.Field
}

func (e AppError) Error() string {
	return e.GetCode() + ": " + e.GetMessage()
}

func (e AppError) ToProto() *Error {
	return &Error{
		Code:    e.GetCode(),
		Message: e.GetMessage(),
		Field:   e.GetField(),
	}
}
