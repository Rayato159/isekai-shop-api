package exception

type HistoryOfPurchasingRecordingException struct{}

func (e *HistoryOfPurchasingRecordingException) Error() string {
	return "Error inserting history of purchasing"
}
