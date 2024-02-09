package api

type Payload struct {
	X int `json:"x"`
	Y int `json:"y"`
}

type ResponseBody struct {
	Input  Payload `json:"input"`
	Output int     `json:"output"`
}

type ResponseBodyIsOdd struct {
	Output bool `json:"output"`
}
