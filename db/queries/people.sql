--------------------------------
-- people
--------------------------------

-- name: GetPeople :one
SELECT * FROM people WHERE name = ? LIMIT 1 ;

-- name: ListPeople :many
SELECT * FROM people ORDER BY name ;

-- name: AddPeople :one
INSERT INTO people(id,name, pic, groups, emails, phones, social_nets, wallets, locations, events, notes)
VALUES(NULL,?,?,'[]','[]','[]','[]','[]','[]','[]','[]') RETURNING * ;

-- name: DelPeople :exec
DELETE FROM people WHERE name = ? ;


--------------------------------
-- groups
--------------------------------

-- name: GetPeopleGroups :one
SELECT groups FROM people WHERE name = ? LIMIT 1 ;

-- name: AddPeopleGroup :one
UPDATE people SET groups = json_insert(groups,"$[#]",?) WHERE name = ? RETURNING * ;

-- NOTE:
-- sqlc can not compile json_each function so this was manually writed in 
-- ../../contacts/people.sql.aux.go
--name:DelPeopleGroup :exec
-- UPDATE people
-- SET groups = json_remove(groups,(
--         SELECT fullkey FROM json_each(groups) WHERE value = ?
-- ))
-- WHERE name = ? ;
