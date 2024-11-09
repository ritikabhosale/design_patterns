package main

func main() {
	file1 := File{Name: "mumbai"}
	file2 := File{Name: "nagpur"}
	file3 := File{Name: "pune"}

	folder1 := Folder{
		Name: "jalgaon",
		Children: []Inode{
			File{Name: "parola"},
			File{Name: "amalner"},
		}}

	folder2 := Folder{
		Name:     "maharashtra",
		Children: []Inode{file1, file2, folder1, file3},
	}

	folder2.print(" ")

	println("--------------")

	clonedFolder := folder2.clone()
	clonedFolder.print(" ")
}
