package menu

type Pagination struct {
	CurrentElements int  `json:"currentElements"`
	Last            bool `json:"last"`
	HasNext         bool `json:"hasNext"`
	TotalCount      int  `json:"totalCount"`
	BlockEnd        int  `json:"blockEnd"`
	TotalElements   int  `json:"totalElements"`
	BlockStart      int  `json:"blockStart"`
	Size            int  `json:"size"`
	TotalPages      int  `json:"totalPages"`
	HasPrevious     bool `json:"hasPrevious"`
	Page            int  `json:"page"`
	CurrentPage     int  `json:"currentPage"`
	First           bool `json:"first"`
}
