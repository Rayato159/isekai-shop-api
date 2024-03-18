package exception

type ItemCreatingException struct{}

func (e *ItemCreatingException) Error() string {
	return "Failed to create item"
}
