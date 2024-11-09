package main

type Inode interface {
	print(indentation string)
	clone() Inode
}
