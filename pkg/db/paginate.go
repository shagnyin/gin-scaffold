package db

func Paginate(page, pageSize uint64) (offset, limit int32) {
	if page == 0 {
		page = 1
	}
	if pageSize > 100 {
		pageSize = 100
	}
	if pageSize <= 0 {
		pageSize = 10
	}

	return int32((page - 1) * pageSize), int32(pageSize)
}
