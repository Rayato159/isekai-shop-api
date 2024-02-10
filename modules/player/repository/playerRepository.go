package repository

type PlayerRepository interface {
	InsertPlayer() (string, error)
}
