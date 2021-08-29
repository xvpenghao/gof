package strategy

// Strategy 定义策略的统一接口
type Strategy interface {
	Execute(a, b int) int
}

type ConcreteStrategyAdd struct{}

func (this *ConcreteStrategyAdd) Execute(a, b int) int {
	return a + b
}

type ConcreteStrategySubtract struct{}

func (this *ConcreteStrategySubtract) Execute(a, b int) int {
	return a - b
}

// Context 上下文定义了客户端关注的接口，在上下文类中添加一个成员变量用于保存对于策略对象的引用。
type Context struct {
	s Strategy
}

func (this *Context) SetStrategy(s Strategy) {
	this.s = s
}

func (this *Context) ExecuteStrategy(a, b int) int {
	return this.s.Execute(a, b)
}
