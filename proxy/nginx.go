package proxy

type Nginx struct {
	application       *Application
	maxAllowedRequest int
	rateLimiter       map[string]int
}

func newNgninxServer() *Nginx {
	return &Nginx{
		application:       &Application{},
		maxAllowedRequest: 2,
		rateLimiter:       make(map[string]int),
	}
}

func (server *Nginx) handleRequest(url string, method string) (int, string) {
	allowed := server.checkRateLimit(url)
	if !allowed {
		return 403, "Not Allowed"
	}
	return server.application.handleRequest(url, method)
}

func (server *Nginx) checkRateLimit(url string) bool {
	if server.rateLimiter[url] == 0 {
		server.rateLimiter[url] = 1
	}
	if server.rateLimiter[url] > server.maxAllowedRequest {
		return false
	}
	server.rateLimiter[url]++
	return true
}
