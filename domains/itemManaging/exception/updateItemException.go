package exception

import "fmt"

type ItemEditingException struct {
	ItemID uint64
}

func (e *ItemEditingException) Error() string {
	return fmt.Sprintf("Failed to update item with id %d", e.ItemID)
}
