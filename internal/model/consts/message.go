package consts

import "errors"

var (
	NotFoundErr    = errors.New("object not found")
	InvalidJsonErr = errors.New("invalid data")
	InvalidURLErr  = errors.New("invalid URL param")
	ServerErr      = errors.New("server error")
	InvalidRequest = errors.New("bad request")
)
