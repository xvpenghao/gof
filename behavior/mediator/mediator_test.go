package mediator

import "testing"

func TestMediator(t *testing.T) {
	u1 := &User{Name: "小张"}
	u1.sendMsg("你好小明")

	u2 := &User{Name: "小明"}
	u2.sendMsg("你好 小张")
}
