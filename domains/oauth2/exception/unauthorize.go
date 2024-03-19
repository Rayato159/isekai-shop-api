package exception

type UnAuthorize struct{}

func (e *UnAuthorize) Error() string {
	return "unauthorize access"
}
