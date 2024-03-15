package exception

type AdminIDNotfoundException struct{}

func (e *AdminIDNotfoundException) Error() string {
	return "Admin ID not found"
}
