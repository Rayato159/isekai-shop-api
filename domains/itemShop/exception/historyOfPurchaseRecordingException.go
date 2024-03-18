package exception

type HistoryOfPurchaseRecordingException struct{}

func (e *HistoryOfPurchaseRecordingException) Error() string {
	return "Error inserting history of purchasing"
}
