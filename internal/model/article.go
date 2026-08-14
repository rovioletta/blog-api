package model

import "time"

type Article struct {
	ID      uint64
	Title   string
	Content string
	Tags    []string
}

type ParamsFilter struct {
	SearchTitle string
	SearchTags  []string
	CreatedFrom time.Time
	CreatedTo   time.Time
	UpdatedFrom time.Time
	UpdatedTo   time.Time
}

type Pagination struct {
	Limit  uint64
	Offset uint64
}

type Sort struct {
	Field string
	Order string
}

type ArticleFilter struct {
	Filter     *ParamsFilter
	Pagination *Pagination
	Sort       []Sort
}
