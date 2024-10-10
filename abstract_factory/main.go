package main

func main() {
	factory, _ := GetBrandFactory("puma")

	shirt := factory.makeShirt()
	shirt.setColor()
	shirt.setSize()

	factory, _ = GetBrandFactory("nike")

	shirt = factory.makeShirt()
	shirt.setColor()
	shirt.setSize()
}
