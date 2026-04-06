package common

type PaginationMeta struct {
	Limit   int  `json:"limit"`
	Offset  int  `json:"offset"`
	Total   int  `json:"total"`
	HasNext bool `json:"has_next"`
}

func BuildPaginationMeta(limit, offset, total int) PaginationMeta {
	hasNext := offset+limit < total
	return PaginationMeta{
		Limit:   limit,
		Offset:  offset,
		Total:   total,
		HasNext: hasNext,
	}
}
