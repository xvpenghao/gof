package decorator

import "testing"

func TestDecorator(t *testing.T) {
	// 抽象构件，具体构件
	var noddles Noddles = &Ramen{
		Name:  "拉面",
		Price: 12,
	}

	t.Log(noddles.Description(), noddles.PriceM())

	// 抽象装饰类 ,具体装饰类
	var noddlesDecorator = &EggNoddlesDecorator{
		Name:  "鸡蛋",
		Price: 2,
	}
	noddlesDecorator.SetNoddles(noddles)

	t.Log("=====装饰器=====")
	t.Log(noddlesDecorator.Description(), noddlesDecorator.PriceM())

}
