package models

// Package pagination defines helper models for pagination related functionalities

const (
	// PageDefaultNumber int value 1
	PageDefaultNumber int = 1
	// PageDefaultSize int value 10
	PageDefaultSize int = 10
	// PageDefaultSortBy default sortBy string value
	PageDefaultSortBy string = "created_at"
	// PageDefaultSortDirectionDesc default sort direction descending order status
	PageDefaultSortDirectionDesc bool = true
	// PageSortDirectionAscending string value asc
	PageSortDirectionAscending string = "asc"
	// PageSortDirectionDescending string value desc
	PageSortDirectionDescending string = "desc"
	// SortByTags sort by tags on customers table
	SortByTags string = "tags"
)

// Page object for pagination purpose. Not persisted
type Page struct {
	Number            *int
	Size              *int
	SortBy            *string
	SortDirectionDesc *bool
}

// PageInfo holds pagination response info
type PageInfo struct {
	Page            int
	Size            int
	HasNextPage     bool
	HasPreviousPage bool
	TotalCount      int64
}

// NewPage creates a new pagination Page object
func NewPage(n int, s int, sBy string, sDirectionD bool) Page {
	return Page{
		Number:            &n,
		Size:              &s,
		SortBy:            &sBy,
		SortDirectionDesc: &sDirectionD,
	}
}

// NewPageWithDefaultSorting creates a new pagination Page object with system default values
func NewPageWithDefaultSorting(n int, s int) Page {
	return Page{
		Number: &n,
		Size:   &s,
	}
}
