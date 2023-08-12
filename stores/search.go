package stores

import "strings"

type SearchOpts struct {
	Search string
	Sort   string
	Order  string
	Limit  int
	Page   int
}

func NewSearchOpts(opts ...Option) *SearchOpts {
	var searchOptions SearchOpts

	for _, opt := range opts {
		opt(&searchOptions)
	}

	return &searchOptions
}

type Option func(options *SearchOpts)

func WithSort(sort string) Option {
	return func(options *SearchOpts) {
		s := strings.ToLower(sort)

		if s == "desc" {
			options.Sort = "DESC"
			return
		}

		options.Sort = "ASC"
	}
}

func WithOrder(order string) Option {
	return func(options *SearchOpts) {
		options.Order = order
	}
}

func WithLimit(limit int) Option {
	return func(options *SearchOpts) {
		if limit < 1 || limit > 250 {
			options.Limit = 25
			return
		}

		options.Limit = limit
	}
}

func WithPage(page int) Option {
	return func(options *SearchOpts) {
		if page < 1 {
			options.Page = 1
			return
		}

		options.Page = page
	}
}

func WithSearch(s string) Option {
	return func(options *SearchOpts) {
		options.Search = s
	}
}
