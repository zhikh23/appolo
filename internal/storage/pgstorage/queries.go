package pgstorage

const (
	save_material_query = `
		INSERT INTO materials (name, description, tags, url) VALUES
			($1, $2, $3, $4)
		RETURNING id
	`

	select_by_id = `
		SELECT
			id, name, description, tags, url
		FROM materials
		WHERE id=$1
	`

	select_by_tags = `
		SELECT
			id, name, description, tags, url
		FROM materials
		WHERE $1 && tags
	`
)
