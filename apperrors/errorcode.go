package apperrors

type MyAppError struct {
	ErrCode        // レスポンスとログに表示するエラーコード
	Message string // レスポンスに表示するエラーメッセージ
	Err     error  // エラーチェーンのための内部エラー
}

type ErrCode string

const (
	Unknown ErrCode = "U000"

	InsertDataFailed ErrCode = "S001"
	FetchDataFailed  ErrCode = "S002" // failed to exec select query
	EmptyData        ErrCode = "S003" // chosse article is not found
)

func (myErr *MyAppError) Error() string {
	return myErr.Err.Error()
}

func (myErr *MyAppError) Unwrap() error {
	return myErr.Err
}

func (code ErrCode) Wrap(err error, message string) error {
	return &MyAppError{
		ErrCode: code,
		Message: message,
		Err:     err,
	}
}
