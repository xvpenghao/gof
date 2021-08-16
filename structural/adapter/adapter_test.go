package adapter

import (
	"testing"
)

// 对象适配器
func TestAdapterObj(t *testing.T) {
	// 小明买了 110v 的电视
	tv := TV{}

	// 将家用电220v 转化成 目标所需要的 110，
	p220v := &PowerPort220V{}
	ap := AdapterTV220v{adaptee: p220v} // 对外提供 100v，实际上还是 220v
	// 播放电视需要，提供110v
	tv.Playll0v(ap.ConvertVoltage())

	// 为什么适配器，用接口呢
	// 如果小明，此时又买了，电动牙刷 25v，其实也是将 220v的电压 转化为 25v
	// 其实 这是只需要 更好一个适配器，就可以了，无需改动代码。ap.ConvertVoltage() 不需要修改
}
