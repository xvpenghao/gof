package builder

// HouseV2 房子属性
type HouseV2 struct {
	WindowType string
	DoorType   string
	Floor      int
}

type HoseV2OptsFun func(h *HouseV2)

func WithWindowType(windowType string) HoseV2OptsFun {
	return func(h *HouseV2) {
		h.WindowType = windowType
	}
}

func WithDoorType(doorType string) HoseV2OptsFun {
	return func(h *HouseV2) {
		h.DoorType = doorType
	}
}

func WithFloor(floor int) HoseV2OptsFun {
	return func(h *HouseV2) {
		h.Floor = floor
	}
}

// GetHouseV2 创建一个对象
func GetHouseV2(opts ...HoseV2OptsFun) *HouseV2 {
	h := &HouseV2{
		WindowType: "*",
		DoorType:   "*",
		Floor:      -1,
	}
	for _, o := range opts {
		o(h)
	}
	return h
}
