package migrations

import (
	"github.com/elyarsadig/studybud-go/internal/domain"
	"github.com/elyarsadig/studybud-go/pkg/logger"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB, logging logger.Logger) error {
	err := db.AutoMigrate(
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
	if err != nil {
		return err
	}
	logging.Info("successfully migrated the DB")
	return nil
}
