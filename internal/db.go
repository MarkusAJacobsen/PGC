package internal

type DB interface {
	Connect() (err error)
	Create(query string, obj interface{}) (err error)
	Read(query string, obj interface{}, key string)
	Update()
	Delete()
	Do(query string, obj interface{}) (err error)
}
