package exception

import "fmt"

type DeleteInventoryException struct {
	ItemID uint64
}

func (e *DeleteInventoryException) Error() string {
	return fmt.Sprintf("Failed to delete inventory for itemID: %d", e.ItemID)
}
