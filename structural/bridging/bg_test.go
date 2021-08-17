package bridging

import "testing"

func TestBridging(t *testing.T) {
	// 是一些实体的高阶控制层，改层自身不完成具体的工作，
	// 它需要将具体的工作委派给实现部分层
	var computer Computer = &Mac{}

	// 设置具体的打印机
	var printer Printer = &Hp{}
	computer.SetPrinter(printer)
	computer.Print()
}
