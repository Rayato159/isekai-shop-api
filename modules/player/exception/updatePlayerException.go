package exception

import "fmt"

type ProfileEditingException struct {
	PlayerID string
}

func (e *ProfileEditingException) Error() string {
	return fmt.Sprintf("Error updating player: %s", e.PlayerID)
}
