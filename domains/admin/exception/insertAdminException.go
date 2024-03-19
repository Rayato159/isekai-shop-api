package exception

import "fmt"

type InsertAdmin struct {
	AdminID string
}

func (e *InsertAdmin) Error() string {
	return fmt.Sprintf("Failed to insert admin with ID: %s", e.AdminID)
}
