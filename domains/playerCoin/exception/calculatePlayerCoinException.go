package exception

import "fmt"

type PlayerCoinShowingException struct {
	PlayerID string
}

func (e *PlayerCoinShowingException) Error() string {
	return fmt.Sprintf("Failed to calculate player coin for playerID: %s", e.PlayerID)
}
