package recipe_model

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Recipe struct {
	gorm.Model
	Name        string
	Description string
	Ingredients []Ingredient `gorm:"many2many:has_quantity;"`
}

type Quantity struct {
	gorm.Model
	RecipeID     uint `gorm:"primaryKey"`
	IngredientID uint `gorm:"primaryKey"`
	Amount       uint
	Unit         string
}

type Ingredient struct {
	gorm.Model
	Name    string
	StoreID uint
	Store   Store
}

type Store struct {
	gorm.Model
	Name string
}

func Init_database(db_file string) (db *gorm.DB, err error) {
	db, err = gorm.Open(sqlite.Open(db_file), &gorm.Config{})
	return
}

func Setup_database(db *gorm.DB) (err error) {
	err = db.SetupJoinTable(&Recipe{}, "Ingredients", &Quantity{})
	if err != nil {
		return
	}
	err = db.AutoMigrate(&Recipe{}, &Ingredient{}, &Store{})
	if err != nil {
		return
	}
	return
}
