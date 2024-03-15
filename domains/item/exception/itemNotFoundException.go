package exception

import "fmt"

type ItemNotFoundException struct {
	ItemID uint64
}

func (e *ItemNotFoundException) Error() string {
	return fmt.Sprintf("Item with ID %d not found", e.ItemID)
}
