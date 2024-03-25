package storage

type Storage interface {
	SaveHash(payload string, hash string) (string, error)
	CheckHash(payload string) (bool, error)
	GetHash(payload string) (string, error)
}
