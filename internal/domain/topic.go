package domain

type Topic struct {
	ID   uint   `gorm:"primaryKey"`
	Name string `gorm:"type:varchar(200);not null"`
}

type Topics struct {
	List  []TopicWithDetails
	Count int64
}

type TopicWithDetails struct {
	Name      string
	RoomCount int64
}
