package errors

type IError interface {
	GetCode() string
	GetMessage() string
	GetField() string
}

type IInternalError interface {
	GetCode() string
	GetRequestId() string
}

type ProtoConstructable interface {
	ToErrorProto() *Error
}

type AppError struct {
	Code     string
	Message  string
	Field    string
	Internal IInternalError
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

func (e AppError) ToErrorProto() *Error {
	return &Error{
		Code:    e.GetCode(),
		Message: e.GetMessage(),
		Field:   e.GetField(),
	}
}

type Internal struct {
	Code      string
	RequestId string
}

func (i Internal) GetCode() string {
	return i.Code
}

func (i Internal) GetRequestId() string {
	return i.RequestId
}
