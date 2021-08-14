package builder

import "testing"

type House0 struct {
	WindowType string
	DoorType   string
	Floor      int
}

func GetHouse0(windowType, doorType string, floor int) *House0 {
	return &House0{
		WindowType: windowType,
		DoorType:   doorType,
		Floor:      floor,
	}
}

func TestGetHouse0(t *testing.T) {
	GetHouse0("", "", 10)
	GetHouse0("铁窗户", "铁木", -1)
	GetHouse0("", "木门", 10)

	// 如果参数变量20呢？
	// 怎么解决，只需要传递，需要参数即可，或者不需要传递参数变量就可以创建 对应对象？
}

// 解决问题方法一
func TestBuild(t *testing.T) {
	// 冰屋子生成器
	director := NewDirector(&IglooBuilder{})
	h := director.BuildHouse()
	t.Logf("%+v", h)

	// 如果需要普通屋子，则需要创建 自己的生成器即可
}

//  解决问题方法二 利用golang的特性，创建对象
func TestBuildV2(t *testing.T) {
	h := GetHouseV2(
		WithWindowType("铁门"),
		WithDoorType("铁窗口"),
		WithFloor(100),
	)

	t.Logf("%+v", h)

}
