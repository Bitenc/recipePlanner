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

// Clears and removes testing environment
func shutdown(testfile string) {
	os.Remove(testfile)
}

func TestCreateRecipeEntry(t *testing.T) {
	recipe := Recipe{Name: "Testnoodles", Description: "Add noodles to hot water, rinse, eat. Yumm."}
	result := testDB.Create(&recipe)

	if result.Error != nil {
		t.Errorf("Create query returned %v, want nil", result.Error)
	}

}
