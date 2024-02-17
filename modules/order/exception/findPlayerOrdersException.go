package exception

type FindPlayerOrdersException struct{}

func (e *FindPlayerOrdersException) Error() string {
	return "Error finding player orders"
}
