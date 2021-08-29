package memento

import "testing"

func TestMemento(t *testing.T) {
	// 负责人
	caretaker := &Caretaker{}

	// 原始对象
	originator := &Originator{State: "A"}

	// 添加到备忘录
	t.Log(originator.GetState())
	caretaker.AddMemento(originator.CreateMemento())

	// 添加到备忘录
	originator.SetState("B")
	t.Log(originator.GetState())
	caretaker.AddMemento(originator.CreateMemento())

	// 原始对象的恢复, 从负责人那里获取
	originator.RestoreMemento(caretaker.GetMemento(0))

	// 恢复对象 打印
	t.Log(originator.GetState())
}
