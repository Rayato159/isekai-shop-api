package exception

type ItemListingException struct{}

func (e *ItemListingException) Error() string {
	return "Find items failed"
}
