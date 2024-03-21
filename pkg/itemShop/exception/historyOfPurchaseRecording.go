package exception

type HistoryOfPurchaseRecording struct{}

func (e *HistoryOfPurchaseRecording) Error() string {
	return "recording history of purchase failed"
}
