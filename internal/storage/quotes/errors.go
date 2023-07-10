package quotes

import "errors"

// ErrDBEmpty is returned when the db is empty.
var ErrDBEmpty = errors.New("db is empty")
