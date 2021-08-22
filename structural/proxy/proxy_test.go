package proxy

import "testing"

func TestProxy(t *testing.T) {
	// app := Application{}
	// app.HandlerRequest("/app/status","GET")
	// app.HandlerRequest("/app/status","GET")

	// 可以在不修改-服务端代码情况下，通过代理，来实现 路由的访问控制
	pxy := NewNginxProxy()
	t.Log(pxy.HandlerRequest("/app/status", "GET"))
	t.Log(pxy.HandlerRequest("/app/status", "GET"))
	t.Log(pxy.HandlerRequest("/app/status", "GET"))
}
