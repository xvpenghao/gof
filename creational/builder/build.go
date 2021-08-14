package builder

// IBuilder 房子构建接口
type IBuilder interface {
	SetWindowType() // 设置窗户类型
	SetDoorType()   // 设置门的类型
	SetNumFloor()   // 设置地板数量
	GetHouse() House
}

// House 房子属性
type House struct {
	WindowType string
	DoorType   string
	Floor      int
}

// IglooBuilder 冰屋建造者 让具体的 构造者去设置自己需要的属性值
type IglooBuilder struct {
	House
}

func (this *IglooBuilder) SetWindowType() {
	this.WindowType = "木质窗户"
}

func (this *IglooBuilder) SetDoorType() {
	this.DoorType = "木质门"
}

func (this *IglooBuilder) SetNumFloor() {
	this.Floor = 20
}

// GetHouse 得到房子
func (this *IglooBuilder) GetHouse() House {
	this.House = House{
		WindowType: this.WindowType,
		DoorType:   this.DoorType,
		Floor:      this.Floor,
	}
	return this.House
}

// Director 管理者
type Director struct {
	Builder IBuilder
}

// NewDirector 创建一个管理者
func NewDirector(b IBuilder) *Director {
	return &Director{Builder: b}
}

func (this *Director) BuildHouse() House {
	this.Builder.SetWindowType()
	this.Builder.SetDoorType()
	this.Builder.SetNumFloor()
	return this.Builder.GetHouse()
}
