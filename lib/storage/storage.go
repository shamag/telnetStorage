package storage

type IStorage interface {
	Set(string, string)
	Get(string) (string, bool)
	Exist(string) bool
	Delete(key string) bool
}

type Storage map[string]string
