package exception

import "fmt"

type ItemEditing struct {
	ItemID uint64
}

func (e *ItemEditing) Error() string {
	return fmt.Sprintf("editing item id: %d failed", e.ItemID)
}
