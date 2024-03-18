package exception

type ItemCountingException struct{}

func (e *ItemCountingException) Error() string {
	return "Count items failed"
}
