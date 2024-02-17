package exception

type NoPermissionException struct{}

func (e *NoPermissionException) Error() string {
	return "No permission"
}
