package internal

type DB interface {
	Create(query string, obj interface{}) (err error)
	Read()
	Update()
	Delete()
}
