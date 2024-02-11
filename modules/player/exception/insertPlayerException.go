package exception

import "fmt"

type InsertPlayerException struct {
	PlayerID string
}

func (e *InsertPlayerException) Error() string {
	return fmt.Sprintf("Error inserting player: %s", e.PlayerID)
}
