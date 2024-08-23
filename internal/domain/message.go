package domain

import (
	"time"
)

type Message struct {
	ID      uint      `gorm:"primaryKey"`
	Updated time.Time `gorm:"type:timestamp with time zone;not null;autoUpdateTime"`
	Created time.Time `gorm:"type:timestamp with time zone;not null;autoCreateTime"`
	Body    string    `gorm:"type:text;not null"`
	RoomID  uint      `gorm:"not null;index:idx_message_room_id"`
	UserID  uint      `gorm:"not null;index:idx_message_user_id"`
	Room    Room      `gorm:"foreignKey:RoomID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;deferrable:InitiallyDeferred"`
	User    User      `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;deferrable:InitiallyDeferred"`
	Since   string    `gorm:"-"`
}

type Messages struct {
	MessageList []Message
	Count       int64
}
