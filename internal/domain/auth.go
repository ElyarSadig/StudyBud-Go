package domain

type AuthGroup struct {
	ID   uint   `gorm:"primaryKey"`
	Name string `gorm:"type:varchar(150);not null;unique"`
}

type AuthPermission struct {
	ID            uint        `gorm:"primaryKey"`
	ContentTypeID uint        `gorm:"not null;index:idx_auth_permission_content_type_id"`
	Codename      string      `gorm:"type:varchar(100);not null"`
	Name          string      `gorm:"type:varchar(255);not null"`
	ContentType   ContentType `gorm:"foreignKey:ContentTypeID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;deferrable:InitiallyDeferred"`
}

type AuthGroupPermission struct {
	ID           uint           `gorm:"primaryKey"`
	GroupID      uint           `gorm:"not null;index:idx_auth_group_permissions_group_id"`
	PermissionID uint           `gorm:"not null;index:idx_auth_group_permissions_permission_id"`
	Group        AuthGroup      `gorm:"foreignKey:GroupID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;deferrable:InitiallyDeferred"`
	Permission   AuthPermission `gorm:"foreignKey:PermissionID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;deferrable:InitiallyDeferred"`
}
