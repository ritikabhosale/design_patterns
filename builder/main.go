package main

func getBuilder(s string) MealBuilder {
	switch s {
	case "veg":
		return &VegMealBuilder{}
	case "nonveg":
		return &NonVegMealBuilder{}
	}

	return nil
}

func main() {
	nonVegMealBuilder := getBuilder("nonveg")
	director := NewDirector(nonVegMealBuilder)

	meal := director.MakeMeal()
	print(meal.String())
}
