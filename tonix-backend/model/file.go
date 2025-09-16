package model

import (
	"time"

	"github.com/jmoiron/sqlx"
)

type File struct {
	Id        string    `db:"id"`
	Filename  string    `db:"filename"`
	Path      string    `db:"path"`
	Mimetype  string    `db:"mimetype"`
	AuthorId  string    `db:"author_id"`
	CreatedAt time.Time `db:"created_at"`
}

type FileFeatures struct {
	DB *sqlx.DB
}

func Files(db *sqlx.DB) *FileFeatures {
	return &FileFeatures{DB: db}
}

func (f *FileFeatures) AddFile(file *File) (*string, error) {
	query, err := f.DB.NamedQuery("INSERT INTO files (author_id, filename, path, mimetype) VALUES (:author_id, :filename, :path, :mimetype) RETURNING id", file)
	if err != nil {
		return nil, err
	}

	query.Next()
	var id string
	if err = query.Scan(&id); err != nil {
		return nil, err
	}

	return &id, nil
}

func (f *FileFeatures) ById(fileId string) (*File, error) {
	var file File
	if err := f.DB.Get(&file, "SELECT * FROM files WHERE id = $1", fileId); err != nil {
		return nil, err
	}

	return &file, nil
}

func (f *FileFeatures) IsAuthor(userId string, fileId string) (bool, error) {
	var fileIds []string = make([]string, 1)
	if err := f.DB.Select(&fileIds, "SELECT * FROM files WHERE author_id = $1", userId); err != nil {
		return false, err
	}

	return len(fileIds) != 0, nil
}
