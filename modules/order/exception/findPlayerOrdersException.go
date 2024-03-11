package exception

type PlayerOrderListingException struct{}

func (e *PlayerOrderListingException) Error() string {
	return "Error finding player orders"
}
