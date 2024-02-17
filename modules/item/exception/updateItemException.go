package exception

import "fmt"

type UpdateItemException struct {
	ItemID uint64
}

func (e *UpdateItemException) Error() string {
	return fmt.Sprintf("Failed to update item with id %d", e.ItemID)
}
