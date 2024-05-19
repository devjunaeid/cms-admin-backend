package types

type ErrorDefault struct {
	Status int    `json:"status"`
	Err    string `json:"err"`
}

type SuccessDefault struct {
	Status int    `json:"status"`
	Msg    string `json:"msg"`
}

type JwtTokenRes struct {
	Token  string `json:"token"`
	Status int    `json:"status"`
}
