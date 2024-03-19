package exception

import "fmt"

type ItemNotEnough struct {
	ItemID uint64
}

func (e *ItemNotEnough) Error() string {
	return fmt.Sprintf("itemID: %d is not enough", e.ItemID)
}
