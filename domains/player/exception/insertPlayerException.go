package exception

import "fmt"

type PlayerCreatingException struct {
	PlayerID string
}

func (e *PlayerCreatingException) Error() string {
	return fmt.Sprintf("Error inserting player: %s", e.PlayerID)
}
