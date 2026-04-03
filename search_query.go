package openfoodfacts

import (
	"net/url"
	"strconv"
)

type SearchQuery struct {
	Keyword string `json:"query"`
	CategoryTags	[]string `json:"category_tags"`

	Page int `json:"page"`
	PageSize int `json:"page_size"`	
}

func (q SearchQuery) ToQueryString() string {
	params := url.Values{}
	params.Set("q", q.Keyword)

	if (len(q.CategoryTags) > 0) {
		for _, tag := range q.CategoryTags {
			params.Add("categories_tags", tag)
		}
	}

	if (q.Page > 0) {
		params.Set("page", strconv.Itoa(q.Page))
	}

	if (q.PageSize > 0) {
		params.Set("page_size", strconv.Itoa(q.PageSize))
	}

	return params.Encode()
}	
