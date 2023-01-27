// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0
// source: people.sql

package contacts

import (
	"context"
	"database/sql"
)

const addPeople = `-- name: AddPeople :one
;

INSERT INTO people(id,name, pic, groups, emails, phones, social_nets, wallets, locations, events, notes)
VALUES(NULL,?,'','[]','[]','[]','[]','[]','[]','[]','[]') RETURNING id, name, pic, "groups", emails, phones, social_nets, wallets, locations, events, notes
`

func (q *Queries) AddPeople(ctx context.Context, name string) (People, error) {
	row := q.db.QueryRowContext(ctx, addPeople, name)
	var i People
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Pic,
		&i.Groups,
		&i.Emails,
		&i.Phones,
		&i.SocialNets,
		&i.Wallets,
		&i.Locations,
		&i.Events,
		&i.Notes,
	)
	return i, err
}

const addPeopleGroup = `-- name: AddPeopleGroup :one
;

UPDATE people SET groups = json_insert(groups,"$[#]",?) WHERE name = ? RETURNING id, name, pic, "groups", emails, phones, social_nets, wallets, locations, events, notes
`

type AddPeopleGroupParams struct {
	JsonInsert interface{} `db:"json_insert" json:"json_insert"`
	Name       string      `db:"name" json:"name"`
}

func (q *Queries) AddPeopleGroup(ctx context.Context, arg AddPeopleGroupParams) (People, error) {
	row := q.db.QueryRowContext(ctx, addPeopleGroup, arg.JsonInsert, arg.Name)
	var i People
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Pic,
		&i.Groups,
		&i.Emails,
		&i.Phones,
		&i.SocialNets,
		&i.Wallets,
		&i.Locations,
		&i.Events,
		&i.Notes,
	)
	return i, err
}

const delPeople = `-- name: DelPeople :exec
;

DELETE FROM people WHERE name = ?
`

func (q *Queries) DelPeople(ctx context.Context, name string) error {
	_, err := q.db.ExecContext(ctx, delPeople, name)
	return err
}

const getPeople = `-- name: GetPeople :one

SELECT id, name, pic, "groups", emails, phones, social_nets, wallets, locations, events, notes FROM people WHERE name = ? LIMIT 1
`

// ------------------------------
// people
// ------------------------------
func (q *Queries) GetPeople(ctx context.Context, name string) (People, error) {
	row := q.db.QueryRowContext(ctx, getPeople, name)
	var i People
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Pic,
		&i.Groups,
		&i.Emails,
		&i.Phones,
		&i.SocialNets,
		&i.Wallets,
		&i.Locations,
		&i.Events,
		&i.Notes,
	)
	return i, err
}

const getPeopleGroups = `-- name: GetPeopleGroups :one
;



SELECT groups FROM people WHERE name = ? LIMIT 1
`

// ------------------------------
// groups
// ------------------------------
func (q *Queries) GetPeopleGroups(ctx context.Context, name string) (sql.NullString, error) {
	row := q.db.QueryRowContext(ctx, getPeopleGroups, name)
	var groups sql.NullString
	err := row.Scan(&groups)
	return groups, err
}

const listPeople = `-- name: ListPeople :many
;

SELECT id, name, pic, "groups", emails, phones, social_nets, wallets, locations, events, notes FROM people ORDER BY name
`

func (q *Queries) ListPeople(ctx context.Context) ([]People, error) {
	rows, err := q.db.QueryContext(ctx, listPeople)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []People
	for rows.Next() {
		var i People
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Pic,
			&i.Groups,
			&i.Emails,
			&i.Phones,
			&i.SocialNets,
			&i.Wallets,
			&i.Locations,
			&i.Events,
			&i.Notes,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
