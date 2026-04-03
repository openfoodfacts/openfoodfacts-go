package openfoodfacts

import (
	"fmt"
	"net/url"
	"strconv"
	"strings"
)

type SearchQuery struct {
	Keyword string `json:"q"`
	Lang    string `json:"lang"` // e.g. "en", "fr" — auto-prefixes tags if set

	CategoryTags          []string `json:"categories_tags"`
	BrandsTags            []string `json:"brands_tags"`
	LabelsTags            []string `json:"labels_tags"`
	CountriesTags         []string `json:"countries_tags"`
	StatesTags            []string `json:"states_tags"`
	IngredientsTags       []string `json:"ingredients_tags"`
	AllergensTags         []string `json:"allergens_tags"`
	IngredientsAnalysis   []string `json:"ingredients_analysis_tags"`

	Page     int `json:"page"`
	PageSize int `json:"page_size"`
}

func (q SearchQuery) prefixTag(tag string) string {
	if q.Lang != "" && !strings.Contains(tag, ":") {
		return q.Lang + ":" + tag
	}
	return tag
}

func (q SearchQuery) ToQueryString() string {
	params := url.Values{}

	var qParts []string

	if q.Keyword != "" {
		qParts = append(qParts, q.Keyword)
	}

	for _, tag := range q.CategoryTags {
		qParts = append(qParts, fmt.Sprintf(`categories_tags:"%s"`, q.prefixTag(tag)))
	}

	for _, tag := range q.BrandsTags {
		qParts = append(qParts, fmt.Sprintf(`brands_tags:"%s"`, q.prefixTag(tag)))
	}

	for _, tag := range q.LabelsTags {
		qParts = append(qParts, fmt.Sprintf(`labels_tags:"%s"`, q.prefixTag(tag)))
	}

	for _, tag := range q.CountriesTags {
		qParts = append(qParts, fmt.Sprintf(`countries_tags:"%s"`, q.prefixTag(tag)))
	}

	for _, tag := range q.StatesTags {
		qParts = append(qParts, fmt.Sprintf(`states_tags:"%s"`, q.prefixTag(tag)))
	}

	for _, tag := range q.IngredientsTags {
		qParts = append(qParts, fmt.Sprintf(`ingredients_tags:"%s"`, q.prefixTag(tag)))
	}

	for _, tag := range q.AllergensTags {
		qParts = append(qParts, fmt.Sprintf(`allergens_tags:"%s"`, q.prefixTag(tag)))
	}

	for _, tag := range q.IngredientsAnalysis {
		qParts = append(qParts, fmt.Sprintf(`ingredients_analysis_tags:"%s"`, q.prefixTag(tag)))
	}

	if len(qParts) > 0 {
		params.Set("q", strings.Join(qParts, " "))
	}

	if q.Page > 0 {
		params.Set("page", strconv.Itoa(q.Page))
	}

	if q.PageSize > 0 {
		params.Set("page_size", strconv.Itoa(q.PageSize))
	}

	return params.Encode()
}
