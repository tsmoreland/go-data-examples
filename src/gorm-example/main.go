package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"github.com/tsmoreland/go-data-examples/src/gormexample/entities"
	"github.com/tsmoreland/go-data-examples/src/gormexample/infrastructure"
	"github.com/tsmoreland/go-data-examples/src/gormexample/models"
	"github.com/tsmoreland/go-data-examples/src/gormexample/shared"
	"time"
)

func main() {
	db, err := gorm.Open("postgres", "host=localhost port=5432 user=postgres password=P@55w0rd! dbname=gormexample sslmode=disable")
	if err != nil {
		panic(err.Error())
	}
	defer shared.CloseWithErrorLogging(db)
	sqlDB := db.DB()
	defer shared.CloseWithErrorLogging(sqlDB)
	infrastructure.VerifyConnectionOrPanic(sqlDB)
	infrastructure.ResetTables(db)

	infrastructure.CrudDemo(db)

	entities.CreateEmployeeTable(db)
	entities.CreateCalendarTable(db)
	entities.CreateAppointmentsTable(db)

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
	}

	for e := range employees {
		db.Debug().Save(&e)
	}

	db.Debug().Save(&entities.Employee{
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
	})

	e := entities.Employee{}
	c := entities.Calendar{}
	db.First(&e).Related(&c)

	fmt.Println(e) // e.Calendar will be nil, First does not load children by default
	fmt.Println(c)
}
