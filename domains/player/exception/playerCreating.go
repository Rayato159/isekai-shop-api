package exception

import "fmt"

type PlayerCreating struct {
	PlayerID string
}

func (e *PlayerCreating) Error() string {
	return fmt.Sprintf("playerID: %s failed to create", e.PlayerID)
}
