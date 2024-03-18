package exception

import "fmt"

type NotEnoughItemException struct {
	ItemID uint64
}

func (e *NotEnoughItemException) Error() string {
	return fmt.Sprintf("Not enough item with id: %d", e.ItemID)
}
