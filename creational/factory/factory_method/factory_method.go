package factory_method

import "fmt"

// 参考链接 https://juejin.cn/post/6844903704189992967

// IDialogFactory 对工厂的抽象，让具体的实现者，来创建自己的新的产品
type IDialogFactory interface {
	CreateButton() IButton
}

// IButton 对产品的抽象
type IButton interface {
	Render()
}

type WindowDialogFactory struct{}

// CreateButton 具体的工厂来创建产品 这里没有参数，是因为工厂的产品只有一个
func (this *WindowDialogFactory) CreateButton() IButton {
	return &WindowButton{}
}

// WindowButton 产品
type WindowButton struct{}

func (this *WindowButton) Render() {
	fmt.Println("我 window 被渲染了")
}

type WebDialogFactory struct{}

// CreateButton 具体的工厂来创建产品
func (this *WebDialogFactory) CreateButton() IButton {
	return &WebButton{}
}

// WebButton 产品
type WebButton struct{}

func (this *WebButton) Render() {
	fmt.Println("我是 web 我被渲染了")
}
