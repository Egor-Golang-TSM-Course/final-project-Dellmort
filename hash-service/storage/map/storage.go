package storagemap

import "errors"

var (
	ErrNotFound = errors.New("hash not found")
)

type StorageMap struct {
	hashes map[string]string
}

func NewStorageMap() *StorageMap {
	return &StorageMap{
		hashes: make(map[string]string),
	}
}

func (s *StorageMap) SaveHash(payload string, hash string) (string, error) {
	if _, ok := s.hashes[payload]; !ok {
		s.hashes[payload] = hash
	}

	return hash, nil
}

func (s *StorageMap) CheckHash(payload string) (bool, error) {
	// check if hash exists in map
	if _, ok := s.hashes[payload]; ok {
		return true, nil
	}

	return false, nil
}

func (s *StorageMap) GetHash(payload string) (string, error) {
	if hash, ok := s.hashes[payload]; ok {
		return hash, nil
	}

	return "", ErrNotFound
}
