package apperrors

type MyAppError struct {
	ErrCode        // error code about responses and logs
	Message string // error message for responses
	Err     error  `json:"-"` // internal error for internal process & judging error type
}

type ErrCode string

const (
	Unknown ErrCode = "U000"

	InsertDataFailed    ErrCode = "S001" // database insert error
	FetchDataFailed     ErrCode = "S002" // failed to exec select query
	EmptyData           ErrCode = "S003" // chosse article is not found
	NoTargetData        ErrCode = "S004" // Not found target data
	UpdateDataFailed    ErrCode = "S005" // failed to update a number of nices
	ReqBodyDecodeFailed ErrCode = "R001" // failed to decode json which is received from req body
	BadParam            ErrCode = "R002" // e.g. to error for strconv.Atoi(1)
)

// error wraper
func (myErr *MyAppError) Error() string {
	return myErr.Err.Error() // wrap Err member of myErr
}

func (myErr *MyAppError) Unwrap() error {
	return myErr.Err // unwrap myErr
}

func (code ErrCode) Wrap(err error, message string) error {
	return &MyAppError{
		ErrCode: code,
		Message: message,
		Err:     err,
	}
}
