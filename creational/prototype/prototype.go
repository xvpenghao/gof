package prototype

import "fmt"

type inode interface {
	print(string)
	clone() inode
}

// 文件
type file struct {
	name string
}

func (this *file) print(indentation string) { // 文件打印方式
	fmt.Println(this.name + indentation)
}

func (this *file) clone() inode {
	return &file{name: this.name + "_clone"}
}

// 文件夹
type folder struct {
	children []inode
	name     string
}

func (f *folder) print(indentation string) { // 文件夹的打印方式
	fmt.Println(indentation + f.name)
	for _, i := range f.children { // 文件
		i.print(indentation + indentation)
	}
}

// 也能实现递归的拷贝
func (f *folder) clone() inode {
	cloneFolder := &folder{
		children: nil,
		name:     f.name + "_clone",
	}
	var cloneChildren []inode
	for _, i := range f.children {
		cloneChildren = append(cloneChildren, i.clone())
	}

	cloneFolder.children = cloneChildren
	return cloneFolder
}
