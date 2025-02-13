package data

type Database interface {
	Connect(uri string) (err error)
}
