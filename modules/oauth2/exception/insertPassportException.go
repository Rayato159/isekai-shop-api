package exception

import "fmt"

type InsertPassportException struct {
	PlayerID string
}

func (e *InsertPassportException) Error() string {
	return fmt.Sprintf("Error inserting passport: %s", e.PlayerID)
}
