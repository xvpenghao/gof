package compose

import "testing"

func TestCompose(t *testing.T) {
	file1 := &File{Name: "File1"}
	file2 := &File{Name: "File2"}
	file3 := &File{Name: "File3"}

	folder1 := &Folder{
		Name: "Folder1",
	}

	folder2 := &Folder{
		Name: "Folder2",
	}
	// 文件
	folder2.add(file2)
	folder2.add(file3)

	folder1.add(file1)
	// 文件夹to
	folder2.add(folder1)

	folder2.Search("rose")
}
