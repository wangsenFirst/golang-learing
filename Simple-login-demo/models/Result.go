package models

type Result struct {
	Code    int32       `json:"code""`
	Message string      `json:"message""`
	Data    interface{} `json:"data""`
}

func Success(code int32, massege string, data interface{}) (result Result) {
	result = Result{
		code,
		massege,
		data,
	}
	return
}
