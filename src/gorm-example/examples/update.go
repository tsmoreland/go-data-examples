package examples

import (
	"github.com/jinzhu/gorm"
	"github.com/tsmoreland/go-data-examples/src/gormexample/entities"
	"github.com/tsmoreland/go-data-examples/src/gormexample/models"
)

func UpdateHarleyQuinn(db *gorm.DB, harley *entities.Employee, harleyCalendar *entities.Calendar) {
	// normally these calls would take &harley but she's already a pointer

	db.Where(harley).First(harley).Related(harleyCalendar)
	db.Debug().Model(harley).Update("job_category_id", models.JobCategoryAntiHero)
	db.Debug().Model(harley).Update("job_category_name", "Anti Hero")
	db.Debug().Model(harley).Updates(
		map[string]interface{}{
			"job_category_id":   models.JobCategoryAntiHero,
			"job_category_name": "Anti Hero",
		})
}

func BulkUpdateWhenSetValueIsKnown(db *gorm.DB) {
	db.
		Debug().
		Table("appointments").
		Where("length = ?", 30).
		Update("Length", 22) // 24 is more accurate, handled by subsequent update
}

func BulkUpdateWhenSetValueIsCalculated(db *gorm.DB) {
	db.
		Debug().
		Table("appointments").
		Where("length = ?", 22).
		Update("length", gorm.Expr("length + 2"))
}
