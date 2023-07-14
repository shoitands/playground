package main

import (
	"testing"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver
//Company name set unique
func TestGORM(t *testing.T) {
	company := Company{Name: "jinzhu"}

	DB.Create(&company)

	for i := 0; i < 1000; i++ {
		if err := DB.Debug().Save(&company).Error; err != nil {
			t.Errorf("%v", err)
		}
	}
	//after some errors getting violation on unique value -
}
