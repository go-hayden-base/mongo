package mongo

import (
	"strings"
)

const (
	ErrNoDefaultConnection    = "No default connection"
	ErrExistConnectionAlias   = "Exist connection alias"
	ErrNoDefaultDatabase      = "No default database"
	ErrNoConnection           = "No connection"
	ErrCannotSwitchCollection = "Can not switch collection"
)

func EqualError(err error, str string) bool {
	return str == err.Error() || strings.HasPrefix(err.Error(), str)
}
