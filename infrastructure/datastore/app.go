package datastore

type AppDataStore struct {
	User UserDataStore
}

// AppDataStore constructor
func New() *AppDataStore {
	ads := AppDataStore{
		User: UserDataStore{},
	}
	return &ads
}