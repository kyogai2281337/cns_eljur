package xlsx

import "errors"

var (
	ErrFileCorrupted       = errors.New("file corrupted")
	ErrUnableToSaveFile    = errors.New("unable to save file")
	ErrUnableToCreateSheet = errors.New("unable to create a sheet")
)
