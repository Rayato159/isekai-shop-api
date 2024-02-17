package exception

import "fmt"

type PlayerNotFoundException struct {
	PlayerID string
}

func (e *PlayerNotFoundException) Error() string {
	return fmt.Sprintf("Error finding player: %s", e.PlayerID)
}
