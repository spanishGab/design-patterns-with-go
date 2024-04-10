package prototype

import "fmt"

type Folder struct {
	name string
	children []INode
}

func (folder *Folder) print(indentation string) {
	fmt.Println(indentation + folder.name)
	for _, child := range folder.children {
		child.print(indentation + indentation)
	}
}

func (folder *Folder) clone() INode {
	folderClone := &Folder{name: folder.name + "_clone"}
	var tempChildren []INode
	for _, child := range folder.children {
		childClone := child.clone()
		tempChildren = append(tempChildren, childClone)
	}
	folderClone.children = tempChildren
	return folderClone
}