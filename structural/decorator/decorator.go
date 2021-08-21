package decorator

import "fmt"

// Noddles 给一个类或者对象增加行为
// 参考链接 https://blog.csdn.net/qibin0506/article/details/51082510
// Noddles ，面条
type Noddles interface {
	Description() string
	PriceM() float64
}

// Ramen 拉面
type Ramen struct {
	Name  string
	Price float64
}

func (this *Ramen) Description() string {
	return this.Name
}

func (this *Ramen) PriceM() float64 {
	return this.Price
}

// NoddlesDecorator 面条的装饰器
type NoddlesDecorator interface {
	SetNoddles(noddles Noddles) // 设置面条
}

// EggNoddlesDecorator 具体的装饰器-鸡蛋面条
type EggNoddlesDecorator struct {
	noddles Noddles
	Name    string
	Price   float64
}

func (this *EggNoddlesDecorator) SetNoddles(noddles Noddles) {
	this.noddles = noddles
}

func (this *EggNoddlesDecorator) Description() string {
	return fmt.Sprintf("装饰[%s]-源物品[%s]", this.Name, this.noddles.Description())
}

func (this *EggNoddlesDecorator) PriceM() float64 {
	return this.Price + this.noddles.PriceM()
}
