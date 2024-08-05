package migrations

import (
	"github.com/elyarsadig/studybud-go/internal/domain"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	return db.AutoMigrate(
		&domain.AuthGroup{},
		&domain.AuthPermission{},
		&domain.AuthGroupPermission{},
		&domain.Message{},
		&domain.Room{},
		&domain.RoomParticipant{},
		&domain.Topic{},
		&domain.User{},
		&domain.UserGroup{},
		&domain.UserPermission{},
		&domain.ContentType{},
	)
}
