package state

type State interface {
	GenerateRandomState() (string, error)
	ParseState(state string) error
}
