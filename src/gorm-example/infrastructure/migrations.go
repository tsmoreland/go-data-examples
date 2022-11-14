package infrastructure

import (
	"github.com/jinzhu/gorm"
	"github.com/tsmoreland/go-data-examples/src/gormexample/entities"
	"github.com/tsmoreland/go-data-examples/src/gormexample/models"
	"time"
)

func ResetCrudTables(db *gorm.DB) {
	db.DropTableIfExists(&models.EmployeeEntity{})
	db.CreateTable(&models.EmployeeEntity{})
}

func CreateTables(db *gorm.DB) {
	entities.CreateEmployeeTable(db)
	entities.CreateCalendarTable(db)
	entities.CreateAppointmentsTable(db)
}

func SeedDb(db *gorm.DB) error {
	tx := db.Begin()
	employees := []entities.Employee{
		{
			FirstName: "Bruce",
			LastName:  "Wayne",
			JobCategory: entities.JobCategoryLink{
				JobCategoryID:   models.JobCategorySuperHero,
				JobCategoryName: "Superhero",
			},
			Calendar: entities.Calendar{
				Name: "Batman TAS",
				Appointments: []entities.Appointment{
					{
						Title:       "On Leather Wings",
						Description: "A mysterious bat-like creature terrorizes Gotham City, causing the police force to pursue Batman. The Dark Knight must find the real perpetrator to clear his name.",
						StartTime:   time.Date(1992, time.May, 6, 9, 30, 0, 0, time.UTC),
						Length:      30,
					},
				},
			},
		},
		{
			FirstName: "Harley",
			LastName:  "Quinn",
			JobCategory: entities.JobCategoryLink{
				JobCategoryID:   models.JobCategoryVillain,
				JobCategoryName: "Villain",
			},
		},
	}

	for e := range employees {
		if err := tx.Create(&e).Error; err != nil {
			tx.Rollback()
			return err
		}
	}

	err := tx.Create(&entities.Employee{
		FirstName: "Clark",
		LastName:  "Kent",
		JobCategory: entities.JobCategoryLink{
			JobCategoryID:   models.JobCategorySuperHero,
			JobCategoryName: "Superhero",
		},
		Calendar: entities.Calendar{
			Name: "Superman TAS",
			Appointments: []entities.Appointment{
				{
					Title:       "World's Finest",
					Description: "The Joker steals a large piece of Kryptonite and then comes to Metropolis, offering to kill Superman for Lex Luthor in exchange for one billion dollars",
					StartTime:   time.Date(1997, time.October, 4, 9, 30, 0, 0, time.UTC),
					Length:      30,
					Attendees:   employees,
				},
			},
		},
	}).Error

	if err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()

	return nil
}
