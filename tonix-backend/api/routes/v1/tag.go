package v1

import (
	"net/http"
	"tonix/backend/api/context"
	"tonix/backend/api/dto/view"
	"tonix/backend/api/utils"
	"tonix/backend/model"

	"github.com/saryginrodion/stackable"
)

var SearchTags = stackable.WrapFunc(
	func(ctx *context.Context, next func() error) error {
		queryValues := ctx.Request.URL.Query()
		paginationOpts, err := utils.PaginationParamsFromQuery(queryValues)
		if err != nil {
			return err
		}

		searchName := utils.QueryGetOrDefault(queryValues, "name", "")
		var searchType *string = nil
		if utils.QueryGetOrDefault(queryValues, "type", "") != "" {
			searchTypeUnref := utils.QueryGetOrDefault(queryValues, "type", "")
			searchType = &searchTypeUnref
		}

		tags := model.Tags(ctx.Shared.DB)
		searchedTagsResults, err := tags.Search(model.SearchTagsOpts{
			Name: searchName,
			Type: searchType,
		}, *paginationOpts)

		if err != nil {
			return err
		}

		tagViews := make([]view.TagView, 0)

		for _, tag := range searchedTagsResults.SelectResult {
			tagViews = append(tagViews, view.ToTagView(&tag))
		}

		ctx.Response, _ = stackable.JsonResponse(
			http.StatusOK,
			view.ToPaginated(&searchedTagsResults.Pagination, tagViews),
		)

		return next()
	},
)
