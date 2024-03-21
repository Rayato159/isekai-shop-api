package exception

import "fmt"

type FindAdmin struct {
	AdminID string
}

func (e *FindAdmin) Error() string {
	return fmt.Sprintf("Error finding admin: %s", e.AdminID)
}
