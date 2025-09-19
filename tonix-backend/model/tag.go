package model

import (
	"github.com/jmoiron/sqlx"
)

const (
	TagTypeUnsorted   = "unsorted"
	TagTypeInstrument = "instrument"
	TagTypeGenre      = "genre"
)

type Tag struct {
	Id     string `db:"id"`
	Name   string `db:"name"`
	Usages int    `db:"usages"`
	Type   string `db:"type"`
}

type TagFeatures struct {
	DB *sqlx.DB
}

func Tags(db *sqlx.DB) *TagFeatures {
	return &TagFeatures{
		DB: db,
	}
}

func (t *TagFeatures) ByName(name string) (*Tag, error) {
	var tag Tag

	err := t.DB.Get(&tag, "SELECT * FROM tags WHERE name = $1", name)
	if err != nil {
		return nil, err
	}

	return &tag, nil
}

func (t *TagFeatures) AddOrCreate(name string, tagType string) (*string, error) {
	foundTag, err := t.ByName(name)

	if foundTag != nil {
		foundTag.Usages++

		t.DB.Exec("UPDATE tags SET usages = $1 WHERE id = $2", foundTag.Usages, foundTag.Id)

		return &foundTag.Id, nil
	}

	var userId string
	row := t.DB.QueryRow(
		"INSERT INTO tags (name, type) VALUES ($1, $2) RETURNING id",
		name, tagType,
	)

	if err = row.Scan(&userId); err != nil {
		return nil, err
	}

	return &userId, err
}

type SearchTagsOpts struct {
	Name string
	Type *string
}

func (t *TagFeatures) Search(opts SearchTagsOpts, paginationOpts PaginationOpts) (*PaginatorResult[Tag], error) {
	paginator := NewPaginator[Tag](t.DB, "tags")

	var queryPart string
	var selectArgs []any

	queryPart += "name LIKE CONCAT('%', ?::text, '%')"
	selectArgs = append(selectArgs, opts.Name)

	if opts.Type != nil {
		queryPart += " AND type = ?"
		selectArgs = append(selectArgs, *opts.Type)
	}

	res, err := paginator.Select(paginationOpts, queryPart, "usages DESC", selectArgs...)
	if err != nil {
		return nil, err
	}

	return res, nil
}
