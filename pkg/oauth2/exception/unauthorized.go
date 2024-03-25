package exception

type Unauthorized struct{}

func (e *Unauthorized) Error() string {
	return "unauthorized"
}
