package utils

import (
	"net/url"
	"strconv"
	"tonix/backend/consts"
	"tonix/backend/model"
)

func QueryGetOrDefault(queryParams url.Values, key string, def string) string {
	if !queryParams.Has(key) {
		return def
	}

	return queryParams.Get(key)
}

func PaginationParamsFromQuery(queryParams url.Values) (*model.PaginationOpts, error) {
	pageString := QueryGetOrDefault(queryParams, "page", "1")
	elementsOnPageString := QueryGetOrDefault(queryParams, "elementsOnPage", strconv.Itoa(consts.ELEMENTS_ON_PAGE_DEFAULT))

	page, err := strconv.Atoi(pageString)
	if err != nil {
		return nil, err
	}

	elementsOnPage, err := strconv.Atoi(elementsOnPageString)
	if err != nil {
		return nil, err
	}

	return &model.PaginationOpts{
		Page:           page,
		ElementsOnPage: elementsOnPage,
	}, err
}
