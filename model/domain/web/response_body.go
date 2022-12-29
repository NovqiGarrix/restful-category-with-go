package web

type ResponseBody struct {
	Code   int         `json:"code"`
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
}

type ErrorResponseBody struct {
	Code   int    `json:"code"`
	Status string `json:"status"`
	Error  string `json:"error"`
}
