package exception

type UnAuthorizeException struct{}

func (e *UnAuthorizeException) Error() string {
	return "Unauthorize"
}
