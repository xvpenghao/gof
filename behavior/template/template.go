package template

import "fmt"

/*
模拟做饭
参考链接 https://learnku.com/articles/33704
通用步骤在抽象类中实现，变化的步骤在具体的子类中实现
*/

type Cooker interface {
	Open()
	Fire()
	Cooke()
	OutFire()
	Close()
}

type CookMenu struct{}

func (this *CookMenu) Open() {
	fmt.Println("打开开关")
}

func (this *CookMenu) Fire() {
	fmt.Println("开火")
}

// Cooke 做菜，交给具体的子类实现
func (this *CookMenu) Cooke() {}

func (this *CookMenu) OutFire() {
	fmt.Println("关火")
}

func (this *CookMenu) Close() {
	fmt.Println("关闭开关")
}

// DoCook 这样写就避免了 子类重写了
func DoCook(cook Cooker) {
	cook.Open()
	cook.Fire()
	cook.Cooke()
	cook.OutFire()
	cook.Close()
}

type Tomato struct {
	CookMenu // 模拟继承
}

func (this *Tomato) Cooke() { // 重写了 父类的方法
	fmt.Println("做西红柿")
}
