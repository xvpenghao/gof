package compose

import "fmt"

// 参考链接 https://refactoringguru.cn/design-patterns/composite/go/example

type Component interface {
	Search(keyword string) // 复杂对象和就 简单对象都有的
}

// File 文件
type File struct {
	Name string
}

func (this *File) Search(keyword string) {
	fmt.Printf("搜索关键字 %s 在文件 %s \n", keyword, this.Name)
}

func (this *File) GetName() string {
	return this.Name
}

// Folder 文件夹
type Folder struct {
	Components []Component
	Name       string
}

func (this *Folder) Search(keyword string) {
	fmt.Printf("递归搜索关键字 %s  在文件夹 %s\n", keyword, this.Name)
	for _, f := range this.Components {
		f.Search(keyword) // f 可能是文件，也可能是文件夹
	}
}

func (this *Folder) add(c Component) {
	this.Components = append(this.Components, c)
}

func (this *Folder) GetName() string {
	return this.Name
}
