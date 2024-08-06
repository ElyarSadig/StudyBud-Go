package configs

type ExtraData struct {
	HealthCheck                bool  `json:"health_check" yaml:"health_check"`
	SessionExpireDuration      int   `yaml:"session_expire_duration" json:"session_expire_duration"`
	MaxAttemptLoginTime        uint8 `yaml:"max_attempt_login_time" json:"max_attempt_login_time"`
	ServicePermissions         ServiceInfo
}

type ServiceInfo struct {
	ServiceName    string `yaml:"service_name" json:"service_name"`
	ServiceCode    string `yaml:"service_code" json:"service_code"`
	ServiceVersion string `yaml:"service_version" json:"service_version"`
}
