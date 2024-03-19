package exception

type AdminIDNotfound struct{}

func (e *AdminIDNotfound) Error() string {
	return "Admin ID not found"
}
