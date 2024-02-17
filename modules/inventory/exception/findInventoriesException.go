package exception

import "fmt"

type FindInventoriesException struct {
	PlayerID string
}

func (e *FindInventoriesException) Error() string {
	return fmt.Sprintf("Failed to find inventories for playerID: %s", e.PlayerID)
}
