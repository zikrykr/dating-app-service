package pkg

const (
	DEFAULT_LIMIT = 10
	MAX_LIMIT     = 1000
)

type Pagination struct {
	CurrentPage     int64  `json:"current_page"`
	CurrentElements int64  `json:"current_elements"`
	TotalPages      int64  `json:"total_pages"`
	TotalElements   int64  `json:"total_elements"`
	SortBy          string `json:"sort_by"`
}

func ValidateLimit(limit int) int {
	if limit < 1 {
		return DEFAULT_LIMIT
	}

	if limit > MAX_LIMIT {
		return MAX_LIMIT
	}
	return limit
}

func ValidatePage(page int) int {
	if page < 1 {
		return 1
	}
	return page
}
