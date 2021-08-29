package observer

import (
	"testing"
)

func TestObserver(t *testing.T) {
	sb := &Subject{}

	ob1 := &ObServer{Id: 1}
	ob2 := &ObServer{Id: 2}
	sb.Register(ob1)
	sb.Register(ob2)

	sb.Notify("收到消息了吗")
	sb.Remove(ob1)

	sb.Notify("收到了")
}
