package seeders

import (
	"gorm.io/gorm"
	"DimasPrasetyo/backend-next/config/fakers"
)
type Seeder struct {
	Seeder interface{}
}
func RegisterSeeders(db *gorm.DB) []Seeder {

	return []Seeder{

		{fakers.AddressFaker(db)},
	}
}

func DBSeed(db *gorm.DB) error {
	for _, seeder := range RegisterSeeders(db) {
		err := db.Debug().Create(seeder.Seeder).Error
		if err != nil {
			return err
		}
	}

	return nil
}
