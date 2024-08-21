package domain

import "time"

type Room struct {
	ID          uint      `gorm:"primaryKey"`
	Name        string    `gorm:"type:varchar(200);not null"`
	Description string    `gorm:"type:text"`
	Updated     time.Time `gorm:"type:timestamp with time zone;not null;autoUpdateTime"`
	Created     time.Time `gorm:"type:timestamp with time zone;not null;autoCreateTime"`
	HostID      uint      `gorm:"index:idx_room_host_id"`
	TopicID     uint      `gorm:"index:idx_room_topic_id"`
	Host        User      `gorm:"foreignKey:HostID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;deferrable:InitiallyDeferred"`
	Topic       Topic     `gorm:"foreignKey:TopicID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;deferrable:InitiallyDeferred"`
	Since       string    `gorm:"-"`
}

type RoomParticipant struct {
	ID     uint `gorm:"primaryKey"`
	RoomID uint `gorm:"not null;index:idx_room_participants_room_id"`
	UserID uint `gorm:"not null;index:idx_room_participants_user_id"`
	Room   Room `gorm:"foreignKey:RoomID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;deferrable:InitiallyDeferred"`
	User   User `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;deferrable:InitiallyDeferred"`
}

type RoomWithDetails struct {
	Room
	ParticipantsCount int64
	Since             string
}

type Rooms struct {
	List  []RoomWithDetails
	Count int64
}

type RoomForm struct {
	TopicName   string
	Name        string
	Description string
}
