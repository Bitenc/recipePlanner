package recipe_model

import (
	"fmt"
	"os"
	"testing"

	"gorm.io/gorm"
)

var testDB *gorm.DB

func TestMain(m *testing.M) {
	testDB = setup("test.db")
	code := m.Run()
	shutdown("test.db")
	os.Exit(code)
}

// Creates testing environment in form of a file "testfile"
func setup(testfile string) *gorm.DB {
	db, err := Init_database(testfile)
	if err != nil {
		fmt.Print(err)
	}
	err = Setup_database(db)
	if err != nil {
		fmt.Print(err)
	}
	return db
}

func expect(t *testing.T, want string, is string) {
	if want != is {
		t.Errorf("ERROR: Wanted string \"%s\", is \"%s\".\n", want, is)
	}
}

// Clears and removes testing environment
func shutdown(testfile string) {
	os.Remove(testfile)
}

func TestCreateRecipeEntry(t *testing.T) {
	store := Store{Name: "Test Mall"}
	recipe := Recipe{
		Name:        "Testnoodles",
		Description: "Add noodles to hot water, rinse, add test sauce, eat. Yumm.",
		Ingredients: []Ingredient{
			{Name: "Test Noodles", Store: store},
			{Name: "Test Sauce", Store: store},
		},
	}
	result := testDB.Create(&recipe)

	var ret_recipe Recipe

	if result.Error != nil {
		t.Errorf("CREATE query returned %v, want nil.\n", result.Error)
	}

	testDB.Where("name = ?", "Testnoodles").First(&ret_recipe)

	expect(t, ret_recipe.Name, "Testnoodles")
	expect(t, ret_recipe.Description, "Add noodles to hot water, rinse, add test sauce, eat. Yumm.")
}
