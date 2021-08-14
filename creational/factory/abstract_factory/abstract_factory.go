package abstract_factory

// IChair 产品-椅子
type IChair interface {
	HasLegs() bool // 是否有腿
}

// ISofa 产品-沙发
type ISofa interface {
	MaterialMade() string // 材料制作 实木，塑料
}

// IFurnitureFactory 家具工厂 对工厂的抽象
type IFurnitureFactory interface {
	CreateChair() IChair // 创建椅子
	CreateSofa() ISofa   // 创建沙发
}

// ModernChar 现代风格椅子
type ModernChar struct{}

func (this *ModernChar) HasLegs() bool {
	return false
}

// ModernSofa 现代风格沙发
type ModernSofa struct{}

func (this *ModernSofa) MaterialMade() string {
	return "塑料"
}

// ModernFactory 现代工厂
type ModernFactory struct{}

func (this *ModernFactory) CreateChair() IChair {
	return &ModernChar{}
}

func (this *ModernFactory) CreateSofa() ISofa {
	return &ModernSofa{}
}
