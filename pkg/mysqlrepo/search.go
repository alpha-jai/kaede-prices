package mysqlrepo

type Pager interface {
	Offset() int
	Limit() int
	Sort() []string

	SetPage(page int)
	SetPerPage(perpage int)
	SetSort(ol []string)
	AddSort(field string, asc bool)
}

type Filterer interface {
	Filters() map[string][]string
	SetFilters(filters map[string][]string)
	SetFilter(k string, v []string)
	AddFilter(k string, v []string)
}

type Searcher interface {
	Keyword() string
	SetKeyword(keyword string)
	KeywordEmbedding() []float32
	RAGUUIDs() []string
	SetRAGUUIDs(ragUUIDs []string)
}

type Lister interface {
	Pager
	Filterer
	Searcher
}
