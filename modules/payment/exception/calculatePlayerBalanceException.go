package exception

import "fmt"

type CalculatePlayerBalanceException struct {
	PlayerID string
}

func (e *CalculatePlayerBalanceException) Error() string {
	return fmt.Sprintf("Failed to calculate player balance for playerID: %s", e.PlayerID)
}
