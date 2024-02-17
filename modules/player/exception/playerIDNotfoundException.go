package exception

type PlayerIDNotFoundException struct{}

func (e *PlayerIDNotFoundException) Error() string {
	return "Player ID not found"
}
