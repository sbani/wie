package engines

// SearchEngine interface makes sure of the `GetLinks` func
type SearchEngine interface {
	GetLinks(query, site string) ([]string, error)
}
