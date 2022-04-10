package main

import "recipePlanner/recipe_model"

func main() {
	db, err := recipe_model.Init_database("recipeDB.db")
	if err != nil {
		panic("failed to connect database")
	}
	recipe_model.Setup_database(db)
}
