package model

type Article struct {
	ID      uint64
	Title   string
	Content string
	Tags    []string
}
