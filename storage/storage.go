package storage

type Storage interface {
	Put(string, string, []byte) error
	Get(string, string) ([]byte, error)
	List() ([]string, error)
	LatestVersion(string) (string, error)
}
