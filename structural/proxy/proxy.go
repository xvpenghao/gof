package proxy

// Server 定了 服务端和 代理对象共有的接口
type Server interface {
	HandlerRequest(string, string) (int, string)
}

type NginxProxy struct {
	App               *Application // 代理真实对象 不写接口，是因为，可能每个的限制都不一样
	MaxAllowedRequest int
	RateLimiter       map[string]int
}

func NewNginxProxy() *NginxProxy {
	return &NginxProxy{
		App:               &Application{},
		MaxAllowedRequest: 2,
		RateLimiter:       map[string]int{},
	}
}

func (this *NginxProxy) HandlerRequest(url, method string) (int, string) {
	// 做一些访问控制的限制
	if this.checkRateLimiting(url) {
		return 403, "Not Allowed"
	}

	return this.App.HandlerRequest(url, method)
}

func (this *NginxProxy) checkRateLimiting(url string) bool {
	if this.RateLimiter[url] == 0 {
		this.RateLimiter[url] = 1
	}

	if this.RateLimiter[url] > this.MaxAllowedRequest {
		return true
	}

	this.RateLimiter[url] = this.RateLimiter[url] + 1
	return false
}

type Application struct{}

func (this *Application) HandlerRequest(url, method string) (int, string) {
	if url == "/app/status" && method == "GET" {
		return 200, "Ok"
	}

	if url == "/create/user" && method == "POST" {
		return 201, "User Created"
	}
	return 404, "Not Ok"
}
