package statics

import (
	"database/sql/driver"
	"strings"
)

//Environment - Environment type
type Environment string

const (
	//DEVELOPMENT - development environment
	DEVELOPMENT Environment = "development"
	//PRODUCTION - productive environment
	PRODUCTION Environment = "prodution"
)

//GetName - return string name for Environment
func (s Environment) GetName() string {
	return string(s)
}

var availableEnvironments = map[string]Environment{
	DEVELOPMENT.GetName(): DEVELOPMENT,
	PRODUCTION.GetName(): PRODUCTION,
}

//GetEnvironmentByString - return Environment if string is valid
func GetEnvironmentByString(scope string) Environment {
	return availableEnvironments[strings.ToLower(scope)]
}

//Value - Implementation of valuer for database/sql
func (s Environment) Value() (driver.Value, error) {
	return string(s), nil
}
