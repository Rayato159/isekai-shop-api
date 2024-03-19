package exception

import "fmt"

type ItemNotFound struct {
	ItemID uint64
}

func (e *ItemNotFound) Error() string {
	return fmt.Sprintf("itemID: %d was not found", e.ItemID)
}
