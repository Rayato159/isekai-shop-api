package exception

type RenewTokenException struct{}

func (e *RenewTokenException) Error() string {
	return "Failed to renew token"
}
