package exception

import "fmt"

type UpdatePlayerException struct {
	PlayerID string
}

func (e *UpdatePlayerException) Error() string {
	return fmt.Sprintf("Error updating player: %s", e.PlayerID)
}
