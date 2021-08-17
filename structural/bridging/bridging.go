package bridging

import "fmt"

type Computer interface {
	Print()             // 打印，将具体的工作委派为 打印机
	SetPrinter(Printer) // 设置打印机
}

type Mac struct {
	Printer Printer
}

// Print 将工作委派给具体的 实现对象
func (this *Mac) Print() {
	this.Printer.PrintFile()
}

// SetPrinter 设置具体的打印机
func (this *Mac) SetPrinter(p Printer) {
	this.Printer = p
}

type Printer interface {
	PrintFile()
}

type Hp struct{}

func (this *Hp) PrintFile() {
	fmt.Println("打印，通过 惠普打印机")
}
