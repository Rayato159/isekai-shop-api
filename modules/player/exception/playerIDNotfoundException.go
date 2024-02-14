package exception

type PlayerIDNotfoundException struct{}

func (e *PlayerIDNotfoundException) Error() string {
	return "Player ID not found"
}
