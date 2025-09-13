package model

type File struct {
	Id        string `db:"id"`
	Filename  string `db:"filename"`
	Mimetype  string `db:"mimetype"`
	Author_id string `db:"author_id"`
}
