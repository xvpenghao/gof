package simple_factory

import "fmt"

// 简单工厂，对具有相似的对象的创建进行统一管理，
// 避免多次创建的同一个对象的麻烦
// 对外隐藏 复杂对象的实例化步骤

type IOperation interface {
	PrintName()
}

// OperationAdd add
type OperationAdd struct {
	Name string
}

func (this *OperationAdd) PrintName() {
	fmt.Println(this.Name)
}

// OperationSub sub
type OperationSub struct {
	Name string
}

func (this *OperationSub) PrintName() {
	fmt.Println(this.Name)
}

// OperationFactory 定义操作符创建工厂
type OperationFactory struct{}

// CreateOperation 根据 操作符名称，创建实例对象
// 如果后期新增新的 操作符，则需要修改这个地方，这方法也是核心方法，
// 如果处理问题，则会影响到其他对象的创建
func (this *OperationFactory) CreateOperation(operationName string) IOperation {
	switch operationName {
	case "+":
		return &OperationAdd{Name: operationName}
	case "-":
		return &OperationSub{Name: operationName}
	default:

	}
	return nil
}
