package exception

import "fmt"

type PlayerItemsFinding struct {
	PlayerID string
}

func (e *PlayerItemsFinding) Error() string {
	return fmt.Sprintf("finding player items for playerID: %s", e.PlayerID)
}
