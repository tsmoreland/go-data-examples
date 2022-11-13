package entities

import (
	"github.com/jinzhu/gorm"
	"github.com/tsmoreland/go-data-examples/src/gormexample/models"
	"time"
)

func CreateTables(db *gorm.DB) {
	CreateEmployeeTable(db)
	CreateCalendarTable(db)
	CreateAppointmentsTable(db)
}

func SeedDb(db *gorm.DB) {
	employees := []Employee{
		{
			FirstName: "Bruce",
			LastName:  "Wayne",
			JobCategory: JobCategoryLink{
				JobCategoryID:   models.JobCategorySuperHero,
				JobCategoryName: "Superhero",
			},
			Calendar: Calendar{
				Name: "Batman TAS",
				Appointments: []Appointment{
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
			JobCategory: JobCategoryLink{
				JobCategoryID:   models.JobCategoryVillain,
				JobCategoryName: "Villain",
			},
		},
	}

	for e := range employees {
		db.Debug().Create(&e)
	}

	db.Debug().Create(&Employee{
		FirstName: "Clark",
		LastName:  "Kent",
		JobCategory: JobCategoryLink{
			JobCategoryID:   models.JobCategorySuperHero,
			JobCategoryName: "Superhero",
		},
		Calendar: Calendar{
			Name: "Superman TAS",
			Appointments: []Appointment{
				{
					Title:       "World's Finest",
					Description: "The Joker steals a large piece of Kryptonite and then comes to Metropolis, offering to kill Superman for Lex Luthor in exchange for one billion dollars",
					StartTime:   time.Date(1997, time.October, 4, 9, 30, 0, 0, time.UTC),
					Length:      30,
					Attendees:   employees,
				},
			},
		},
	})
}
