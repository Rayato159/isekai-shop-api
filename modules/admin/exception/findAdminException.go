package exception

import "fmt"

type FindAdminException struct {
	AdminID string
}

func (e *FindAdminException) Error() string {
	return fmt.Sprintf("Error finding admin: %s", e.AdminID)
}
