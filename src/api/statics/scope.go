package statics

import (
	"database/sql/driver"
	"strings"
)

//Scope - Deployment instance
type Scope string

const (
	//BETA - instance not productive
	BETA Scope = "beta"
	//MASTER - instance productive
	MASTER Scope = "master"
)

//GetName - return string name for Scope
func (s Scope) GetName() string {
	return string(s)
}

var availableScopes = map[string]Scope{
	BETA.GetName(): BETA,
	MASTER.GetName(): MASTER,
}

//GetScopeByString - return Scope if string is valid
func GetScopeByString(scope string) Scope {
	return availableScopes[strings.ToLower(scope)]
}

//Value - Implementation of valuer for database/sql
func (s Scope) Value() (driver.Value, error) {
	return string(s), nil
}
