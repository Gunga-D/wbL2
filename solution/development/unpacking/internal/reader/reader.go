package reader

type Reader interface {
	ReadAll(key string) (string, error)
}
