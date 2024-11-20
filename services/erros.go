package services

import "errors"

var ErrEmptyData = errors.New("fetch 0 data from db.Query") // to deal with fetching 0 data as an error
