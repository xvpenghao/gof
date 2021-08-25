package command

import "fmt"

// Device 接受者
type Device interface {
	On()
	Off()
}

// TV 具体的接受者
type TV struct {
	isRunning bool
}

func (this *TV) On() {
	this.isRunning = true
	fmt.Println("Turning tv on")
}

func (this *TV) Off() {
	this.isRunning = false
	fmt.Println("Turning tv off")
}

type Cmd interface {
	Execute()
}

// OnCmd 将请求者 和 接受者 解耦 将 接受者的 方法 拆封成具体的命令
type OnCmd struct {
	Dev Device
}

func (this *OnCmd) Execute() () {
	this.Dev.On()
}

type OffCmd struct {
	Dev Device
}

func (this *OffCmd) Execute() () {
	this.Dev.Off()
}

// Button 请求者
type Button struct {
	Cmd Cmd
}

func (this *Button) Press() {
	// 具体执行什么命令，就看传递设置什么命令了
	this.Cmd.Execute()
}
