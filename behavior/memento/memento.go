package memento

// Originator 原发器
type Originator struct {
	State string
}

func (this *Originator) CreateMemento() *Memento {
	return &Memento{State: this.State}
}

func (this *Originator) RestoreMemento(m *Memento) {
	this.State = m.GetSaveState()
}

func (this *Originator) SetState(s string) {
	this.State = s
}

func (this *Originator) GetState() string {
	return this.State
}

// Memento 备忘录
type Memento struct {
	State string // 用户存储 原发器的状态
}

// GetSaveState 得到保存的 原发器的对象
func (this *Memento) GetSaveState() string {
	return this.State
}

// Caretaker 负责人
type Caretaker struct {
	MementoList []*Memento
}

func (this *Caretaker) AddMemento(m *Memento) {
	this.MementoList = append(this.MementoList, m)
}

func (this *Caretaker) GetMemento(index int) *Memento {
	return this.MementoList[index]
}
