package types

type ErrorDefault struct {
	Status int    `json:"status"`
	Err    string `json:"err"`
}

type SuccessDefault struct {
	Status int    `json:"status"`
	Msg    string `json:"msg"`
}
