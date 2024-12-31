package formatter

type response struct {
	Message string      `json:"message"`
	Result  interface{} `json:"result"`
}

func SendResponse(msg string, res interface{}) response {
	resp := response{
		Message: msg,
		Result:  res,
	}

	return resp
}
