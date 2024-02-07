package pgstorage

import "github.com/lib/pq"

type MaterialDto struct {
	Id 			uint64			`db:"id"`
	Name 		string			`db:"name"`
	Description string			`db:"description"`
	Url 		string			`db:"url"`
	Tags 		pq.StringArray	`db:"tags"`
}
