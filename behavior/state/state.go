package state

import "fmt"

// 参考链接 https://segmentfault.com/a/1190000021094565
// 实现类似 电灯开关的 转换

type IState interface { // 状态接口
	handler(ctx *Context)
}

type Context struct { // 环境类 持有 state
	State IState
}

func (this *Context) Request() {
	this.State.handler(this)
}

func (this *Context) SetState(s IState) {
	this.State = s
}

type SwitchOnState struct{}

func (this *SwitchOnState) handler(context *Context) () {
	fmt.Println("SwitchOnState on")
	context.SetState(&SwitchOffState{})
}

type SwitchOffState struct{}

func (this *SwitchOffState) handler(context *Context) () {
	fmt.Println("SwitchOnState off")
	context.SetState(&SwitchOnState{})
}
