package simple_factory

import (
	"fmt"
	"testing"
)

type TOperationAdd struct {
	Name string
}

func (this *TOperationAdd) PrintName() {
	fmt.Println(this.Name)
}

type TOperationSub struct {
	Name string
}

func (this *TOperationSub) PrintName() {
	fmt.Println(this.Name)
}

func TestExample(t *testing.T) {
	// 创建 + 操作对象
	add := &TOperationAdd{Name: "+"}
	add.PrintName()

	// 创建 + 对象2
	add2 := &TOperationAdd{Name: "+"}
	add2.PrintName()

	// 创建 - 操作对象
	sub := &TOperationSub{Name: "-"}
	sub.PrintName()

	// 能够优化对象的创建，不像这样麻烦，例如，我只需要传递一个参数，然后直接返回一个对象
}

// 通过传入指定操作符名称，来简化创建对象
func TestOperationFactory_CreateOperation(t *testing.T) {
	operationFactory := &OperationFactory{}
	add := operationFactory.CreateOperation("+")
	sub := operationFactory.CreateOperation("-")

	add.PrintName()
	sub.PrintName()
}
