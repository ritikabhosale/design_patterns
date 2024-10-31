package main

type Director struct {
	builder MealBuilder
}

func (d *Director) MakeMeal() Meal {
	d.builder.WithMainCourse()
	d.builder.WithSideDish()
	d.builder.WithBread()

	return d.builder.GetMeal()
}

func NewDirector(mealBuilder MealBuilder) *Director {
	return &Director{builder: mealBuilder}
}
