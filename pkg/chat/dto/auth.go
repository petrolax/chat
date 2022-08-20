package dto

type Auth struct {
	Id       int64 `gorm:"primaryKey"`
	Username string
	Login    string
	Password string
}
