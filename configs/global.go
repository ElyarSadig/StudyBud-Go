package configs

type CtxKey string

const UserCtxKey CtxKey = "user"

const (
	DefaultAvatar                  = "/static/images/avatar.svg"
	USERS_DB_NAME                  = "users"
	USER_PERMISSIONS_DB_NAME       = "user_permissions"
	USER_GROUPS_DB_NAME            = "user_groups"
	TOPICS_DB_NAME                 = "topics"
	ROOMS_DB_NAME                  = "rooms"
	ROOM_PARTICIPANTS_DB_NAME      = "room_participants"
	MESSAGES_DB_NAME               = "messages"
	CONTENT_TYPES_DB_NAME          = "content_types"
	AUTH_PERMISSIONS_DB_NAME       = "auth_permissions"
	AUTH_GROUPS_DB_NAME            = "auth_groups"
	AUTH_GROUP_PERMISSIONS_DB_NAME = "auth_group_permissions"
)

type ExtraData struct {
	HealthCheck           bool  `json:"health_check" yaml:"health_check"`
	SessionExpireDuration int   `yaml:"session_expire_duration" json:"session_expire_duration"`
	MaxAttemptLoginTime   uint8 `yaml:"max_attempt_login_time" json:"max_attempt_login_time"`
	ServicePermissions    ServiceInfo
}

type ServiceInfo struct {
	ServiceName    string `yaml:"service_name" json:"service_name"`
	ServiceCode    string `yaml:"service_code" json:"service_code"`
	ServiceVersion string `yaml:"service_version" json:"service_version"`
}
