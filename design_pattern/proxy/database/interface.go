package proxy_database

type IDatabase interface {
	Set(input Input)
	Get(id string) string
}

type Input struct {
	Id      string
	Content string
}
