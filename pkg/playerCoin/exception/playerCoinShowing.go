package exception

import "fmt"

type PlayerCoinShowing struct{}

func (e *PlayerCoinShowing) Error() string {
	return fmt.Sprintf("player coin showing failed")
}
