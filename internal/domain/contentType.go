package domain

type ContentType struct {
	ID       uint   `gorm:"primaryKey"`
	AppLabel string `gorm:"type:varchar(100);not null"`
	Model    string `gorm:"type:varchar(100);not null"`
}
