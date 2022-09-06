package ecode

var msg = map[int]string{
	Success:         "ok",
	InvalidParams:   "invalid params",
	ErrorCreateUser: "failed to create user",
}

func GetMsg(code int) string {
	return msg[code]
}
