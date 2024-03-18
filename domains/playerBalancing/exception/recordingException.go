package exception

type PlayerBalanceRecording struct{}

func (e *PlayerBalanceRecording) Error() string {
	return "Failed to insert balancing"
}
