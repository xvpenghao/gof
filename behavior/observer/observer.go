package observer

import "fmt"

type ISubject interface {
	Register(observer IObserver) // 订阅
	Remove(observer IObserver)
	Notify(msg string) // 发布
}

type Subject struct {
	ObserverList []IObserver // 订阅者列表
}

func (this *Subject) Register(observer IObserver) () {
	this.ObserverList = append(this.ObserverList, observer)
}

func (this *Subject) Remove(observer IObserver) {
	for index, o := range this.ObserverList {
		if o == observer {
			this.ObserverList = append(this.ObserverList[:index], this.ObserverList[index+1:]...)
		}
	}
}

func (this *Subject) Notify(msg string) () {
	for _, o := range this.ObserverList {
		o.Update(msg)
	}
}

// IObserver 观察者（订阅者s）
type IObserver interface {
	Update(msg string)
}

type ObServer struct {
	Id int
}

func (this *ObServer) Update(msg string) () {
	fmt.Println("ObServer", this.Id, msg)
}
