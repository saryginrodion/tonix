package model

import (
	"database/sql"
	"time"

	"github.com/jmoiron/sqlx"
)

type Samplepack struct {
	Id            Id             `db:"id"`
	Name          string         `db:"name"`
	Cost          int            `db:"cost"`
	PreviewId     sql.NullString `db:"preview_id"`
	IconId        sql.NullString `db:"icon_id"`
	AuthorId      Id             `db:"author_id"`
	PurchaseCount int            `db:"purchase_count"`
	Status        string         `db:"status"`
	Likes         int            `db:"likes"`
	CreatedAt     time.Time      `db:"created_at"`
}

type SamplepackTag struct {
	Name         string `db:"name"`
	SamplepackId Id     `db:"samplepack_id"`
}

type PurchasedSamplepack struct {
	UserId       Id `db:"user_id"`
	SamplepackId Id `db:"samplepack_id"`
	CreatedAt    time.Time
}

type SamplepackFeatures struct {
	db *sqlx.DB
}

func Samplepacks(db *sqlx.DB) *SamplepackFeatures {
	return &SamplepackFeatures{
		db: db,
	}
}

// func (f *SamplepackFeatures) New(author_id string)
