package prototype

import "fmt"

type File struct {
	name string
}

func (file *File) print(indentation string) {
	fmt.Println(indentation + file.name)
}

func (file *File) clone() INode {
	return &File{
		name: file.name + "_clone",
	}
}
