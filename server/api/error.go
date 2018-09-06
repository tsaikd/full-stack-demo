package api

import "github.com/tsaikd/KDGoLib/errutil"

// errors
var (
	ErrAPINotFound  = errutil.NewFactory("API not found")
	ErrDatabaseGet1 = errutil.NewFactory("get %v from database failed")

	ErrInvalidParam = errutil.NewNamedFactory("ErrInvalidParam", "invalid parameters")
)
