package pagination

type Params struct {
    Page int
    Size int
}

type Page[T any] struct {
    Items      []T   `json:"items"`
    Total      int64 `json:"total"`
    Page       int   `json:"page"`
    Size       int   `json:"size"`
    TotalPages int   `json:"total_pages"`
}
