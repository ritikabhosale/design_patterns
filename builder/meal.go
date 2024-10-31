package main

import "fmt"

type MealBuilder interface {
	WithMainCourse()
	WithSideDish()
	WithBread()
	GetMeal() Meal
}

type Meal struct {
	mainCourse string
	sideDish   string
	bread      string
}

func (m Meal) String() string {
	return fmt.Sprintf("Non-Veg Meal:\nMain Course: %s\nSide Dish: %s\nBread: %s\n",
		m.mainCourse, m.sideDish, m.bread)
}
