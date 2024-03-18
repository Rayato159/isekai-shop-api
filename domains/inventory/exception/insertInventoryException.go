package exception

import "fmt"

type InsertInventoryException struct {
	PlayerID string
	ItemID   uint64
}

func (e *InsertInventoryException) Error() string {
	return fmt.Sprintf("Failed to insert inventory for playerID: %s and itemID: %d", e.PlayerID, e.ItemID)
}
