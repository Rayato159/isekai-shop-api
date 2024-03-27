package exception

type PlayerCoinShowing struct{}

func (e *PlayerCoinShowing) Error() string {
	return "player coin showing failed"
}
