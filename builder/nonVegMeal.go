package main

type NonVegMealBuilder struct {
	mainCourse string
	sideDish   string
	bread      string
}

func (v *NonVegMealBuilder) WithMainCourse() {
	v.mainCourse = "Grilled Chicken Tikka"
}

func (v *NonVegMealBuilder) WithSideDish() {
	v.sideDish = "Chicken Mushroom"
}

func (v *NonVegMealBuilder) WithBread() {
	v.bread = "Garlic Naan"
}

func (v NonVegMealBuilder) GetMeal() Meal {
	return Meal{
		mainCourse: v.mainCourse,
		sideDish:   v.sideDish,
		bread:      v.bread,
	}
}
