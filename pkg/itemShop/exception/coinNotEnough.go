package exception

type CoinNotEnough struct{}

func (e *CoinNotEnough) Error() string {
	return "coin is not enough"
}
