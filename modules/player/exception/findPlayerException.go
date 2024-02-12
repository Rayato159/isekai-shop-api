package exception

import "fmt"

type FindPlayerException struct {
	PlayerID string
}

func (e *FindPlayerException) Error() string {
	return fmt.Sprintf("Error finding player: %s", e.PlayerID)
}
