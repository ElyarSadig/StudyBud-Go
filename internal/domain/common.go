package domain

type Bridger interface {
	None()
}

// Bridge you can get specific useCase or repository
// Example:
//
//	userUseCase := Bridge[domain.UserUseCase](config.USER_DB, u.useCases)
func Bridge[T Bridger](dbName string, bridges map[string]Bridger) T {
	return bridges[dbName].(T)
}
