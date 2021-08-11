package factory_method

import "testing"

func TestExample(t *testing.T) {
	var factory IDialogFactory
	// 对工厂的抽象，让具体的实现者创建的自己的工厂，和产品，将责任划分下去
	factory = new(WindowDialogFactory)
	factory.CreateButton().Render()

	// web 工厂
	factory = new(WebDialogFactory)
	factory.CreateButton().Render()

}
