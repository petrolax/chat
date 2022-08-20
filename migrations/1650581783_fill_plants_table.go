package migrations

import (
	"github.com/carprice-tech/migorm"
	"github.com/jinzhu/gorm"
)

func init() {
	migorm.RegisterMigration(&migrationfillPlantsTable{})
}

type migrationfillPlantsTable struct{}

type fillPlantsTable struct {
	ID   uint   `gorm:"primaryKey"`
	Name string `gorm:"not null"`
}

func (fillPlantsTable) TableName() string {
	return "plants"
}

func (m *migrationfillPlantsTable) Up(db *gorm.DB, log migorm.Logger) error {
	names := []fillPlantsTable{
		{
			Name: "Rose",
		},
		{
			Name: "Gladiolus",
		},
		{
			Name: "Trifolium",
		},
		{
			Name: "Tulipa",
		},
	}

	for _, name := range names {
		if err := db.Create(&name).Error; err != nil {
			return err
		}
	}
	return nil
}

func (m *migrationfillPlantsTable) Down(db *gorm.DB, log migorm.Logger) error {
	return nil
}
