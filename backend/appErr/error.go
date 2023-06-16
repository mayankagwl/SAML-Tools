package appErr

type ErrorResponse struct {
	Error            string `json:"error"`
	ErrorDescription string `json:"error_description"`
}

func NewError(err, desc string) ErrorResponse {
	return ErrorResponse{
		Error:            err,
		ErrorDescription: desc,
	}
}
