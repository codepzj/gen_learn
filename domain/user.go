package domain

type User struct {
	Id       string `gorm:"primaryKey"`
	Username string `gorm:"type:varchar(128);unique;not null"`
	Password string `gorm:"type:varchar(128);not null"`
	Books    []Book `gorm:"foreignKey:Uid;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
