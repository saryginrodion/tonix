package model

import (
	"time"

	"github.com/jmoiron/sqlx"
)

type Samplepack struct {
	Id            string    `db:"id"`
	Name          string    `db:"name"`
	Cost          int       `db:"cost"`
	PreviewId     string    `db:"preview_id"`
	IconId        string    `db:"icon_id"`
	AuthorId      string    `db:"author_id"`
	PurchaseCount int       `db:"purchase_count"`
	Status        string    `db:"status"`
	Likes         int       `db:"likes"`
	CreatedAt     time.Time `db:"created_at"`
}

type SamplepackTag struct {
	Name         string `db:"name"`
	SamplepackId string `db:"samplepack_id"`
}

type PurchasedSamplepack struct {
	UserId       string `db:"user_id"`
	SamplepackId string `db:"samplepack_id"`
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
