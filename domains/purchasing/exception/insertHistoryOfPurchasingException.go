package exception

type PurchasingHistoryRecording struct{}

func (e *PurchasingHistoryRecording) Error() string {
	return "Error inserting history of purchasing"
}
