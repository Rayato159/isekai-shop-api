package exception

type CountItemsException struct{}

func (e *CountItemsException) Error() string {
	return "Count items failed"
}
