package migrations

import (
	"github.com/carprice-tech/migorm"
	"github.com/jinzhu/gorm"
)

func init() {
	migorm.RegisterMigration(&migrationInit{})
}

type migrationInit struct{}

type plant struct {
	ID   uint   `gorm:"primaryKey"`
	Name string `gorm:"not null;unique"`
}

func (plant) TableName() string {
	return "plants"
}

func (m *migrationInit) Up(db *gorm.DB, log migorm.Logger) error {
	return db.AutoMigrate(&plant{}).Error
}

func (m *migrationInit) Down(db *gorm.DB, log migorm.Logger) error {
	return db.DropTableIfExists(&plant{}).Error
}
