package errorHandler

type Error struct {
	httpCode int
	message  string
	params   []any
}
