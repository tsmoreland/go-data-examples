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
	db.DropTableIfExists(&entities.Employee{}, &entities.Calendar{}, &entities.Appointment{}, "appointment_user")

	entities.CreateEmployeeTable(db)
	entities.CreateCalendarTable(db)
	entities.CreateAppointmentsTable(db)
}

func SeedDb(db *gorm.DB) error {
	tx := db.Begin()
	employees := map[string]*entities.Employee{
		"Batman": {
			FirstName: "Bruce",
			LastName:  "Wayne",
			JobCategory: entities.JobCategoryLink{
				JobCategoryID:   models.JobCategorySuperHero,
				JobCategoryName: "Superhero",
			},
		},
		"Harley": {
			FirstName: "Harley",
			LastName:  "Quinn",
			JobCategory: entities.JobCategoryLink{
				JobCategoryID:   models.JobCategoryVillain,
				JobCategoryName: "Villain",
			},
		},
		"Poison Ivy": {
			FirstName: "Pamela",
			LastName:  "Isley",
			JobCategory: entities.JobCategoryLink{
				JobCategoryID:   models.JobCategoryAntiHero,
				JobCategoryName: "Antihero",
			},
		},
		"Wonder Woman": {
			FirstName: "Diana",
			LastName:  "Prince",
			JobCategory: entities.JobCategoryLink{
				JobCategoryID:   models.JobCategorySuperHero,
				JobCategoryName: "Superhero",
			},
		},
		"The Flash": {
			FirstName: "Barry",
			LastName:  "Allen",
			JobCategory: entities.JobCategoryLink{
				JobCategoryID:   models.JobCategorySuperHero,
				JobCategoryName: "Superhero",
			},
		},
		"Superman": {
			FirstName: "Clark",
			LastName:  "Kent",
			JobCategory: entities.JobCategoryLink{
				JobCategoryID:   models.JobCategorySuperHero,
				JobCategoryName: "Superhero",
			},
		},
		"Green Lantern": {},
	}

	for _, e := range employees {
		e.Calendar = entities.Calendar{Name: "Animated Series"}
	}

	employees["Batman"].AddAppointment(&entities.Appointment{
		Title:       "On Leather Wings",
		Description: "A mysterious bat-like creature terrorizes Gotham City, causing the police force to pursue Batman. The Dark Knight must find the real perpetrator to clear his name.",
		StartTime:   time.Date(1992, time.May, 6, 9, 30, 0, 0, time.UTC),
		Length:      30,
	})
	employees["Batman"].AddAppointment(&entities.Appointment{
		Title:       "Secret Origins",
		Description: "As alien invaders who were accidentally awakened on Mars begin to take over the Earth, Superman and Batman rescue Martian Manhunter, who telepathically summons Wonder Woman, Hawkgirl, The Flash, and Green Lantern to help stop the invasion. Superman gathers everyone at the newly built Watchtower, funded by Batman, and asks them to be part of a team, which he dubs the Justice League.",
		StartTime:   time.Date(2001, time.November, 17, 9, 30, 0, 0, time.UTC),
		Length:      72,
		Attendees:   []*entities.Employee{employees["Superman"], employees["Wonder Woman"], employees["The Flash"]},
	})
	employees["Superman"].AddAppointment(&entities.Appointment{
		Title:       "World's Finest",
		Description: "The Joker steals a large piece of Kryptonite and then comes to Metropolis, offering to kill Superman for Lex Luthor in exchange for one billion dollars",
		StartTime:   time.Date(1997, time.October, 4, 9, 30, 0, 0, time.UTC),
		Length:      30,
		Attendees:   []*entities.Employee{employees["Batman"]},
	})
	employees["Wonder Woman"].AddAppointment(&entities.Appointment{
		Title:       "This Little Piggy",
		Description: "When Circe changes Wonder Woman into a pig, Batman and Zatanna must find her and change her back.",
		StartTime:   time.Date(2004, time.August, 28, 9, 30, 0, 0, time.UTC),
		Length:      24,
		Attendees:   []*entities.Employee{employees["Batman"]},
	})

	for _, e := range employees {
		if err := tx.Create(&e).Error; err != nil {
			tx.Rollback()
			return err
		}
	}

	tx.Commit()

	return nil
}
