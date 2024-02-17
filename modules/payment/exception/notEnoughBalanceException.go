package exception

type NotEnoughBalanceException struct{}

func (e *NotEnoughBalanceException) Error() string {
	return "Not enough balance"
}
