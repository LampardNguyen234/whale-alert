package db

type KeyValueReaderWriter interface {
	KeyValueReader
	KeyValueWriter
}

type KeyValueReader interface {
	GetByKey(key []byte) ([]byte, error)
}

type KeyValueWriter interface {
	SetByKey(key []byte, value []byte) error
}
