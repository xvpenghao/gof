package state

import "testing"

func TestState(t *testing.T) {
	c := Context{State: &SwitchOnState{}}
	c.Request() // 不断的做 状态切换
	c.Request()
	c.Request()
}
