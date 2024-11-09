package main

type File struct {
	Name string
}

func (f File) clone() Inode {
	return File{Name: f.Name + "_clone"}
}

func (f File) print(indentation string) {
	println(indentation + f.Name)
}
