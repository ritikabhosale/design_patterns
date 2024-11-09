package main

type Folder struct {
	Name     string
	Children []Inode
}

func (f Folder) clone() Inode {
	clonedChildren := make([]Inode, 0)

	for _, child := range f.Children {
		clonedChildren = append(clonedChildren, child.clone())
	}

	return Folder{Name: f.Name + "_clone", Children: clonedChildren}
}

func (f Folder) print(indentation string) {
	println(indentation + f.Name)

	for _, child := range f.Children {
		child.print(indentation + indentation)
	}
}
