package confighandler

import "time"

type Config[T any] struct {
	ApplicationData   ApplicationData `yaml:"application_data" json:"application_data"`
	Log               *Log            `yaml:"log" json:"log"`
	Database          Database        `json:"database" yaml:"database"`
	Redis             Redis           `json:"redis" yaml:"redis"`
	Development       bool            `yaml:"development" json:"development"`
	TestMode          bool            `yaml:"test_mode" json:"test_mode"`
	HttpAddress       string          `yaml:"http_address" json:"http_address"`
	DefaultUploadPath string          `yaml:"default_upload_path" json:"default_upload_path"`
	FromEnvFile       bool            `yaml:"from_env_file" json:"from_env_file"` // true to load from .env instead of OS default env
	ExtraData         T               `yaml:"extra_data" json:"extra_data"`
}

type Redis struct {
	DialRetry       int           `yaml:"dial_retry"`
	MaxConn         int           `yaml:"max_conn"`
	IdleConn        int           `yaml:"idle_conn"`
	Address         string        `yaml:"address"`
	Port            string        `yaml:"port"`
	Password        string        `yaml:"password"`
	DB              int           `yaml:"db"`
	MaxRetries      int           `yaml:"max_retries"`
	MinRetryBackoff time.Duration `yaml:"min_retries_backoff"`
	MaxRetryBackoff time.Duration `yaml:"max_retries_backoff"`
	DialTimeout     time.Duration `yaml:"dial_timeout"`
	ReadTimeout     time.Duration `yaml:"read_timeout"`
	WriteTimeout    time.Duration `yaml:"write_timeout"`
	PoolSize        int           `yaml:"pool_size"` // It could giving value based on `10 * runtime.NumCPU()`
	MinIdleConns    int           `yaml:"min_idle_conns"`
	MaxConnAge      time.Duration `yaml:"max_conn_age"`
	PoolTimeout     time.Duration `yaml:"pool_timeout"`
	IdleTimeout     time.Duration `yaml:"idle_timeout"`
	Sentinel        Sentinel      `yaml:"sentinel"`
	PingTimeout     time.Duration `yaml:"ping_timeout"`
}

type Sentinel struct {
	Enabled    bool     `yaml:"enabled"`
	MasterName string   `yaml:"master_name"`
	Addresses  []string `yaml:"addresses"`
}

type Database struct {
	Name     string `yaml:"name" json:"name"`
	Host     string `yaml:"host" json:"host"`
	Port     string `yaml:"port" json:"port"`
	User     string `yaml:"user" json:"user"`
	Password string `yaml:"password" json:"password"`
}

type CircuitBreaker struct {
	ErrorThreshold   int           `yaml:"error_threshold" json:"error_threshold"`
	SuccessThreshold int           `yaml:"success_threshold" json:"success_threshold"`
	Timeout          time.Duration `yaml:"timeout" json:"timeout"`
}

type ApplicationData struct {
	Name    string `json:"name" yaml:"name"`
	Version string `json:"version" yaml:"version"`
}

type Log struct {
	Debug        bool   `yaml:"debug" json:"debug"`
	Handler      uint8  `yaml:"handler" json:"handler"` // Handler 0= console handler, 1= text handler, 2= json handler
	EnableCaller bool   `yaml:"enable_caller" json:"enable_caller"`
	SentryDSN    string `yaml:"sentry_dsn" json:"sentry_dsn"`
}
