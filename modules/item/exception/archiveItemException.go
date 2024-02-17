package exception

import "fmt"

type ArchiveItemException struct {
	ItemID uint64
}

func (e *ArchiveItemException) Error() string {
	return fmt.Sprintf("Failed to archive item with id %d", e.ItemID)
}
