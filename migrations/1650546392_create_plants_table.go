package migrations

import (
	"github.com/carprice-tech/migorm"
	"github.com/jinzhu/gorm"
)

func init() {
	migorm.RegisterMigration(&migrationInit{})
}

type migrationInit struct{}

type auth struct {
	Id       int64  `gorm:"primaryKey"`
	Username string `gorm:"not null;unique"`
	Login    string `gorm:"not null;unique"`
	Password string `gorm:"not null;unique"`
}

func (auth) AuthTableName() string {
	return "auth_data_table"
}

func (m *migrationInit) Up(db *gorm.DB, log migorm.Logger) error {
	return db.AutoMigrate(&auth{}).Error
}

func (m *migrationInit) Down(db *gorm.DB, log migorm.Logger) error {
	return db.DropTableIfExists(&auth{}).Error
}
