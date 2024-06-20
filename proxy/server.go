package proxy

type server interface {
	handleRequest(url string, method string) (int, string)
}
