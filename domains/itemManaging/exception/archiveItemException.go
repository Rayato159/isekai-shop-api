package exception

import "fmt"

type ItemArchivingException struct {
	ItemID uint64
}

func (e *ItemArchivingException) Error() string {
	return fmt.Sprintf("Failed to archive item with id %d", e.ItemID)
}
