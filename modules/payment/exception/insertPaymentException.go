package exception

type PaymentRecordingException struct{}

func (e *PaymentRecordingException) Error() string {
	return "Failed to insert payment"
}
