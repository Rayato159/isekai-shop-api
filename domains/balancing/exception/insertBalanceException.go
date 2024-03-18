package exception

type BalancingRecordingException struct{}

func (e *BalancingRecordingException) Error() string {
	return "Failed to insert balancing"
}
