package exception

type InsertItemException struct{}

func (e *InsertItemException) Error() string {
	return "Failed to insert item"
}
