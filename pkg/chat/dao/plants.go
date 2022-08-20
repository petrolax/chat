package dao

type Plant struct {
	ID   uint `gorm:"primaryKey"`
	Name string
}
