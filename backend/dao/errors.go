package dao

type ErrorTypeDao int

const (
	ErrTypeDatabaseConnection ErrorTypeDao = iota + 1
	ErrTypeDatabaseQuery
	ErrTypeDatabaseExec
	ErrTypeScanRows
	ErrTypeAffectRows
	ErrTypeNoSuchUser
	ErrTypeNoSuchProduct
	ErrTypeNoSuchComment
	ErrTypeSysOpenFile
	ErrTypeSysReadFile
	ErrTypeSysSaveFile
	ErrTypeEmailSend
	ErrTypeProductAlreadyExist
	ErrTypeUserAlreadyExist
	ErrTypeWrongCaptcha
	ErrTypeWrongPassword
	ErrTypeWrongRequestFormat
)

type ErrorDao struct {
	Type    ErrorTypeDao
	Message string
}

func (e *ErrorDao) Error() string {
	return e.Message
}

func NewErrorDao(errType ErrorTypeDao, message string) *ErrorDao {
	return &ErrorDao{
		Type:    errType,
		Message: message,
	}
}
