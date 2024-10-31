package main

type VegMealBuilder struct {
	mainCourse string
	sideDish   string
	bread      string
}

func (v *VegMealBuilder) WithMainCourse() {
	v.mainCourse = "Paneer tikka masala"
}

func (v *VegMealBuilder) WithBread() {
	v.bread = "Garlic Naan"
}

func (v *VegMealBuilder) WithSideDish() {
	v.sideDish = "Jeera rice"
}

func (v VegMealBuilder) GetMeal() Meal {
	return Meal{
		mainCourse: v.mainCourse,
		sideDish:   v.sideDish,
		bread:      v.bread,
	}
}
