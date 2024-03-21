package exception

type NoPermission struct{}

func (e *NoPermission) Error() string {
	return "no permission"
}
