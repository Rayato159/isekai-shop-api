package exception

type OrderRecordingException struct{}

func (e *OrderRecordingException) Error() string {
	return "Error inserting order"
}
