package db

import "database/sql"

func CheckError(e error) (exists bool, err error) {
	if e == sql.ErrNoRows {
		return false, nil
	}
	if e != nil {
		return false, err
	}
	return true, nil
}
