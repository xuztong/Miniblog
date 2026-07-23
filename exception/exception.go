package exception

func NewApiException(code int, message string) *ApiException {
	return &ApiException{
		Code:    code,
		Message: message,
	}
}

type ApiException struct {
	Code     int    `json:"code"`
	Message  string `json:"message"`
	HttpCode int    `json:"-"`
}

func (a *ApiException) Error() string {
	return a.Message
}

func (a *ApiException) WithMessage(msg string) *ApiException {
	a.Message = msg
	return a
}

func (a *ApiException) WithHttpCode(code int) *ApiException {
	a.HttpCode = code
	return a
}
