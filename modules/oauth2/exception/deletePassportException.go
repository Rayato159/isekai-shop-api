package exception

type DeletePassportException struct{}

func (e *DeletePassportException) Error() string {
	return "Error deleting passport"
}
