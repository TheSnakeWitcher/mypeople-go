package contacts

import (
	"context"
	"database/sql"
	"fmt"
)

const delPeopleGroup = `
UPDATE people
SET groups = json_remove(groups,(
        SELECT fullkey FROM json_each(groups) WHERE value = ?
))
WHERE name = ? ;
`

type DelPeopleGroupParams struct {
	Name  string `db:"name"`
	Group string `db:"group"`
}

func (q *Queries) DelPeopleGroup(ctx context.Context, arg DelPeopleGroupParams) error {
	_, err := q.db.ExecContext(ctx, delPeople, arg.Group, arg.Name)
	return err
}

type TestQueryParams struct {
    Name string
}

const testQuery1 = `
SELECT * FROM people WHERE name = ? ;
`

const testQuery2 = `
SELECT groups FROM people WHERE name = ? ;
`

func (q *Queries) TestQuery(ctx context.Context, arg TestQueryParams) (*sql.Rows, error) {
	data, err := q.db.QueryContext(ctx, testQuery1,arg.Name)
	fmt.Print("test query 1 data: ", data,"\n\n")

	data, err = q.db.QueryContext(ctx, testQuery2,arg.Name)
	fmt.Print("test query 2 data: ", data,"\n\n")

	return data,err
}
