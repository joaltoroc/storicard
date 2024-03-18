package transaction

type Storage interface {
	GetFile(fileName string) ([]byte, error)
}
