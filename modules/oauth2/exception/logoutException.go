package exception

type LogoutException struct{}

func (e *LogoutException) Error() string {
	return "Failed to logout user"
}
