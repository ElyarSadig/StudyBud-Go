package domain

import "time"

type User struct {
    ID          uint      `gorm:"primaryKey"`
    Password    string    `gorm:"type:varchar(128);not null"`
    LastLogin   time.Time `gorm:"type:timestamp with time zone"`
    IsSuperuser bool      `gorm:"type:boolean;not null"`
    Username    string    `gorm:"type:varchar(150);not null;unique"`
    FirstName   string    `gorm:"type:varchar(150)"`
    LastName    string    `gorm:"type:varchar(150)"`
    Email       string    `gorm:"type:varchar(254);unique"`
    IsStaff     bool      `gorm:"type:boolean;not null"`
    IsActive    bool      `gorm:"type:boolean;not null"`
    DateJoined  time.Time `gorm:"type:timestamp with time zone;not null;autoCreateTime"`
    Bio         string    `gorm:"type:text"`
    Name        string    `gorm:"type:varchar(200)"`
    Avatar      string    `gorm:"type:varchar(100)"`
}

type UserGroup struct {
    ID      uint      `gorm:"primaryKey"`
    UserID  uint      `gorm:"not null;index:idx_user_groups_user_id"`
    GroupID uint      `gorm:"not null;index:idx_user_groups_group_id"`
    User    User      `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;deferrable:InitiallyDeferred"`
    Group   AuthGroup `gorm:"foreignKey:GroupID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;deferrable:InitiallyDeferred"`
}

type UserPermission struct {
    ID           uint          `gorm:"primaryKey"`
    UserID       uint          `gorm:"not null;index:idx_user_permissions_user_id"`
    PermissionID uint          `gorm:"not null;index:idx_user_permissions_permission_id"`
    User         User          `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;deferrable:InitiallyDeferred"`
    Permission   AuthPermission `gorm:"foreignKey:PermissionID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;deferrable:InitiallyDeferred"`
}
