package adapter

import "fmt"

// https://blog.csdn.net/carson_ho/article/details/54910430
// 小明买了一个进口电视，只支持110V，而家里都是220v，不兼容，需要调整

type TV struct{}

func (this *TV) Playll0v(voltage int) {
	fmt.Println("我播放了 ", voltage)
}

// 对象适配器

// Target 目标 客户端需要的接口 期待得到的 110v的电压
type Target interface {
	ConvertVoltage() int // 转换为电压
}

// PowerPort220V 适配者
type PowerPort220V struct{}

func (this *PowerPort220V) Output220v() {
	fmt.Println("输出220v")
}

// AdapterTV220v 适配器
type AdapterTV220v struct {
	adaptee *PowerPort220V
}

func (this *AdapterTV220v) ConvertVoltage() int {
	this.adaptee.Output220v()
	fmt.Println(" 220v电压 转化为 110v")
	return 110
}
