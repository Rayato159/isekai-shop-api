package exception

import "fmt"

type InventoryFilling struct {
	PlayerID string
	ItemID   uint64
}

func (e *InventoryFilling) Error() string {
	return fmt.Sprintf("inventory filling for playerID: %s and itemID: %d", e.PlayerID, e.ItemID)
}
