package exception

import "fmt"

type InsertAdminException struct {
	AdminID string
}

func (e *InsertAdminException) Error() string {
	return fmt.Sprintf("Failed to insert admin with ID: %s", e.AdminID)
}
