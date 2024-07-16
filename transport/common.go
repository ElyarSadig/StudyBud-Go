package transport

import "time"

type HttpMethod string

const (
	POST   HttpMethod = "post"
	GET    HttpMethod = "get"
	DELETE HttpMethod = "delete"
	PUT    HttpMethod = "put"
)

const (
	_defaultReadTimeout     = 5 * time.Second
	_defaultWriteTimeout    = 5 * time.Second
	_defaultShutdownTimeout = 3 * time.Second
)
