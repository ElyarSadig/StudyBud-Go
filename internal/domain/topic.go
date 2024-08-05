package domain

type Topic struct {
	ID   uint   `gorm:"primaryKey"`
	Name string `gorm:"type:varchar(200);not null"`
}
