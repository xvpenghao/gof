package abstract_factory

import "testing"

// 测试抽象工厂
func TestAf(t *testing.T) {
	// 客户需要 现代风格家具
	var factory IFurnitureFactory
	factory = &ModernFactory{} // 一个产品族的工厂

	// 创建的产品页都是 现代风格的
	factory.CreateChair()
	factory.CreateSofa()

	// 客户需要 装饰风艺术
	// 则只需要 创建对应的 风格工厂类，和 对应风格的产品
}
