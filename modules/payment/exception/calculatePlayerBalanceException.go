package exception

import "fmt"

type PlayerBalanceShowingException struct {
	PlayerID string
}

func (e *PlayerBalanceShowingException) Error() string {
	return fmt.Sprintf("Failed to calculate player balance for playerID: %s", e.PlayerID)
}
