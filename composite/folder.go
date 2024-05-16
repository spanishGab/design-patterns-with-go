package composite

import "fmt"

type Folder struct {
	components []IComponent
	name 			 string
}

func (folder *Folder) search(keyword string) {
	fmt.Printf("Searching recursively for keyword %s in folder %s\n", keyword, folder.name)
	for _, component := range folder.components {
		component.search(keyword)
	}
}

func (folder *Folder) add(component IComponent) {
	folder.components = append(folder.components, component)
}
