package exception

type InsertPaymentException struct{}

func (e *InsertPaymentException) Error() string {
	return "Failed to insert payment"
}
