package model

import (
	"github.com/jmoiron/sqlx"
)

type PaginationOpts struct {
	Page           int
	ElementsOnPage int
}

type PaginationData struct {
	Pages          int
	Page           int
	Count          int
	ElementsOnPage int
}

type PaginatorResult[T any] struct {
	Pagination   PaginationData
	SelectResult []T
}

type Paginator[T any] struct {
	db        *sqlx.DB
	tableName string
}

func NewPaginator[T any](db *sqlx.DB, tableName string) *Paginator[T] {
	return &Paginator[T]{
		db:        db,
		tableName: tableName,
	}
}

// Returns select results with LIMIT and OFFSET
//
// whereConditions - conditions that passed to WHERE. String builded will be:
//
// ```sql
// SELECT * FROM {Paginator.tableName} {selectPartStatement} LIMIT ? OFFSET ?
//
// SELECT COUNT(*) FROM {Paginator.tableName} {countSelectPartStatement}
// ```
func (p *Paginator[T]) Select(opts PaginationOpts, selectPartStatement string, countSelectPartStatement string, args ...any) (*PaginatorResult[T], error) {
	var results []T
	offset := opts.ElementsOnPage * (opts.Page - 1)
	query := "SELECT * FROM " + p.tableName + " " + selectPartStatement + " LIMIT ? OFFSET ?"

	argsWithPagination := args[:]
	argsWithPagination = append(argsWithPagination, opts.ElementsOnPage)
	argsWithPagination = append(argsWithPagination, offset)

	query = p.db.Rebind(query)
	if err := p.db.Select(&results, query, argsWithPagination...); err != nil {
		return nil, err
	}

	countQuery := "SELECT COUNT(*) FROM " + p.tableName + " " + countSelectPartStatement
	countQuery = p.db.Rebind(countQuery)

	row := p.db.QueryRow(countQuery, args...)

	var count int
	if err := row.Scan(&count); err != nil {
		return nil, err
	}

	pages := count / opts.ElementsOnPage

	if count%opts.ElementsOnPage != 0 {
		pages++
	}

	return &PaginatorResult[T]{
		Pagination: PaginationData{
			ElementsOnPage: opts.ElementsOnPage,
			Pages:          pages,
			Page:           opts.Page,
			Count:          count,
		},
		SelectResult: results,
	}, nil
}
