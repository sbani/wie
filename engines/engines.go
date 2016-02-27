package engines

// Interface every search engine has to apply
type SearchEngine interface {
    GetLinks(query, site string) ([]string, error)
}
