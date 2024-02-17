package exception

type InsertOrderException struct{}

func (e *InsertOrderException) Error() string {
	return "Error inserting order"
}
