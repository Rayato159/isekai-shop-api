package exception

type PlayerHistoryOfPurchasingListingException struct{}

func (e *PlayerHistoryOfPurchasingListingException) Error() string {
	return "Error finding history of purchasing"
}
